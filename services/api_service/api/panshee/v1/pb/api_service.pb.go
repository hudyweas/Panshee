// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0-devel
// 	protoc        v3.21.6
// source: panshee/v1/proto/api_service.proto

package pb

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
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

var File_panshee_v1_proto_api_service_proto protoreflect.FileDescriptor

var file_panshee_v1_proto_api_service_proto_rawDesc = []byte{
	0x0a, 0x22, 0x70, 0x61, 0x6e, 0x73, 0x68, 0x65, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x61, 0x6e, 0x73, 0x68, 0x65,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32,
	0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1f, 0x70, 0x61, 0x6e, 0x73, 0x68, 0x65, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x20, 0x70, 0x61, 0x6e, 0x73, 0x68, 0x65, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x70, 0x61, 0x6e, 0x73, 0x68, 0x65, 0x65, 0x2f, 0x76, 0x31,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x5f, 0x77, 0x61, 0x6c, 0x6c, 0x65,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x93, 0x08, 0x0a, 0x0a, 0x41, 0x70, 0x69, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x7d, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x12, 0x27, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x61, 0x6e, 0x73, 0x68,
	0x65, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x70, 0x61, 0x6e, 0x73, 0x68, 0x65, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x22, 0x2a, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x0d, 0x22, 0x08, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x3a, 0x01, 0x2a, 0xda, 0x41,
	0x14, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x2c, 0x75, 0x73, 0x65, 0x72, 0x2c, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x88, 0x01, 0x0a, 0x09, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x55,
	0x73, 0x65, 0x72, 0x12, 0x26, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x61, 0x6e, 0x73, 0x68, 0x65,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x70, 0x61, 0x6e, 0x73, 0x68, 0x65, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x22, 0x0e, 0x2f, 0x76,
	0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x3a, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x3a, 0x01, 0x2a, 0xda,
	0x41, 0x0e, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x2c, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x12, 0x6d, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x24, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x70, 0x61, 0x6e, 0x73, 0x68, 0x65, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x61, 0x6e, 0x73, 0x68, 0x65, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x22, 0x20, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x13, 0x12, 0x11, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65,
	0x3d, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x2a, 0x7d, 0xda, 0x41, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0xc2, 0x01, 0x0a, 0x12, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x41, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x32, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x61, 0x6e,
	0x73, 0x68, 0x65, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65,
	0x66, 0x72, 0x65, 0x73, 0x68, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x44, 0x54, 0x4f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x33, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x70, 0x61, 0x6e, 0x73, 0x68, 0x65, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x44, 0x54, 0x4f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x43, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x26, 0x22, 0x21, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x70, 0x61,
	0x72, 0x65, 0x6e, 0x74, 0x3d, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x2a, 0x7d, 0x2f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x3a, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x3a, 0x01, 0x2a, 0xda, 0x41, 0x14,
	0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x2c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x90, 0x01, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57,
	0x61, 0x6c, 0x6c, 0x65, 0x74, 0x12, 0x29, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x61, 0x6e, 0x73,
	0x68, 0x65, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x61, 0x6e, 0x73, 0x68, 0x65, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x44, 0x54,
	0x4f, 0x22, 0x34, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x25, 0x22, 0x1b, 0x2f, 0x76, 0x31, 0x2f, 0x7b,
	0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x3d, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x2a, 0x7d, 0x2f, 0x77,
	0x61, 0x6c, 0x6c, 0x65, 0x74, 0x73, 0x3a, 0x06, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0xda, 0x41,
	0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x9b, 0x01, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x57,
	0x61, 0x6c, 0x6c, 0x65, 0x74, 0x73, 0x12, 0x2d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x61, 0x6e,
	0x73, 0x68, 0x65, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65,
	0x74, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x44, 0x54, 0x4f, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x61, 0x6e, 0x73,
	0x68, 0x65, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74,
	0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x44, 0x54, 0x4f, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x12, 0x1d, 0x2f,
	0x76, 0x31, 0x2f, 0x7b, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x3d, 0x75, 0x73, 0x65, 0x72, 0x2f,
	0x2a, 0x2f, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x73, 0x2f, 0x2a, 0x7d, 0xda, 0x41, 0x06, 0x70,
	0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x95, 0x01, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x12, 0x29, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x61, 0x6e,
	0x73, 0x68, 0x65, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x2a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x61, 0x6e, 0x73, 0x68, 0x65, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x57,
	0x61, 0x6c, 0x6c, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2e, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x2a, 0x1d, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x70, 0x61, 0x72, 0x65,
	0x6e, 0x74, 0x3d, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x2a, 0x2f, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74,
	0x73, 0x2f, 0x2a, 0x7d, 0xda, 0x41, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x42, 0x98, 0x01,
	0x0a, 0x14, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x61, 0x6e, 0x73, 0x68, 0x65, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x42, 0x0f, 0x41, 0x70, 0x69, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x04, 0x2e, 0x3b, 0x70, 0x62, 0x92,
	0x41, 0x66, 0x12, 0x64, 0x0a, 0x13, 0x50, 0x61, 0x6e, 0x73, 0x68, 0x65, 0x65, 0x5f, 0x61, 0x70,
	0x69, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x48, 0x0a, 0x0d, 0x4a, 0x61, 0x6b,
	0x75, 0x62, 0x20, 0x48, 0x75, 0x64, 0x7a, 0x69, 0x61, 0x6b, 0x12, 0x1c, 0x68, 0x74, 0x74, 0x70,
	0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68,
	0x75, 0x64, 0x79, 0x77, 0x65, 0x61, 0x73, 0x2f, 0x1a, 0x19, 0x6a, 0x61, 0x6b, 0x75, 0x62, 0x2e,
	0x6b, 0x2e, 0x68, 0x75, 0x64, 0x7a, 0x69, 0x61, 0x6b, 0x40, 0x67, 0x6d, 0x61, 0x69, 0x6c, 0x2e,
	0x63, 0x6f, 0x6d, 0x32, 0x03, 0x31, 0x2e, 0x30, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_panshee_v1_proto_api_service_proto_goTypes = []interface{}{
	(*CreateUserRequest)(nil),             // 0: api.panshee.v1.proto.CreateUserRequest
	(*LoginUserRequest)(nil),              // 1: api.panshee.v1.proto.LoginUserRequest
	(*GetUserRequest)(nil),                // 2: api.panshee.v1.proto.GetUserRequest
	(*RefreshAccessTokenDTORequest)(nil),  // 3: api.panshee.v1.proto.RefreshAccessTokenDTORequest
	(*CreateWalletRequest)(nil),           // 4: api.panshee.v1.proto.CreateWalletRequest
	(*GetWalletDTODataRequest)(nil),       // 5: api.panshee.v1.proto.GetWalletDTODataRequest
	(*DeleteWalletRequest)(nil),           // 6: api.panshee.v1.proto.DeleteWalletRequest
	(*User)(nil),                          // 7: api.panshee.v1.proto.User
	(*LoginUserResponse)(nil),             // 8: api.panshee.v1.proto.LoginUserResponse
	(*RefreshAccessTokenDTOResponse)(nil), // 9: api.panshee.v1.proto.RefreshAccessTokenDTOResponse
	(*WalletDTO)(nil),                     // 10: api.panshee.v1.proto.WalletDTO
	(*GetWalletDTODataResponse)(nil),      // 11: api.panshee.v1.proto.GetWalletDTODataResponse
	(*DeleteWalletResponse)(nil),          // 12: api.panshee.v1.proto.DeleteWalletResponse
}
var file_panshee_v1_proto_api_service_proto_depIdxs = []int32{
	0,  // 0: api.panshee.v1.proto.ApiService.CreateUser:input_type -> api.panshee.v1.proto.CreateUserRequest
	1,  // 1: api.panshee.v1.proto.ApiService.LoginUser:input_type -> api.panshee.v1.proto.LoginUserRequest
	2,  // 2: api.panshee.v1.proto.ApiService.GetUser:input_type -> api.panshee.v1.proto.GetUserRequest
	3,  // 3: api.panshee.v1.proto.ApiService.RefreshAccessToken:input_type -> api.panshee.v1.proto.RefreshAccessTokenDTORequest
	4,  // 4: api.panshee.v1.proto.ApiService.CreateWallet:input_type -> api.panshee.v1.proto.CreateWalletRequest
	5,  // 5: api.panshee.v1.proto.ApiService.GetWallets:input_type -> api.panshee.v1.proto.GetWalletDTODataRequest
	6,  // 6: api.panshee.v1.proto.ApiService.DeleteWallet:input_type -> api.panshee.v1.proto.DeleteWalletRequest
	7,  // 7: api.panshee.v1.proto.ApiService.CreateUser:output_type -> api.panshee.v1.proto.User
	8,  // 8: api.panshee.v1.proto.ApiService.LoginUser:output_type -> api.panshee.v1.proto.LoginUserResponse
	7,  // 9: api.panshee.v1.proto.ApiService.GetUser:output_type -> api.panshee.v1.proto.User
	9,  // 10: api.panshee.v1.proto.ApiService.RefreshAccessToken:output_type -> api.panshee.v1.proto.RefreshAccessTokenDTOResponse
	10, // 11: api.panshee.v1.proto.ApiService.CreateWallet:output_type -> api.panshee.v1.proto.WalletDTO
	11, // 12: api.panshee.v1.proto.ApiService.GetWallets:output_type -> api.panshee.v1.proto.GetWalletDTODataResponse
	12, // 13: api.panshee.v1.proto.ApiService.DeleteWallet:output_type -> api.panshee.v1.proto.DeleteWalletResponse
	7,  // [7:14] is the sub-list for method output_type
	0,  // [0:7] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_panshee_v1_proto_api_service_proto_init() }
func file_panshee_v1_proto_api_service_proto_init() {
	if File_panshee_v1_proto_api_service_proto != nil {
		return
	}
	file_panshee_v1_proto_api_user_proto_init()
	file_panshee_v1_proto_api_token_proto_init()
	file_panshee_v1_proto_api_wallet_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_panshee_v1_proto_api_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_panshee_v1_proto_api_service_proto_goTypes,
		DependencyIndexes: file_panshee_v1_proto_api_service_proto_depIdxs,
	}.Build()
	File_panshee_v1_proto_api_service_proto = out.File
	file_panshee_v1_proto_api_service_proto_rawDesc = nil
	file_panshee_v1_proto_api_service_proto_goTypes = nil
	file_panshee_v1_proto_api_service_proto_depIdxs = nil
}
