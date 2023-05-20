// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: proto/customers.proto

package proto

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
	CustomersService_Create_FullMethodName = "/proto.CustomersService/Create"
	CustomersService_Get_FullMethodName    = "/proto.CustomersService/Get"
	CustomersService_Update_FullMethodName = "/proto.CustomersService/Update"
)

// CustomersServiceClient is the client API for CustomersService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CustomersServiceClient interface {
	Create(ctx context.Context, in *CustomerCreateRequest, opts ...grpc.CallOption) (*CustomerResponse, error)
	Get(ctx context.Context, in *CustomerGetRequest, opts ...grpc.CallOption) (*CustomerGetResponse, error)
	Update(ctx context.Context, in *CustomerUpdateRequest, opts ...grpc.CallOption) (*CustomerResponse, error)
}

type customersServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCustomersServiceClient(cc grpc.ClientConnInterface) CustomersServiceClient {
	return &customersServiceClient{cc}
}

func (c *customersServiceClient) Create(ctx context.Context, in *CustomerCreateRequest, opts ...grpc.CallOption) (*CustomerResponse, error) {
	out := new(CustomerResponse)
	err := c.cc.Invoke(ctx, CustomersService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customersServiceClient) Get(ctx context.Context, in *CustomerGetRequest, opts ...grpc.CallOption) (*CustomerGetResponse, error) {
	out := new(CustomerGetResponse)
	err := c.cc.Invoke(ctx, CustomersService_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customersServiceClient) Update(ctx context.Context, in *CustomerUpdateRequest, opts ...grpc.CallOption) (*CustomerResponse, error) {
	out := new(CustomerResponse)
	err := c.cc.Invoke(ctx, CustomersService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustomersServiceServer is the server API for CustomersService service.
// All implementations should embed UnimplementedCustomersServiceServer
// for forward compatibility
type CustomersServiceServer interface {
	Create(context.Context, *CustomerCreateRequest) (*CustomerResponse, error)
	Get(context.Context, *CustomerGetRequest) (*CustomerGetResponse, error)
	Update(context.Context, *CustomerUpdateRequest) (*CustomerResponse, error)
}

// UnimplementedCustomersServiceServer should be embedded to have forward compatible implementations.
type UnimplementedCustomersServiceServer struct {
}

func (UnimplementedCustomersServiceServer) Create(context.Context, *CustomerCreateRequest) (*CustomerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedCustomersServiceServer) Get(context.Context, *CustomerGetRequest) (*CustomerGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedCustomersServiceServer) Update(context.Context, *CustomerUpdateRequest) (*CustomerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}

// UnsafeCustomersServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CustomersServiceServer will
// result in compilation errors.
type UnsafeCustomersServiceServer interface {
	mustEmbedUnimplementedCustomersServiceServer()
}

func RegisterCustomersServiceServer(s grpc.ServiceRegistrar, srv CustomersServiceServer) {
	s.RegisterService(&CustomersService_ServiceDesc, srv)
}

func _CustomersService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CustomerCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomersServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CustomersService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomersServiceServer).Create(ctx, req.(*CustomerCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomersService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CustomerGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomersServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CustomersService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomersServiceServer).Get(ctx, req.(*CustomerGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomersService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CustomerUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomersServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CustomersService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomersServiceServer).Update(ctx, req.(*CustomerUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CustomersService_ServiceDesc is the grpc.ServiceDesc for CustomersService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CustomersService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.CustomersService",
	HandlerType: (*CustomersServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _CustomersService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _CustomersService_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _CustomersService_Update_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/customers.proto",
}
