// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: shortdb/domain/field/v1/field.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Type int32

const (
	Type_TYPE_UNSPECIFIED Type = 0
	Type_TYPE_INTEGER     Type = 1
	Type_TYPE_STRING      Type = 2
	Type_TYPE_BOOLEAN     Type = 3
)

// Enum value maps for Type.
var (
	Type_name = map[int32]string{
		0: "TYPE_UNSPECIFIED",
		1: "TYPE_INTEGER",
		2: "TYPE_STRING",
		3: "TYPE_BOOLEAN",
	}
	Type_value = map[string]int32{
		"TYPE_UNSPECIFIED": 0,
		"TYPE_INTEGER":     1,
		"TYPE_STRING":      2,
		"TYPE_BOOLEAN":     3,
	}
)

func (x Type) Enum() *Type {
	p := new(Type)
	*p = x
	return p
}

func (x Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Type) Descriptor() protoreflect.EnumDescriptor {
	return file_shortdb_domain_field_v1_field_proto_enumTypes[0].Descriptor()
}

func (Type) Type() protoreflect.EnumType {
	return &file_shortdb_domain_field_v1_field_proto_enumTypes[0]
}

func (x Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Type.Descriptor instead.
func (Type) EnumDescriptor() ([]byte, []int) {
	return file_shortdb_domain_field_v1_field_proto_rawDescGZIP(), []int{0}
}

var File_shortdb_domain_field_v1_field_proto protoreflect.FileDescriptor

var file_shortdb_domain_field_v1_field_proto_rawDesc = []byte{
	0x0a, 0x23, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x64, 0x62, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x64, 0x62, 0x2e, 0x64,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x2a, 0x51,
	0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x10, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x49, 0x4e, 0x54, 0x45, 0x47, 0x45, 0x52, 0x10, 0x01, 0x12, 0x0f,
	0x0a, 0x0b, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x54, 0x52, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12,
	0x10, 0x0a, 0x0c, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x42, 0x4f, 0x4f, 0x4c, 0x45, 0x41, 0x4e, 0x10,
	0x03, 0x42, 0x3a, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x62, 0x61, 0x74, 0x61, 0x7a, 0x6f, 0x72, 0x2f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x6c, 0x69, 0x6e,
	0x6b, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x64, 0x62, 0x2f, 0x64, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_shortdb_domain_field_v1_field_proto_rawDescOnce sync.Once
	file_shortdb_domain_field_v1_field_proto_rawDescData = file_shortdb_domain_field_v1_field_proto_rawDesc
)

func file_shortdb_domain_field_v1_field_proto_rawDescGZIP() []byte {
	file_shortdb_domain_field_v1_field_proto_rawDescOnce.Do(func() {
		file_shortdb_domain_field_v1_field_proto_rawDescData = protoimpl.X.CompressGZIP(file_shortdb_domain_field_v1_field_proto_rawDescData)
	})
	return file_shortdb_domain_field_v1_field_proto_rawDescData
}

var file_shortdb_domain_field_v1_field_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_shortdb_domain_field_v1_field_proto_goTypes = []interface{}{
	(Type)(0), // 0: shortdb.domain.field.v1.Type
}
var file_shortdb_domain_field_v1_field_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_shortdb_domain_field_v1_field_proto_init() }
func file_shortdb_domain_field_v1_field_proto_init() {
	if File_shortdb_domain_field_v1_field_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_shortdb_domain_field_v1_field_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_shortdb_domain_field_v1_field_proto_goTypes,
		DependencyIndexes: file_shortdb_domain_field_v1_field_proto_depIdxs,
		EnumInfos:         file_shortdb_domain_field_v1_field_proto_enumTypes,
	}.Build()
	File_shortdb_domain_field_v1_field_proto = out.File
	file_shortdb_domain_field_v1_field_proto_rawDesc = nil
	file_shortdb_domain_field_v1_field_proto_goTypes = nil
	file_shortdb_domain_field_v1_field_proto_depIdxs = nil
}
