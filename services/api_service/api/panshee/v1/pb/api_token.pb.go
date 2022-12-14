// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0-devel
// 	protoc        v3.21.6
// source: panshee/v1/proto/api_token.proto

package pb

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Structure of a token
type TokenDTO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Token string
	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	// Expiration time of the token
	ExpirationTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=expiration_time,json=expirationTime,proto3" json:"expiration_time,omitempty"`
}

func (x *TokenDTO) Reset() {
	*x = TokenDTO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_panshee_v1_proto_api_token_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenDTO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenDTO) ProtoMessage() {}

func (x *TokenDTO) ProtoReflect() protoreflect.Message {
	mi := &file_panshee_v1_proto_api_token_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenDTO.ProtoReflect.Descriptor instead.
func (*TokenDTO) Descriptor() ([]byte, []int) {
	return file_panshee_v1_proto_api_token_proto_rawDescGZIP(), []int{0}
}

func (x *TokenDTO) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *TokenDTO) GetExpirationTime() *timestamppb.Timestamp {
	if x != nil {
		return x.ExpirationTime
	}
	return nil
}

// A representation of a request for refreshing access token
type RefreshAccessTokenDTORequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id of the user to who wants to refresh token
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// Refresh token used to refresh access token
	RefreshToken *TokenDTO `protobuf:"bytes,2,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
}

func (x *RefreshAccessTokenDTORequest) Reset() {
	*x = RefreshAccessTokenDTORequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_panshee_v1_proto_api_token_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RefreshAccessTokenDTORequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefreshAccessTokenDTORequest) ProtoMessage() {}

func (x *RefreshAccessTokenDTORequest) ProtoReflect() protoreflect.Message {
	mi := &file_panshee_v1_proto_api_token_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefreshAccessTokenDTORequest.ProtoReflect.Descriptor instead.
func (*RefreshAccessTokenDTORequest) Descriptor() ([]byte, []int) {
	return file_panshee_v1_proto_api_token_proto_rawDescGZIP(), []int{1}
}

func (x *RefreshAccessTokenDTORequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *RefreshAccessTokenDTORequest) GetRefreshToken() *TokenDTO {
	if x != nil {
		return x.RefreshToken
	}
	return nil
}

// Response with new access_token
type RefreshAccessTokenDTOResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Access token returning to user
	AccessToken *TokenDTO `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
}

func (x *RefreshAccessTokenDTOResponse) Reset() {
	*x = RefreshAccessTokenDTOResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_panshee_v1_proto_api_token_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RefreshAccessTokenDTOResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefreshAccessTokenDTOResponse) ProtoMessage() {}

func (x *RefreshAccessTokenDTOResponse) ProtoReflect() protoreflect.Message {
	mi := &file_panshee_v1_proto_api_token_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefreshAccessTokenDTOResponse.ProtoReflect.Descriptor instead.
func (*RefreshAccessTokenDTOResponse) Descriptor() ([]byte, []int) {
	return file_panshee_v1_proto_api_token_proto_rawDescGZIP(), []int{2}
}

func (x *RefreshAccessTokenDTOResponse) GetAccessToken() *TokenDTO {
	if x != nil {
		return x.AccessToken
	}
	return nil
}

var File_panshee_v1_proto_api_token_proto protoreflect.FileDescriptor

var file_panshee_v1_proto_api_token_proto_rawDesc = []byte{
	0x0a, 0x20, 0x70, 0x61, 0x6e, 0x73, 0x68, 0x65, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x14, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x61, 0x6e, 0x73, 0x68, 0x65, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76,
	0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x65, 0x0a, 0x08, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x44, 0x54, 0x4f, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x43, 0x0a, 0x0f,
	0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x0e, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d,
	0x65, 0x22, 0x85, 0x01, 0x0a, 0x1c, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x41, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x44, 0x54, 0x4f, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1b, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12,
	0x48, 0x0a, 0x0d, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x61, 0x6e,
	0x73, 0x68, 0x65, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x44, 0x54, 0x4f, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0c, 0x72, 0x65, 0x66,
	0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x62, 0x0a, 0x1d, 0x52, 0x65, 0x66,
	0x72, 0x65, 0x73, 0x68, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x44,
	0x54, 0x4f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x0c, 0x61, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x61, 0x6e, 0x73, 0x68, 0x65, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x44, 0x54, 0x4f,
	0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x2a, 0x0a,
	0x14, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x61, 0x6e, 0x73, 0x68, 0x65, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x42, 0x0a, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x04, 0x2e, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_panshee_v1_proto_api_token_proto_rawDescOnce sync.Once
	file_panshee_v1_proto_api_token_proto_rawDescData = file_panshee_v1_proto_api_token_proto_rawDesc
)

func file_panshee_v1_proto_api_token_proto_rawDescGZIP() []byte {
	file_panshee_v1_proto_api_token_proto_rawDescOnce.Do(func() {
		file_panshee_v1_proto_api_token_proto_rawDescData = protoimpl.X.CompressGZIP(file_panshee_v1_proto_api_token_proto_rawDescData)
	})
	return file_panshee_v1_proto_api_token_proto_rawDescData
}

var file_panshee_v1_proto_api_token_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_panshee_v1_proto_api_token_proto_goTypes = []interface{}{
	(*TokenDTO)(nil),                      // 0: api.panshee.v1.proto.TokenDTO
	(*RefreshAccessTokenDTORequest)(nil),  // 1: api.panshee.v1.proto.RefreshAccessTokenDTORequest
	(*RefreshAccessTokenDTOResponse)(nil), // 2: api.panshee.v1.proto.RefreshAccessTokenDTOResponse
	(*timestamppb.Timestamp)(nil),         // 3: google.protobuf.Timestamp
}
var file_panshee_v1_proto_api_token_proto_depIdxs = []int32{
	3, // 0: api.panshee.v1.proto.TokenDTO.expiration_time:type_name -> google.protobuf.Timestamp
	0, // 1: api.panshee.v1.proto.RefreshAccessTokenDTORequest.refresh_token:type_name -> api.panshee.v1.proto.TokenDTO
	0, // 2: api.panshee.v1.proto.RefreshAccessTokenDTOResponse.access_token:type_name -> api.panshee.v1.proto.TokenDTO
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_panshee_v1_proto_api_token_proto_init() }
func file_panshee_v1_proto_api_token_proto_init() {
	if File_panshee_v1_proto_api_token_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_panshee_v1_proto_api_token_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenDTO); i {
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
		file_panshee_v1_proto_api_token_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RefreshAccessTokenDTORequest); i {
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
		file_panshee_v1_proto_api_token_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RefreshAccessTokenDTOResponse); i {
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
			RawDescriptor: file_panshee_v1_proto_api_token_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_panshee_v1_proto_api_token_proto_goTypes,
		DependencyIndexes: file_panshee_v1_proto_api_token_proto_depIdxs,
		MessageInfos:      file_panshee_v1_proto_api_token_proto_msgTypes,
	}.Build()
	File_panshee_v1_proto_api_token_proto = out.File
	file_panshee_v1_proto_api_token_proto_rawDesc = nil
	file_panshee_v1_proto_api_token_proto_goTypes = nil
	file_panshee_v1_proto_api_token_proto_depIdxs = nil
}
