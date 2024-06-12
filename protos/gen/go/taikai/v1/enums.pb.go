// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: taikai/v1/enums.proto

package taikaiv1

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

type HelloType int32

const (
	HelloType_HELLO_GENERIC  HelloType = 0
	HelloType_HELLO_SPECIFIC HelloType = 1
)

// Enum value maps for HelloType.
var (
	HelloType_name = map[int32]string{
		0: "HELLO_GENERIC",
		1: "HELLO_SPECIFIC",
	}
	HelloType_value = map[string]int32{
		"HELLO_GENERIC":  0,
		"HELLO_SPECIFIC": 1,
	}
)

func (x HelloType) Enum() *HelloType {
	p := new(HelloType)
	*p = x
	return p
}

func (x HelloType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (HelloType) Descriptor() protoreflect.EnumDescriptor {
	return file_taikai_v1_enums_proto_enumTypes[0].Descriptor()
}

func (HelloType) Type() protoreflect.EnumType {
	return &file_taikai_v1_enums_proto_enumTypes[0]
}

func (x HelloType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use HelloType.Descriptor instead.
func (HelloType) EnumDescriptor() ([]byte, []int) {
	return file_taikai_v1_enums_proto_rawDescGZIP(), []int{0}
}

var File_taikai_v1_enums_proto protoreflect.FileDescriptor

var file_taikai_v1_enums_proto_rawDesc = []byte{
	0x0a, 0x15, 0x74, 0x61, 0x69, 0x6b, 0x61, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x75, 0x6d,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x74, 0x61, 0x69, 0x6b, 0x61, 0x69, 0x2e,
	0x76, 0x31, 0x2a, 0x32, 0x0a, 0x09, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x11, 0x0a, 0x0d, 0x48, 0x45, 0x4c, 0x4c, 0x4f, 0x5f, 0x47, 0x45, 0x4e, 0x45, 0x52, 0x49, 0x43,
	0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x48, 0x45, 0x4c, 0x4c, 0x4f, 0x5f, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x43, 0x10, 0x01, 0x42, 0x9e, 0x01, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x2e, 0x74,
	0x61, 0x69, 0x6b, 0x61, 0x69, 0x2e, 0x76, 0x31, 0x42, 0x0a, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x66, 0x6f, 0x72, 0x67, 0x65, 0x75, 0x74, 0x61, 0x68, 0x2f, 0x74, 0x61, 0x69,
	0x6b, 0x61, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67,
	0x6f, 0x2f, 0x74, 0x61, 0x69, 0x6b, 0x61, 0x69, 0x2f, 0x76, 0x31, 0x3b, 0x74, 0x61, 0x69, 0x6b,
	0x61, 0x69, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x54, 0x58, 0x58, 0xaa, 0x02, 0x09, 0x54, 0x61, 0x69,
	0x6b, 0x61, 0x69, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x09, 0x54, 0x61, 0x69, 0x6b, 0x61, 0x69, 0x5c,
	0x56, 0x31, 0xe2, 0x02, 0x15, 0x54, 0x61, 0x69, 0x6b, 0x61, 0x69, 0x5c, 0x56, 0x31, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0a, 0x54, 0x61, 0x69,
	0x6b, 0x61, 0x69, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_taikai_v1_enums_proto_rawDescOnce sync.Once
	file_taikai_v1_enums_proto_rawDescData = file_taikai_v1_enums_proto_rawDesc
)

func file_taikai_v1_enums_proto_rawDescGZIP() []byte {
	file_taikai_v1_enums_proto_rawDescOnce.Do(func() {
		file_taikai_v1_enums_proto_rawDescData = protoimpl.X.CompressGZIP(file_taikai_v1_enums_proto_rawDescData)
	})
	return file_taikai_v1_enums_proto_rawDescData
}

var file_taikai_v1_enums_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_taikai_v1_enums_proto_goTypes = []any{
	(HelloType)(0), // 0: taikai.v1.HelloType
}
var file_taikai_v1_enums_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_taikai_v1_enums_proto_init() }
func file_taikai_v1_enums_proto_init() {
	if File_taikai_v1_enums_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_taikai_v1_enums_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_taikai_v1_enums_proto_goTypes,
		DependencyIndexes: file_taikai_v1_enums_proto_depIdxs,
		EnumInfos:         file_taikai_v1_enums_proto_enumTypes,
	}.Build()
	File_taikai_v1_enums_proto = out.File
	file_taikai_v1_enums_proto_rawDesc = nil
	file_taikai_v1_enums_proto_goTypes = nil
	file_taikai_v1_enums_proto_depIdxs = nil
}
