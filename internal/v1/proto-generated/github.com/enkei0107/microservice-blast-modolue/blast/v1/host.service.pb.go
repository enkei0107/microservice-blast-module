// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: core/blast/v1/service/host.service.proto

package blastv1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_core_blast_v1_service_host_service_proto protoreflect.FileDescriptor

var file_core_blast_v1_service_host_service_proto_rawDesc = []byte{
	0x0a, 0x28, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x62, 0x6c, 0x61, 0x73, 0x74, 0x2f, 0x76, 0x31, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x68, 0x6f, 0x73, 0x74, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x68, 0x6f, 0x73, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x28, 0x63, 0x6f, 0x72, 0x65, 0x2f,
	0x62, 0x6c, 0x61, 0x73, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x2f, 0x68, 0x6f, 0x73, 0x74, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x2a, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x62, 0x6c, 0x61, 0x73, 0x74, 0x2f,
	0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2f, 0x68, 0x6f, 0x73, 0x74,
	0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32,
	0x7b, 0x0a, 0x0b, 0x48, 0x6f, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x6c,
	0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x48, 0x6f, 0x73, 0x74, 0x12, 0x22, 0x2e, 0x68,
	0x6f, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x48, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x24, 0x2e, 0x68, 0x6f, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x48, 0x6f, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x3a, 0x01,
	0x2a, 0x22, 0x09, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x6f, 0x73, 0x74, 0x73, 0x42, 0xc7, 0x01, 0x0a,
	0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x68, 0x6f, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x42, 0x10, 0x48, 0x6f, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x6b, 0x65, 0x69, 0x30, 0x31, 0x30, 0x37, 0x2f, 0x6d,
	0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2d, 0x62, 0x6c, 0x61, 0x73,
	0x74, 0x2d, 0x6d, 0x6f, 0x64, 0x6f, 0x6c, 0x75, 0x65, 0x2f, 0x62, 0x6c, 0x61, 0x73, 0x74, 0x2f,
	0x76, 0x31, 0x3b, 0x62, 0x6c, 0x61, 0x73, 0x74, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x48, 0x56, 0x53,
	0xaa, 0x02, 0x0f, 0x48, 0x6f, 0x73, 0x74, 0x2e, 0x56, 0x31, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0xca, 0x02, 0x0f, 0x48, 0x6f, 0x73, 0x74, 0x5c, 0x56, 0x31, 0x5c, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0xe2, 0x02, 0x1b, 0x48, 0x6f, 0x73, 0x74, 0x5c, 0x56, 0x31, 0x5c, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x11, 0x48, 0x6f, 0x73, 0x74, 0x3a, 0x3a, 0x56, 0x31, 0x3a, 0x3a, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_core_blast_v1_service_host_service_proto_goTypes = []any{
	(*CreateHostRequest)(nil),  // 0: host.v1.request.CreateHostRequest
	(*CreateHostResponse)(nil), // 1: host.v1.response.CreateHostResponse
}
var file_core_blast_v1_service_host_service_proto_depIdxs = []int32{
	0, // 0: host.v1.service.HostService.CreateHost:input_type -> host.v1.request.CreateHostRequest
	1, // 1: host.v1.service.HostService.CreateHost:output_type -> host.v1.response.CreateHostResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_core_blast_v1_service_host_service_proto_init() }
func file_core_blast_v1_service_host_service_proto_init() {
	if File_core_blast_v1_service_host_service_proto != nil {
		return
	}
	file_core_blast_v1_request_host_request_proto_init()
	file_core_blast_v1_response_host_response_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_core_blast_v1_service_host_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_core_blast_v1_service_host_service_proto_goTypes,
		DependencyIndexes: file_core_blast_v1_service_host_service_proto_depIdxs,
	}.Build()
	File_core_blast_v1_service_host_service_proto = out.File
	file_core_blast_v1_service_host_service_proto_rawDesc = nil
	file_core_blast_v1_service_host_service_proto_goTypes = nil
	file_core_blast_v1_service_host_service_proto_depIdxs = nil
}
