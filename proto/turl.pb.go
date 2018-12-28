// Code generated by protoc-gen-go. DO NOT EDIT.
// source: turl.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
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

type URL struct {
	Url                  string   `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Hash                 string   `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *URL) Reset()         { *m = URL{} }
func (m *URL) String() string { return proto.CompactTextString(m) }
func (*URL) ProtoMessage()    {}
func (*URL) Descriptor() ([]byte, []int) {
	return fileDescriptor_7838909e0ddbd4cd, []int{0}
}

func (m *URL) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_URL.Unmarshal(m, b)
}
func (m *URL) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_URL.Marshal(b, m, deterministic)
}
func (m *URL) XXX_Merge(src proto.Message) {
	xxx_messageInfo_URL.Merge(m, src)
}
func (m *URL) XXX_Size() int {
	return xxx_messageInfo_URL.Size(m)
}
func (m *URL) XXX_DiscardUnknown() {
	xxx_messageInfo_URL.DiscardUnknown(m)
}

var xxx_messageInfo_URL proto.InternalMessageInfo

func (m *URL) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *URL) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*URL)(nil), "proto.URL")
}

func init() { proto.RegisterFile("turl.proto", fileDescriptor_7838909e0ddbd4cd) }

var fileDescriptor_7838909e0ddbd4cd = []byte{
	// 187 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x29, 0x2d, 0xca,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x52, 0x32, 0xe9, 0xf9, 0xf9, 0xe9,
	0x39, 0xa9, 0xfa, 0x89, 0x05, 0x99, 0xfa, 0x89, 0x79, 0x79, 0xf9, 0x25, 0x89, 0x25, 0x99, 0xf9,
	0x79, 0xc5, 0x10, 0x45, 0x4a, 0xda, 0x5c, 0xcc, 0xa1, 0x41, 0x3e, 0x42, 0x02, 0x5c, 0xcc, 0xa5,
	0x45, 0x39, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x20, 0xa6, 0x90, 0x10, 0x17, 0x4b, 0x46,
	0x62, 0x71, 0x86, 0x04, 0x13, 0x58, 0x08, 0xcc, 0x36, 0xaa, 0xe5, 0xe2, 0x0e, 0x29, 0x2d, 0xca,
	0x09, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0x15, 0xb2, 0xe4, 0x62, 0xf3, 0x4b, 0x2d, 0x07, 0x69,
	0xe7, 0x82, 0x98, 0xa6, 0x17, 0x1a, 0xe4, 0x23, 0x85, 0xc4, 0x56, 0x12, 0x6f, 0xba, 0xfc, 0x64,
	0x32, 0x93, 0xa0, 0x12, 0x8f, 0x7e, 0x99, 0xa1, 0x3e, 0xc8, 0x69, 0xfa, 0x79, 0xa9, 0xe5, 0x56,
	0x8c, 0x5a, 0x42, 0xc6, 0x5c, 0x6c, 0xee, 0xa9, 0x25, 0xf8, 0xb4, 0x0a, 0x82, 0xb5, 0x72, 0x0b,
	0x71, 0xea, 0x97, 0xe8, 0x57, 0x83, 0x6c, 0xaf, 0x4d, 0x62, 0x03, 0xcb, 0x1a, 0x03, 0x02, 0x00,
	0x00, 0xff, 0xff, 0x7f, 0x2b, 0x49, 0xe7, 0xe5, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TurlServiceClient is the client API for TurlService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TurlServiceClient interface {
	NewURL(ctx context.Context, in *URL, opts ...grpc.CallOption) (*URL, error)
	GetURL(ctx context.Context, in *URL, opts ...grpc.CallOption) (*URL, error)
}

type turlServiceClient struct {
	cc *grpc.ClientConn
}

func NewTurlServiceClient(cc *grpc.ClientConn) TurlServiceClient {
	return &turlServiceClient{cc}
}

func (c *turlServiceClient) NewURL(ctx context.Context, in *URL, opts ...grpc.CallOption) (*URL, error) {
	out := new(URL)
	err := c.cc.Invoke(ctx, "/proto.TurlService/NewURL", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *turlServiceClient) GetURL(ctx context.Context, in *URL, opts ...grpc.CallOption) (*URL, error) {
	out := new(URL)
	err := c.cc.Invoke(ctx, "/proto.TurlService/GetURL", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TurlServiceServer is the server API for TurlService service.
type TurlServiceServer interface {
	NewURL(context.Context, *URL) (*URL, error)
	GetURL(context.Context, *URL) (*URL, error)
}

func RegisterTurlServiceServer(s *grpc.Server, srv TurlServiceServer) {
	s.RegisterService(&_TurlService_serviceDesc, srv)
}

func _TurlService_NewURL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(URL)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TurlServiceServer).NewURL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.TurlService/NewURL",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TurlServiceServer).NewURL(ctx, req.(*URL))
	}
	return interceptor(ctx, in, info, handler)
}

func _TurlService_GetURL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(URL)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TurlServiceServer).GetURL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.TurlService/GetURL",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TurlServiceServer).GetURL(ctx, req.(*URL))
	}
	return interceptor(ctx, in, info, handler)
}

var _TurlService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.TurlService",
	HandlerType: (*TurlServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NewURL",
			Handler:    _TurlService_NewURL_Handler,
		},
		{
			MethodName: "GetURL",
			Handler:    _TurlService_GetURL_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "turl.proto",
}