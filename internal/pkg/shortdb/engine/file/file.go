package file

import (
	"fmt"
	"io/ioutil"
	"os"
	"sync"

	"github.com/golang/protobuf/proto"
	"github.com/spf13/viper"

	"github.com/batazor/shortlink/internal/pkg/shortdb/engine/options"
	v1 "github.com/batazor/shortlink/internal/pkg/shortdb/query/v1"
	table "github.com/batazor/shortlink/internal/pkg/shortdb/table/v1"
)

type file struct {
	mc sync.Mutex

	name string
	path string

	database *table.DataBase
}

func New(opts ...options.Option) (*file, error) {
	viper.AutomaticEnv()
	viper.SetDefault("SHORTDB_DEFAULT_DATABASE", "public.db") // ShortDB default database

	var err error
	f := &file{
		name: viper.GetString("SHORTDB_DEFAULT_DATABASE"),
		database: &table.DataBase{
			Tables: make(map[string]*table.Table),
		},
	}

	for _, opt := range opts {
		if err := opt(f); err != nil {
			panic(err)
		}
	}

	// if not set path, set temp directory
	if f.path == "" {
		f.path, err = ioutil.TempDir(os.TempDir(), "shortdb_")
		if err != nil {
			return nil, err
		}
	}

	// init db
	err = f.init()
	if err != nil {
		return nil, err
	}

	return f, nil
}

func (f *file) Exec(query *v1.Query) error {
	switch query.Type {
	case v1.Type_TYPE_UNSPECIFIED:
		return fmt.Errorf("exec: incorret type")
	case v1.Type_TYPE_SELECT:
		return f.Select()
	case v1.Type_TYPE_UPDATE:
		return f.Update()
	case v1.Type_TYPE_INSERT:
		return f.Insert()
	case v1.Type_TYPE_DELETE:
		return f.Delete()
	case v1.Type_TYPE_CREATE_TABLE:
		return f.CreateTable(query.TableName, query.TableFields)
	case v1.Type_TYPE_DROP_TABLE:
		return f.DropTable(query.TableName)
	}

	return nil
}

func (f *file) init() error {
	f.mc.Lock()
	defer f.mc.Unlock()

	path := fmt.Sprintf("%s/%s", f.path, f.name)

	// create directory if not exist
	err := os.MkdirAll(f.path, os.ModePerm)
	if err != nil {
		return err
	}

	// create file if not exist
	file, err := os.OpenFile(path, os.O_RDONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()

	// read file
	payload, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	if len(payload) != 0 {
		err = proto.Unmarshal(payload, f.database)
		if err != nil {
			return err
		}
	}

	return nil
}

func (f *file) Close() error {
	f.mc.Lock()
	defer f.mc.Unlock()

	path := fmt.Sprintf("%s/%s", f.path, f.name)

	// create file if not exist
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()

	payload, err := proto.Marshal(f.database)
	if err != nil {
		return err
	}

	_, err = file.Write(payload)
	if err != nil {
		return err
	}

	return nil
}
