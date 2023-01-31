// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: wallet_service.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// PansheeWalletServiceClient is the client API for PansheeWalletService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PansheeWalletServiceClient interface {
	GetWalletData(ctx context.Context, in *GetWalletDataRequest, opts ...grpc.CallOption) (*Wallet, error)
}

type pansheeWalletServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPansheeWalletServiceClient(cc grpc.ClientConnInterface) PansheeWalletServiceClient {
	return &pansheeWalletServiceClient{cc}
}

func (c *pansheeWalletServiceClient) GetWalletData(ctx context.Context, in *GetWalletDataRequest, opts ...grpc.CallOption) (*Wallet, error) {
	out := new(Wallet)
	err := c.cc.Invoke(ctx, "/api.panshee.v1.proto.PansheeWalletService/GetWalletData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PansheeWalletServiceServer is the server API for PansheeWalletService service.
// All implementations must embed UnimplementedPansheeWalletServiceServer
// for forward compatibility
type PansheeWalletServiceServer interface {
	GetWalletData(context.Context, *GetWalletDataRequest) (*Wallet, error)
	mustEmbedUnimplementedPansheeWalletServiceServer()
}

// UnimplementedPansheeWalletServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPansheeWalletServiceServer struct {
}

func (UnimplementedPansheeWalletServiceServer) GetWalletData(context.Context, *GetWalletDataRequest) (*Wallet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWalletData not implemented")
}
func (UnimplementedPansheeWalletServiceServer) mustEmbedUnimplementedPansheeWalletServiceServer() {}

// UnsafePansheeWalletServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PansheeWalletServiceServer will
// result in compilation errors.
type UnsafePansheeWalletServiceServer interface {
	mustEmbedUnimplementedPansheeWalletServiceServer()
}

func RegisterPansheeWalletServiceServer(s grpc.ServiceRegistrar, srv PansheeWalletServiceServer) {
	s.RegisterService(&PansheeWalletService_ServiceDesc, srv)
}

func _PansheeWalletService_GetWalletData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWalletDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PansheeWalletServiceServer).GetWalletData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.panshee.v1.proto.PansheeWalletService/GetWalletData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PansheeWalletServiceServer).GetWalletData(ctx, req.(*GetWalletDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PansheeWalletService_ServiceDesc is the grpc.ServiceDesc for PansheeWalletService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PansheeWalletService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.panshee.v1.proto.PansheeWalletService",
	HandlerType: (*PansheeWalletServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetWalletData",
			Handler:    _PansheeWalletService_GetWalletData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "wallet_service.proto",
}