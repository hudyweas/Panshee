// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0-devel
// 	protoc        v3.21.6
// source: panshee/v1/proto/wallet.proto

package pb

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

type GetWalletDataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WalletAddress string `protobuf:"bytes,1,opt,name=walletAddress,proto3" json:"walletAddress,omitempty"`
}

func (x *GetWalletDataRequest) Reset() {
	*x = GetWalletDataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_panshee_v1_proto_wallet_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetWalletDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWalletDataRequest) ProtoMessage() {}

func (x *GetWalletDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_panshee_v1_proto_wallet_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWalletDataRequest.ProtoReflect.Descriptor instead.
func (*GetWalletDataRequest) Descriptor() ([]byte, []int) {
	return file_panshee_v1_proto_wallet_proto_rawDescGZIP(), []int{0}
}

func (x *GetWalletDataRequest) GetWalletAddress() string {
	if x != nil {
		return x.WalletAddress
	}
	return ""
}

type GetWalletDataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetWalletDataResponse) Reset() {
	*x = GetWalletDataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_panshee_v1_proto_wallet_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetWalletDataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWalletDataResponse) ProtoMessage() {}

func (x *GetWalletDataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_panshee_v1_proto_wallet_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWalletDataResponse.ProtoReflect.Descriptor instead.
func (*GetWalletDataResponse) Descriptor() ([]byte, []int) {
	return file_panshee_v1_proto_wallet_proto_rawDescGZIP(), []int{1}
}

var File_panshee_v1_proto_wallet_proto protoreflect.FileDescriptor

var file_panshee_v1_proto_wallet_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x70, 0x61, 0x6e, 0x73, 0x68, 0x65, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x14, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x61, 0x6e, 0x73, 0x68, 0x65, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3c, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x57, 0x61, 0x6c, 0x6c,
	0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a,
	0x0d, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x22, 0x17, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74,
	0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x2b, 0x0a, 0x14,
	0x61, 0x70, 0x69, 0x2e, 0x70, 0x61, 0x6e, 0x73, 0x68, 0x65, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x42, 0x0b, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x04, 0x2e, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_panshee_v1_proto_wallet_proto_rawDescOnce sync.Once
	file_panshee_v1_proto_wallet_proto_rawDescData = file_panshee_v1_proto_wallet_proto_rawDesc
)

func file_panshee_v1_proto_wallet_proto_rawDescGZIP() []byte {
	file_panshee_v1_proto_wallet_proto_rawDescOnce.Do(func() {
		file_panshee_v1_proto_wallet_proto_rawDescData = protoimpl.X.CompressGZIP(file_panshee_v1_proto_wallet_proto_rawDescData)
	})
	return file_panshee_v1_proto_wallet_proto_rawDescData
}

var file_panshee_v1_proto_wallet_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_panshee_v1_proto_wallet_proto_goTypes = []interface{}{
	(*GetWalletDataRequest)(nil),  // 0: api.panshee.v1.proto.GetWalletDataRequest
	(*GetWalletDataResponse)(nil), // 1: api.panshee.v1.proto.GetWalletDataResponse
}
var file_panshee_v1_proto_wallet_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_panshee_v1_proto_wallet_proto_init() }
func file_panshee_v1_proto_wallet_proto_init() {
	if File_panshee_v1_proto_wallet_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_panshee_v1_proto_wallet_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetWalletDataRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_panshee_v1_proto_wallet_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetWalletDataResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_panshee_v1_proto_wallet_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_panshee_v1_proto_wallet_proto_goTypes,
		DependencyIndexes: file_panshee_v1_proto_wallet_proto_depIdxs,
		MessageInfos:      file_panshee_v1_proto_wallet_proto_msgTypes,
	}.Build()
	File_panshee_v1_proto_wallet_proto = out.File
	file_panshee_v1_proto_wallet_proto_rawDesc = nil
	file_panshee_v1_proto_wallet_proto_goTypes = nil
	file_panshee_v1_proto_wallet_proto_depIdxs = nil
}
