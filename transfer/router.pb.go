// Code generated by protoc-gen-go. DO NOT EDIT.
// source: router.proto

package elastic_transfer

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

type Response struct {
	Error                uint32   `protobuf:"varint,1,opt,name=error,proto3" json:"error,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_367072455c71aedc, []int{0}
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

func (m *Response) GetError() uint32 {
	if m != nil {
		return m.Error
	}
	return 0
}

func (m *Response) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type PushParameter struct {
	Identity             string   `protobuf:"bytes,1,opt,name=identity,proto3" json:"identity,omitempty"`
	Data                 []byte   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PushParameter) Reset()         { *m = PushParameter{} }
func (m *PushParameter) String() string { return proto.CompactTextString(m) }
func (*PushParameter) ProtoMessage()    {}
func (*PushParameter) Descriptor() ([]byte, []int) {
	return fileDescriptor_367072455c71aedc, []int{1}
}

func (m *PushParameter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PushParameter.Unmarshal(m, b)
}
func (m *PushParameter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PushParameter.Marshal(b, m, deterministic)
}
func (m *PushParameter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PushParameter.Merge(m, src)
}
func (m *PushParameter) XXX_Size() int {
	return xxx_messageInfo_PushParameter.Size(m)
}
func (m *PushParameter) XXX_DiscardUnknown() {
	xxx_messageInfo_PushParameter.DiscardUnknown(m)
}

var xxx_messageInfo_PushParameter proto.InternalMessageInfo

func (m *PushParameter) GetIdentity() string {
	if m != nil {
		return m.Identity
	}
	return ""
}

func (m *PushParameter) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*Response)(nil), "elastic.transfer.Response")
	proto.RegisterType((*PushParameter)(nil), "elastic.transfer.PushParameter")
}

func init() { proto.RegisterFile("router.proto", fileDescriptor_367072455c71aedc) }

var fileDescriptor_367072455c71aedc = []byte{
	// 185 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x8f, 0xb1, 0x8a, 0x83, 0x40,
	0x10, 0x86, 0xcf, 0x3b, 0x4f, 0x74, 0x50, 0x90, 0xe1, 0x0a, 0xb1, 0x39, 0xb1, 0xb2, 0xda, 0xc2,
	0x7b, 0x80, 0xab, 0x52, 0x47, 0xf6, 0x0d, 0x36, 0x71, 0x92, 0x08, 0xd1, 0x95, 0xd9, 0xb1, 0xc8,
	0xdb, 0x07, 0x37, 0x24, 0x90, 0xa4, 0xfb, 0x7f, 0x98, 0x6f, 0xe6, 0x1b, 0x48, 0xd9, 0x2e, 0x42,
	0xac, 0x66, 0xb6, 0x62, 0x31, 0xa7, 0xb3, 0x71, 0x32, 0xec, 0x95, 0xb0, 0x99, 0xdc, 0x81, 0xb8,
	0x6e, 0x21, 0xd6, 0xe4, 0x66, 0x3b, 0x39, 0xc2, 0x1f, 0xf8, 0x26, 0x66, 0xcb, 0x45, 0x50, 0x05,
	0x4d, 0xa6, 0x6f, 0x05, 0x73, 0xf8, 0x1a, 0xdd, 0xb1, 0xf8, 0xac, 0x82, 0x26, 0xd1, 0x6b, 0xac,
	0xff, 0x21, 0xeb, 0x16, 0x77, 0xea, 0x0c, 0x9b, 0x91, 0x84, 0x18, 0x4b, 0x88, 0x87, 0x9e, 0x26,
	0x19, 0xe4, 0xe2, 0xd9, 0x44, 0x3f, 0x3a, 0x22, 0x84, 0xbd, 0x11, 0xe3, 0xf9, 0x54, 0xfb, 0xdc,
	0x6e, 0x21, 0xd2, 0x5e, 0x0b, 0x37, 0x10, 0xae, 0xab, 0xf0, 0x57, 0xbd, 0x9a, 0xa9, 0xa7, 0x13,
	0x65, 0xf9, 0x3e, 0x70, 0xf7, 0xae, 0x3f, 0x76, 0x91, 0x7f, 0xef, 0xef, 0x1a, 0x00, 0x00, 0xff,
	0xff, 0xd5, 0x13, 0xba, 0xad, 0xee, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RouterClient is the client API for Router service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RouterClient interface {
	Push(ctx context.Context, in *PushParameter, opts ...grpc.CallOption) (*Response, error)
}

type routerClient struct {
	cc *grpc.ClientConn
}

func NewRouterClient(cc *grpc.ClientConn) RouterClient {
	return &routerClient{cc}
}

func (c *routerClient) Push(ctx context.Context, in *PushParameter, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/elastic.transfer.Router/Push", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RouterServer is the server API for Router service.
type RouterServer interface {
	Push(context.Context, *PushParameter) (*Response, error)
}

// UnimplementedRouterServer can be embedded to have forward compatible implementations.
type UnimplementedRouterServer struct {
}

func (*UnimplementedRouterServer) Push(ctx context.Context, req *PushParameter) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Push not implemented")
}

func RegisterRouterServer(s *grpc.Server, srv RouterServer) {
	s.RegisterService(&_Router_serviceDesc, srv)
}

func _Router_Push_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PushParameter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouterServer).Push(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/elastic.transfer.Router/Push",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouterServer).Push(ctx, req.(*PushParameter))
	}
	return interceptor(ctx, in, info, handler)
}

var _Router_serviceDesc = grpc.ServiceDesc{
	ServiceName: "elastic.transfer.Router",
	HandlerType: (*RouterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Push",
			Handler:    _Router_Push_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "router.proto",
}