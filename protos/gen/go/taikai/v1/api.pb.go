// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: taikai/v1/api.proto

package taikaiv1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_taikai_v1_api_proto protoreflect.FileDescriptor

var file_taikai_v1_api_proto_rawDesc = []byte{
	0x0a, 0x13, 0x74, 0x61, 0x69, 0x6b, 0x61, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x69, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x74, 0x61, 0x69, 0x6b, 0x61, 0x69, 0x2e, 0x76, 0x31,
	0x1a, 0x18, 0x74, 0x61, 0x69, 0x6b, 0x61, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xbb, 0x07, 0x0a, 0x03, 0x41, 0x70, 0x69, 0x12, 0x58, 0x0a,
	0x0c, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x73, 0x12, 0x1e, 0x2e,
	0x74, 0x61, 0x69, 0x6b, 0x61, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74,
	0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e,
	0x74, 0x61, 0x69, 0x6b, 0x61, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x73,
	0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x3a, 0x01, 0x2a, 0x22, 0x0a, 0x2f, 0x76, 0x31,
	0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x73, 0x12, 0x57, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x73, 0x12, 0x18, 0x2e, 0x74, 0x61, 0x69, 0x6b, 0x61, 0x69,
	0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x19, 0x2e, 0x74, 0x61, 0x69, 0x6b, 0x61, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x12, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x0c, 0x2a, 0x0a, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x73,
	0x12, 0x50, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x73, 0x12, 0x16,
	0x2e, 0x74, 0x61, 0x69, 0x6b, 0x61, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x74, 0x61, 0x69, 0x6b, 0x61, 0x69, 0x2e,
	0x76, 0x31, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x73, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x11, 0x12, 0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x73, 0x2f, 0x6c, 0x69,
	0x73, 0x74, 0x12, 0x4d, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x73, 0x12,
	0x15, 0x2e, 0x74, 0x61, 0x69, 0x6b, 0x61, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x74, 0x61, 0x69, 0x6b, 0x61, 0x69, 0x2e,
	0x76, 0x31, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x73, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x10, 0x22, 0x0e, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x73, 0x2f, 0x67, 0x65,
	0x74, 0x12, 0x4b, 0x0a, 0x07, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x79, 0x12, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x10, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x0a, 0x12, 0x08, 0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x79, 0x12, 0x47,
	0x0a, 0x05, 0x52, 0x65, 0x61, 0x64, 0x79, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x0e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x08, 0x12,
	0x06, 0x2f, 0x72, 0x65, 0x61, 0x64, 0x79, 0x12, 0x54, 0x0a, 0x09, 0x55, 0x70, 0x73, 0x65, 0x72,
	0x74, 0x4f, 0x72, 0x67, 0x12, 0x1b, 0x2e, 0x74, 0x61, 0x69, 0x6b, 0x61, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x4f, 0x72, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x12, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x0c, 0x3a, 0x01, 0x2a, 0x22, 0x07, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x72, 0x67, 0x12, 0x56, 0x0a,
	0x07, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x72, 0x67, 0x12, 0x19, 0x2e, 0x74, 0x61, 0x69, 0x6b, 0x61,
	0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x72, 0x67, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x74, 0x61, 0x69, 0x6b, 0x61, 0x69, 0x2e, 0x76, 0x31, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x4f, 0x72, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x12, 0x0c, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x72, 0x67,
	0x2f, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x61, 0x0a, 0x0d, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x4f,
	0x72, 0x67, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1f, 0x2e, 0x74, 0x61, 0x69, 0x6b, 0x61, 0x69, 0x2e,
	0x76, 0x31, 0x2e, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x4f, 0x72, 0x67, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x3a, 0x01, 0x2a, 0x22, 0x0c, 0x2f, 0x76, 0x31, 0x2f,
	0x6f, 0x72, 0x67, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x12, 0x66, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74,
	0x4f, 0x72, 0x67, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x1e, 0x2e, 0x74, 0x61, 0x69, 0x6b, 0x61,
	0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x72, 0x67, 0x55, 0x73, 0x65, 0x72,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x74, 0x61, 0x69, 0x6b, 0x61,
	0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x72, 0x67, 0x55, 0x73, 0x65, 0x72,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x0f, 0x12, 0x0d, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x72, 0x67, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73,
	0x12, 0x51, 0x0a, 0x09, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x72, 0x67, 0x12, 0x18, 0x2e,
	0x74, 0x61, 0x69, 0x6b, 0x61, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x74, 0x61, 0x69, 0x6b, 0x61, 0x69,
	0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x0f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x09, 0x2a, 0x07, 0x2f, 0x76, 0x31, 0x2f,
	0x6f, 0x72, 0x67, 0x42, 0x9c, 0x01, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x61, 0x69, 0x6b,
	0x61, 0x69, 0x2e, 0x76, 0x31, 0x42, 0x08, 0x41, 0x70, 0x69, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x6f,
	0x72, 0x67, 0x65, 0x75, 0x74, 0x61, 0x68, 0x2f, 0x74, 0x61, 0x69, 0x6b, 0x61, 0x69, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x74, 0x61, 0x69,
	0x6b, 0x61, 0x69, 0x2f, 0x76, 0x31, 0x3b, 0x74, 0x61, 0x69, 0x6b, 0x61, 0x69, 0x76, 0x31, 0xa2,
	0x02, 0x03, 0x54, 0x58, 0x58, 0xaa, 0x02, 0x09, 0x54, 0x61, 0x69, 0x6b, 0x61, 0x69, 0x2e, 0x56,
	0x31, 0xca, 0x02, 0x09, 0x54, 0x61, 0x69, 0x6b, 0x61, 0x69, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x15,
	0x54, 0x61, 0x69, 0x6b, 0x61, 0x69, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0a, 0x54, 0x61, 0x69, 0x6b, 0x61, 0x69, 0x3a, 0x3a,
	0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_taikai_v1_api_proto_goTypes = []any{
	(*UpsertHellosRequest)(nil),  // 0: taikai.v1.UpsertHellosRequest
	(*DeleteRequest)(nil),        // 1: taikai.v1.DeleteRequest
	(*ListRequest)(nil),          // 2: taikai.v1.ListRequest
	(*GetRequest)(nil),           // 3: taikai.v1.GetRequest
	(*emptypb.Empty)(nil),        // 4: google.protobuf.Empty
	(*UpsertOrgRequest)(nil),     // 5: taikai.v1.UpsertOrgRequest
	(*ListOrgRequest)(nil),       // 6: taikai.v1.ListOrgRequest
	(*UpsertOrgUserRequest)(nil), // 7: taikai.v1.UpsertOrgUserRequest
	(*ListOrgUsersRequest)(nil),  // 8: taikai.v1.ListOrgUsersRequest
	(*Hellos)(nil),               // 9: taikai.v1.Hellos
	(*DeleteResponse)(nil),       // 10: taikai.v1.DeleteResponse
	(*ListOrgResponse)(nil),      // 11: taikai.v1.ListOrgResponse
	(*ListOrgUsersResponse)(nil), // 12: taikai.v1.ListOrgUsersResponse
}
var file_taikai_v1_api_proto_depIdxs = []int32{
	0,  // 0: taikai.v1.Api.UpsertHellos:input_type -> taikai.v1.UpsertHellosRequest
	1,  // 1: taikai.v1.Api.DeleteHellos:input_type -> taikai.v1.DeleteRequest
	2,  // 2: taikai.v1.Api.ListHellos:input_type -> taikai.v1.ListRequest
	3,  // 3: taikai.v1.Api.GetHellos:input_type -> taikai.v1.GetRequest
	4,  // 4: taikai.v1.Api.Healthy:input_type -> google.protobuf.Empty
	4,  // 5: taikai.v1.Api.Ready:input_type -> google.protobuf.Empty
	5,  // 6: taikai.v1.Api.UpsertOrg:input_type -> taikai.v1.UpsertOrgRequest
	6,  // 7: taikai.v1.Api.ListOrg:input_type -> taikai.v1.ListOrgRequest
	7,  // 8: taikai.v1.Api.UpsertOrgUser:input_type -> taikai.v1.UpsertOrgUserRequest
	8,  // 9: taikai.v1.Api.ListOrgUsers:input_type -> taikai.v1.ListOrgUsersRequest
	1,  // 10: taikai.v1.Api.DeleteOrg:input_type -> taikai.v1.DeleteRequest
	9,  // 11: taikai.v1.Api.UpsertHellos:output_type -> taikai.v1.Hellos
	10, // 12: taikai.v1.Api.DeleteHellos:output_type -> taikai.v1.DeleteResponse
	9,  // 13: taikai.v1.Api.ListHellos:output_type -> taikai.v1.Hellos
	9,  // 14: taikai.v1.Api.GetHellos:output_type -> taikai.v1.Hellos
	4,  // 15: taikai.v1.Api.Healthy:output_type -> google.protobuf.Empty
	4,  // 16: taikai.v1.Api.Ready:output_type -> google.protobuf.Empty
	4,  // 17: taikai.v1.Api.UpsertOrg:output_type -> google.protobuf.Empty
	11, // 18: taikai.v1.Api.ListOrg:output_type -> taikai.v1.ListOrgResponse
	4,  // 19: taikai.v1.Api.UpsertOrgUser:output_type -> google.protobuf.Empty
	12, // 20: taikai.v1.Api.ListOrgUsers:output_type -> taikai.v1.ListOrgUsersResponse
	10, // 21: taikai.v1.Api.DeleteOrg:output_type -> taikai.v1.DeleteResponse
	11, // [11:22] is the sub-list for method output_type
	0,  // [0:11] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_taikai_v1_api_proto_init() }
func file_taikai_v1_api_proto_init() {
	if File_taikai_v1_api_proto != nil {
		return
	}
	file_taikai_v1_requests_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_taikai_v1_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_taikai_v1_api_proto_goTypes,
		DependencyIndexes: file_taikai_v1_api_proto_depIdxs,
	}.Build()
	File_taikai_v1_api_proto = out.File
	file_taikai_v1_api_proto_rawDesc = nil
	file_taikai_v1_api_proto_goTypes = nil
	file_taikai_v1_api_proto_depIdxs = nil
}
