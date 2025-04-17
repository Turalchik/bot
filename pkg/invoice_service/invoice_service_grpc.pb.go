// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: invoice_service/invoice_service.proto

package invoice_service

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	InvoiceService_CreateInvoice_FullMethodName   = "/turalchik.invoice_service.v1.InvoiceService/CreateInvoice"
	InvoiceService_DescribeInvoice_FullMethodName = "/turalchik.invoice_service.v1.InvoiceService/DescribeInvoice"
	InvoiceService_ListInvoice_FullMethodName     = "/turalchik.invoice_service.v1.InvoiceService/ListInvoice"
	InvoiceService_RemoveInvoice_FullMethodName   = "/turalchik.invoice_service.v1.InvoiceService/RemoveInvoice"
)

// InvoiceServiceClient is the client API for InvoiceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InvoiceServiceClient interface {
	CreateInvoice(ctx context.Context, in *CreateInvoiceRequest, opts ...grpc.CallOption) (*CreateInvoiceResponse, error)
	DescribeInvoice(ctx context.Context, in *DescribeInvoiceRequest, opts ...grpc.CallOption) (*DescribeInvoiceResponse, error)
	ListInvoice(ctx context.Context, in *ListInvoiceRequest, opts ...grpc.CallOption) (*ListInvoiceResponse, error)
	RemoveInvoice(ctx context.Context, in *RemoveInvoiceRequest, opts ...grpc.CallOption) (*RemoveInvoiceResponse, error)
}

type invoiceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewInvoiceServiceClient(cc grpc.ClientConnInterface) InvoiceServiceClient {
	return &invoiceServiceClient{cc}
}

func (c *invoiceServiceClient) CreateInvoice(ctx context.Context, in *CreateInvoiceRequest, opts ...grpc.CallOption) (*CreateInvoiceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateInvoiceResponse)
	err := c.cc.Invoke(ctx, InvoiceService_CreateInvoice_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *invoiceServiceClient) DescribeInvoice(ctx context.Context, in *DescribeInvoiceRequest, opts ...grpc.CallOption) (*DescribeInvoiceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DescribeInvoiceResponse)
	err := c.cc.Invoke(ctx, InvoiceService_DescribeInvoice_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *invoiceServiceClient) ListInvoice(ctx context.Context, in *ListInvoiceRequest, opts ...grpc.CallOption) (*ListInvoiceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListInvoiceResponse)
	err := c.cc.Invoke(ctx, InvoiceService_ListInvoice_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *invoiceServiceClient) RemoveInvoice(ctx context.Context, in *RemoveInvoiceRequest, opts ...grpc.CallOption) (*RemoveInvoiceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RemoveInvoiceResponse)
	err := c.cc.Invoke(ctx, InvoiceService_RemoveInvoice_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InvoiceServiceServer is the server API for InvoiceService service.
// All implementations must embed UnimplementedInvoiceServiceServer
// for forward compatibility.
type InvoiceServiceServer interface {
	CreateInvoice(context.Context, *CreateInvoiceRequest) (*CreateInvoiceResponse, error)
	DescribeInvoice(context.Context, *DescribeInvoiceRequest) (*DescribeInvoiceResponse, error)
	ListInvoice(context.Context, *ListInvoiceRequest) (*ListInvoiceResponse, error)
	RemoveInvoice(context.Context, *RemoveInvoiceRequest) (*RemoveInvoiceResponse, error)
	mustEmbedUnimplementedInvoiceServiceServer()
}

// UnimplementedInvoiceServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedInvoiceServiceServer struct{}

func (UnimplementedInvoiceServiceServer) CreateInvoice(context.Context, *CreateInvoiceRequest) (*CreateInvoiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateInvoice not implemented")
}
func (UnimplementedInvoiceServiceServer) DescribeInvoice(context.Context, *DescribeInvoiceRequest) (*DescribeInvoiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeInvoice not implemented")
}
func (UnimplementedInvoiceServiceServer) ListInvoice(context.Context, *ListInvoiceRequest) (*ListInvoiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListInvoice not implemented")
}
func (UnimplementedInvoiceServiceServer) RemoveInvoice(context.Context, *RemoveInvoiceRequest) (*RemoveInvoiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveInvoice not implemented")
}
func (UnimplementedInvoiceServiceServer) mustEmbedUnimplementedInvoiceServiceServer() {}
func (UnimplementedInvoiceServiceServer) testEmbeddedByValue()                        {}

// UnsafeInvoiceServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InvoiceServiceServer will
// result in compilation errors.
type UnsafeInvoiceServiceServer interface {
	mustEmbedUnimplementedInvoiceServiceServer()
}

func RegisterInvoiceServiceServer(s grpc.ServiceRegistrar, srv InvoiceServiceServer) {
	// If the following call pancis, it indicates UnimplementedInvoiceServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&InvoiceService_ServiceDesc, srv)
}

func _InvoiceService_CreateInvoice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateInvoiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InvoiceServiceServer).CreateInvoice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InvoiceService_CreateInvoice_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InvoiceServiceServer).CreateInvoice(ctx, req.(*CreateInvoiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InvoiceService_DescribeInvoice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeInvoiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InvoiceServiceServer).DescribeInvoice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InvoiceService_DescribeInvoice_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InvoiceServiceServer).DescribeInvoice(ctx, req.(*DescribeInvoiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InvoiceService_ListInvoice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListInvoiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InvoiceServiceServer).ListInvoice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InvoiceService_ListInvoice_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InvoiceServiceServer).ListInvoice(ctx, req.(*ListInvoiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InvoiceService_RemoveInvoice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveInvoiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InvoiceServiceServer).RemoveInvoice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InvoiceService_RemoveInvoice_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InvoiceServiceServer).RemoveInvoice(ctx, req.(*RemoveInvoiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// InvoiceService_ServiceDesc is the grpc.ServiceDesc for InvoiceService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var InvoiceService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "turalchik.invoice_service.v1.InvoiceService",
	HandlerType: (*InvoiceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateInvoice",
			Handler:    _InvoiceService_CreateInvoice_Handler,
		},
		{
			MethodName: "DescribeInvoice",
			Handler:    _InvoiceService_DescribeInvoice_Handler,
		},
		{
			MethodName: "ListInvoice",
			Handler:    _InvoiceService_ListInvoice_Handler,
		},
		{
			MethodName: "RemoveInvoice",
			Handler:    _InvoiceService_RemoveInvoice_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "invoice_service/invoice_service.proto",
}
