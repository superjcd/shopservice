// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: shopservice.proto

package v1

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

const (
	ShopService_CreateShop_FullMethodName = "/proto.v1.ShopService/CreateShop"
	ShopService_ListShop_FullMethodName   = "/proto.v1.ShopService/ListShop"
	ShopService_UpdateShop_FullMethodName = "/proto.v1.ShopService/UpdateShop"
	ShopService_DeleteShop_FullMethodName = "/proto.v1.ShopService/DeleteShop"
)

// ShopServiceClient is the client API for ShopService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ShopServiceClient interface {
	CreateShop(ctx context.Context, in *CreateShopRequest, opts ...grpc.CallOption) (*CreateShopResponse, error)
	ListShop(ctx context.Context, in *ListShopRequest, opts ...grpc.CallOption) (*ListShopResponse, error)
	UpdateShop(ctx context.Context, in *UpdateShopRequest, opts ...grpc.CallOption) (*UpdateShopResponse, error)
	DeleteShop(ctx context.Context, in *DeleteShopRequest, opts ...grpc.CallOption) (*DeleteShopResponse, error)
}

type shopServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewShopServiceClient(cc grpc.ClientConnInterface) ShopServiceClient {
	return &shopServiceClient{cc}
}

func (c *shopServiceClient) CreateShop(ctx context.Context, in *CreateShopRequest, opts ...grpc.CallOption) (*CreateShopResponse, error) {
	out := new(CreateShopResponse)
	err := c.cc.Invoke(ctx, ShopService_CreateShop_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopServiceClient) ListShop(ctx context.Context, in *ListShopRequest, opts ...grpc.CallOption) (*ListShopResponse, error) {
	out := new(ListShopResponse)
	err := c.cc.Invoke(ctx, ShopService_ListShop_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopServiceClient) UpdateShop(ctx context.Context, in *UpdateShopRequest, opts ...grpc.CallOption) (*UpdateShopResponse, error) {
	out := new(UpdateShopResponse)
	err := c.cc.Invoke(ctx, ShopService_UpdateShop_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopServiceClient) DeleteShop(ctx context.Context, in *DeleteShopRequest, opts ...grpc.CallOption) (*DeleteShopResponse, error) {
	out := new(DeleteShopResponse)
	err := c.cc.Invoke(ctx, ShopService_DeleteShop_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShopServiceServer is the server API for ShopService service.
// All implementations must embed UnimplementedShopServiceServer
// for forward compatibility
type ShopServiceServer interface {
	CreateShop(context.Context, *CreateShopRequest) (*CreateShopResponse, error)
	ListShop(context.Context, *ListShopRequest) (*ListShopResponse, error)
	UpdateShop(context.Context, *UpdateShopRequest) (*UpdateShopResponse, error)
	DeleteShop(context.Context, *DeleteShopRequest) (*DeleteShopResponse, error)
	mustEmbedUnimplementedShopServiceServer()
}

// UnimplementedShopServiceServer must be embedded to have forward compatible implementations.
type UnimplementedShopServiceServer struct {
}

func (UnimplementedShopServiceServer) CreateShop(context.Context, *CreateShopRequest) (*CreateShopResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateShop not implemented")
}
func (UnimplementedShopServiceServer) ListShop(context.Context, *ListShopRequest) (*ListShopResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListShop not implemented")
}
func (UnimplementedShopServiceServer) UpdateShop(context.Context, *UpdateShopRequest) (*UpdateShopResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateShop not implemented")
}
func (UnimplementedShopServiceServer) DeleteShop(context.Context, *DeleteShopRequest) (*DeleteShopResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteShop not implemented")
}
func (UnimplementedShopServiceServer) mustEmbedUnimplementedShopServiceServer() {}

// UnsafeShopServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ShopServiceServer will
// result in compilation errors.
type UnsafeShopServiceServer interface {
	mustEmbedUnimplementedShopServiceServer()
}

func RegisterShopServiceServer(s grpc.ServiceRegistrar, srv ShopServiceServer) {
	s.RegisterService(&ShopService_ServiceDesc, srv)
}

func _ShopService_CreateShop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateShopRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopServiceServer).CreateShop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShopService_CreateShop_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopServiceServer).CreateShop(ctx, req.(*CreateShopRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShopService_ListShop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListShopRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopServiceServer).ListShop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShopService_ListShop_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopServiceServer).ListShop(ctx, req.(*ListShopRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShopService_UpdateShop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateShopRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopServiceServer).UpdateShop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShopService_UpdateShop_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopServiceServer).UpdateShop(ctx, req.(*UpdateShopRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShopService_DeleteShop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteShopRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopServiceServer).DeleteShop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShopService_DeleteShop_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopServiceServer).DeleteShop(ctx, req.(*DeleteShopRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ShopService_ServiceDesc is the grpc.ServiceDesc for ShopService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ShopService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.v1.ShopService",
	HandlerType: (*ShopServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateShop",
			Handler:    _ShopService_CreateShop_Handler,
		},
		{
			MethodName: "ListShop",
			Handler:    _ShopService_ListShop_Handler,
		},
		{
			MethodName: "UpdateShop",
			Handler:    _ShopService_UpdateShop_Handler,
		},
		{
			MethodName: "DeleteShop",
			Handler:    _ShopService_DeleteShop_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "shopservice.proto",
}
