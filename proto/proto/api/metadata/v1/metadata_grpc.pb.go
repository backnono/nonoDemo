// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: api/metadata/v1/metadata.proto

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
	MetaTestService_Create_FullMethodName = "/api.metadata.v1.MetaTestService/Create"
)

// MetaTestServiceClient is the client API for MetaTestService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MetaTestServiceClient interface {
	Create(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Response, error)
}

type metaTestServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMetaTestServiceClient(cc grpc.ClientConnInterface) MetaTestServiceClient {
	return &metaTestServiceClient{cc}
}

func (c *metaTestServiceClient) Create(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, MetaTestService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MetaTestServiceServer is the server API for MetaTestService service.
// All implementations should embed UnimplementedMetaTestServiceServer
// for forward compatibility
type MetaTestServiceServer interface {
	Create(context.Context, *Req) (*Response, error)
}

// UnimplementedMetaTestServiceServer should be embedded to have forward compatible implementations.
type UnimplementedMetaTestServiceServer struct {
}

func (UnimplementedMetaTestServiceServer) Create(context.Context, *Req) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}

// UnsafeMetaTestServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MetaTestServiceServer will
// result in compilation errors.
type UnsafeMetaTestServiceServer interface {
	mustEmbedUnimplementedMetaTestServiceServer()
}

func RegisterMetaTestServiceServer(s grpc.ServiceRegistrar, srv MetaTestServiceServer) {
	s.RegisterService(&MetaTestService_ServiceDesc, srv)
}

func _MetaTestService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Req)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetaTestServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MetaTestService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetaTestServiceServer).Create(ctx, req.(*Req))
	}
	return interceptor(ctx, in, info, handler)
}

// MetaTestService_ServiceDesc is the grpc.ServiceDesc for MetaTestService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MetaTestService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.metadata.v1.MetaTestService",
	HandlerType: (*MetaTestServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _MetaTestService_Create_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/metadata/v1/metadata.proto",
}