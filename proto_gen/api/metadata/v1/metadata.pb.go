// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/metadata/v1/metadata.proto

package v1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Req struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Req) Reset()         { *m = Req{} }
func (m *Req) String() string { return proto.CompactTextString(m) }
func (*Req) ProtoMessage()    {}
func (*Req) Descriptor() ([]byte, []int) {
	return fileDescriptor_441328c9e07c2e92, []int{0}
}

func (m *Req) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Req.Unmarshal(m, b)
}
func (m *Req) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Req.Marshal(b, m, deterministic)
}
func (m *Req) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Req.Merge(m, src)
}
func (m *Req) XXX_Size() int {
	return xxx_messageInfo_Req.Size(m)
}
func (m *Req) XXX_DiscardUnknown() {
	xxx_messageInfo_Req.DiscardUnknown(m)
}

var xxx_messageInfo_Req proto.InternalMessageInfo

func (m *Req) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *Req) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type Response struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_441328c9e07c2e92, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Req)(nil), "api.metadata.v1.Req")
	proto.RegisterType((*Response)(nil), "api.metadata.v1.Response")
}

func init() { proto.RegisterFile("api/metadata/v1/metadata.proto", fileDescriptor_441328c9e07c2e92) }

var fileDescriptor_441328c9e07c2e92 = []byte{
	// 172 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4b, 0x2c, 0xc8, 0xd4,
	0xcf, 0x4d, 0x2d, 0x49, 0x4c, 0x49, 0x2c, 0x49, 0xd4, 0x2f, 0x33, 0x84, 0xb3, 0xf5, 0x0a, 0x8a,
	0xf2, 0x4b, 0xf2, 0x85, 0xf8, 0x13, 0x0b, 0x32, 0xf5, 0xe0, 0x62, 0x65, 0x86, 0x4a, 0xb6, 0x5c,
	0xcc, 0x41, 0xa9, 0x85, 0x42, 0x52, 0x5c, 0x1c, 0xa5, 0xc5, 0xa9, 0x45, 0x79, 0x89, 0xb9, 0xa9,
	0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x70, 0x3e, 0x48, 0xae, 0x20, 0xb1, 0xb8, 0xb8, 0x3c,
	0xbf, 0x28, 0x45, 0x82, 0x09, 0x22, 0x07, 0xe3, 0x2b, 0x71, 0x71, 0x71, 0x04, 0xa5, 0x16, 0x17,
	0xe4, 0xe7, 0x15, 0xa7, 0x1a, 0xf9, 0x71, 0xf1, 0xfb, 0xa6, 0x96, 0x24, 0x86, 0xa4, 0x16, 0x97,
	0x04, 0xa7, 0x16, 0x95, 0x65, 0x26, 0xa7, 0x0a, 0x59, 0x73, 0xb1, 0x39, 0x17, 0xa5, 0x26, 0x96,
	0xa4, 0x0a, 0x89, 0xe8, 0xa1, 0xd9, 0xac, 0x17, 0x94, 0x5a, 0x28, 0x25, 0x89, 0x45, 0x14, 0x62,
	0x9a, 0x12, 0x83, 0x93, 0x60, 0x14, 0x3f, 0x9a, 0x6f, 0x92, 0xd8, 0xc0, 0xbe, 0x30, 0x06, 0x04,
	0x00, 0x00, 0xff, 0xff, 0xe8, 0x2b, 0xa1, 0xc4, 0xe7, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MetaTestServiceClient is the client API for MetaTestService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MetaTestServiceClient interface {
	Create(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Response, error)
}

type metaTestServiceClient struct {
	cc *grpc.ClientConn
}

func NewMetaTestServiceClient(cc *grpc.ClientConn) MetaTestServiceClient {
	return &metaTestServiceClient{cc}
}

func (c *metaTestServiceClient) Create(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/api.metadata.v1.MetaTestService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MetaTestServiceServer is the server API for MetaTestService service.
type MetaTestServiceServer interface {
	Create(context.Context, *Req) (*Response, error)
}

// UnimplementedMetaTestServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMetaTestServiceServer struct {
}

func (*UnimplementedMetaTestServiceServer) Create(ctx context.Context, req *Req) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}

func RegisterMetaTestServiceServer(s *grpc.Server, srv MetaTestServiceServer) {
	s.RegisterService(&_MetaTestService_serviceDesc, srv)
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
		FullMethod: "/api.metadata.v1.MetaTestService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetaTestServiceServer).Create(ctx, req.(*Req))
	}
	return interceptor(ctx, in, info, handler)
}

var _MetaTestService_serviceDesc = grpc.ServiceDesc{
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