package file

import (
	"fmt"

	"github.com/batazor/shortlink/pkg/shortdb/engine/file/cursor"
	v1 "github.com/batazor/shortlink/pkg/shortdb/query/v1"
	table "github.com/batazor/shortlink/pkg/shortdb/table/v1"
)

func (f *file) Select(query *v1.Query) ([]*table.Row, error) {
	f.mc.Lock()
	defer f.mc.Unlock()

	// check table
	t := f.database.Tables[query.TableName]
	if t == nil {
		return nil, fmt.Errorf("at SELECT: not exist table")
	}

	// set flags
	isLimit := query.Limit >= 0

	// response
	response := make([]*table.Row, 0)

	currentRow, err := cursor.New(t, false)
	if err != nil {
		return nil, fmt.Errorf("at SELECT: error create a new cursor")
	}

	for !currentRow.EndOfTable && isLimit {
		// load data
		if t.Pages[currentRow.PageId] == nil {
			pagePath := fmt.Sprintf("%s/%s_%s_%d.page", f.path, f.database.Name, t.Name, currentRow.PageId)
			payload, errLoadPage := f.loadPage(pagePath)
			if errLoadPage != nil {
				return nil, errLoadPage
			}

			if t.Pages == nil {
				t.Pages = make(map[int32]*table.Page, 0)
			}

			t.Pages[currentRow.PageId] = payload
		}

		// get value
		record, errGetValue := currentRow.Value()
		if errGetValue != nil {
			return nil, errGetValue
		}

		for _, field := range query.Fields {
			if record.Value[field] == nil {
				return nil, fmt.Errorf("at SELECT: incorrect name fields %s in table %s", field, query.TableName)
			}
		}
		response = append(response, record)

		if isLimit {
			query.Limit -= 1

			if query.Limit < 1 {
				isLimit = false
			}
		}

		currentRow.Advance()
	}

	return response, nil
}

func (f *file) Update(query *v1.Query) error {
	// TODO implement me
	return nil
}

func (f *file) Insert(query *v1.Query) error {
	f.mc.Lock()
	defer f.mc.Unlock()

	// check the table's existence
	t := f.database.Tables[query.TableName]
	if t == nil {
		return fmt.Errorf("at INSERT INTO: not exist table")
	}

	// check if a new page needs to be created
	_, err := f.addPage(query.TableName)
	if err != nil {
		return fmt.Errorf("at INSERT INTO: error create a new page")
	}

	if t.Stats.PageCount > -1 && t.Pages[t.Stats.PageCount] == nil {
		// load page
		pagePath := fmt.Sprintf("%s/%s_%s_%d.page", f.path, f.database.Name, t.Name, t.Stats.PageCount)
		payload, errLoadPage := f.loadPage(pagePath)
		if errLoadPage != nil {
			return errLoadPage
		}

		if t.Pages == nil {
			t.Pages = make(map[int32]*table.Page, 0)
		}

		t.Pages[t.Stats.PageCount] = payload
	}

	// insert to last page
	currentRow, err := cursor.New(t, true)
	if err != nil {
		return fmt.Errorf("at INSERT INTO: error create a new cursor")
	}

	// iterator to next value
	currentRow.Advance()

	row, err := currentRow.Value()
	if err != nil {
		switch err.(type) {
		case *cursor.ErrorGetPage:
			// hmm...
		default:
			return fmt.Errorf("at INSERT INTO: error get value from cursor")
		}
	}

	// check values and create row record
	record := table.Row{
		Value: make(map[string][]byte),
	}
	for index, field := range query.Fields {
		if t.Fields[field].String() == "" {
			return fmt.Errorf("at INSERT INTO: incorrect type fields %s in table %s", field, query.TableName)
		}

		record.Value[field] = []byte(query.Inserts[0].Items[index])
	}
	row.Value = record.Value

	// update stats
	t.Stats.RowsCount += 1

	return nil
}

func (f *file) Delete(query *v1.Query) error {
	// TODO implement me
	return nil
}
