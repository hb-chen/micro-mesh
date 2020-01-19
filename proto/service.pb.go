// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/service.proto

package proto

import (
	context "context"
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type ApiRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Services             string   `protobuf:"bytes,2,opt,name=services,proto3" json:"services,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ApiRequest) Reset()         { *m = ApiRequest{} }
func (m *ApiRequest) String() string { return proto.CompactTextString(m) }
func (*ApiRequest) ProtoMessage()    {}
func (*ApiRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c33392ef2c1961ba, []int{0}
}

func (m *ApiRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApiRequest.Unmarshal(m, b)
}
func (m *ApiRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApiRequest.Marshal(b, m, deterministic)
}
func (m *ApiRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApiRequest.Merge(m, src)
}
func (m *ApiRequest) XXX_Size() int {
	return xxx_messageInfo_ApiRequest.Size(m)
}
func (m *ApiRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ApiRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ApiRequest proto.InternalMessageInfo

func (m *ApiRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ApiRequest) GetServices() string {
	if m != nil {
		return m.Services
	}
	return ""
}

type Request struct {
	Name                 string     `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Version              string     `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	Services             []*Service `protobuf:"bytes,3,rep,name=services,proto3" json:"services,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_c33392ef2c1961ba, []int{1}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Request) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Request) GetServices() []*Service {
	if m != nil {
		return m.Services
	}
	return nil
}

type Response struct {
	Msg                  string            `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	Chain                []*Response_Chain `protobuf:"bytes,2,rep,name=chain,proto3" json:"chain,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_c33392ef2c1961ba, []int{2}
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

func (m *Response) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *Response) GetChain() []*Response_Chain {
	if m != nil {
		return m.Chain
	}
	return nil
}

type Response_Chain struct {
	Service              *Service          `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	Ctx                  string            `protobuf:"bytes,2,opt,name=ctx,proto3" json:"ctx,omitempty"`
	Chain                []*Response_Chain `protobuf:"bytes,3,rep,name=chain,proto3" json:"chain,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Response_Chain) Reset()         { *m = Response_Chain{} }
func (m *Response_Chain) String() string { return proto.CompactTextString(m) }
func (*Response_Chain) ProtoMessage()    {}
func (*Response_Chain) Descriptor() ([]byte, []int) {
	return fileDescriptor_c33392ef2c1961ba, []int{2, 0}
}

func (m *Response_Chain) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response_Chain.Unmarshal(m, b)
}
func (m *Response_Chain) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response_Chain.Marshal(b, m, deterministic)
}
func (m *Response_Chain) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response_Chain.Merge(m, src)
}
func (m *Response_Chain) XXX_Size() int {
	return xxx_messageInfo_Response_Chain.Size(m)
}
func (m *Response_Chain) XXX_DiscardUnknown() {
	xxx_messageInfo_Response_Chain.DiscardUnknown(m)
}

var xxx_messageInfo_Response_Chain proto.InternalMessageInfo

func (m *Response_Chain) GetService() *Service {
	if m != nil {
		return m.Service
	}
	return nil
}

func (m *Response_Chain) GetCtx() string {
	if m != nil {
		return m.Ctx
	}
	return ""
}

func (m *Response_Chain) GetChain() []*Response_Chain {
	if m != nil {
		return m.Chain
	}
	return nil
}

type Node struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Host                 string   `protobuf:"bytes,2,opt,name=host,proto3" json:"host,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Node) Reset()         { *m = Node{} }
func (m *Node) String() string { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()    {}
func (*Node) Descriptor() ([]byte, []int) {
	return fileDescriptor_c33392ef2c1961ba, []int{3}
}

func (m *Node) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Node.Unmarshal(m, b)
}
func (m *Node) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Node.Marshal(b, m, deterministic)
}
func (m *Node) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Node.Merge(m, src)
}
func (m *Node) XXX_Size() int {
	return xxx_messageInfo_Node.Size(m)
}
func (m *Node) XXX_DiscardUnknown() {
	xxx_messageInfo_Node.DiscardUnknown(m)
}

var xxx_messageInfo_Node proto.InternalMessageInfo

func (m *Node) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Node) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

type Service struct {
	Name                 string     `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Version              string     `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	Node                 *Node      `protobuf:"bytes,3,opt,name=node,proto3" json:"node,omitempty"`
	Services             []*Service `protobuf:"bytes,4,rep,name=services,proto3" json:"services,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Service) Reset()         { *m = Service{} }
func (m *Service) String() string { return proto.CompactTextString(m) }
func (*Service) ProtoMessage()    {}
func (*Service) Descriptor() ([]byte, []int) {
	return fileDescriptor_c33392ef2c1961ba, []int{4}
}

func (m *Service) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Service.Unmarshal(m, b)
}
func (m *Service) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Service.Marshal(b, m, deterministic)
}
func (m *Service) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Service.Merge(m, src)
}
func (m *Service) XXX_Size() int {
	return xxx_messageInfo_Service.Size(m)
}
func (m *Service) XXX_DiscardUnknown() {
	xxx_messageInfo_Service.DiscardUnknown(m)
}

var xxx_messageInfo_Service proto.InternalMessageInfo

func (m *Service) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Service) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Service) GetNode() *Node {
	if m != nil {
		return m.Node
	}
	return nil
}

func (m *Service) GetServices() []*Service {
	if m != nil {
		return m.Services
	}
	return nil
}

func init() {
	proto.RegisterType((*ApiRequest)(nil), "com.hbchen.ApiRequest")
	proto.RegisterType((*Request)(nil), "com.hbchen.Request")
	proto.RegisterType((*Response)(nil), "com.hbchen.Response")
	proto.RegisterType((*Response_Chain)(nil), "com.hbchen.Response.Chain")
	proto.RegisterType((*Node)(nil), "com.hbchen.Node")
	proto.RegisterType((*Service)(nil), "com.hbchen.Service")
}

func init() { proto.RegisterFile("proto/service.proto", fileDescriptor_c33392ef2c1961ba) }

var fileDescriptor_c33392ef2c1961ba = []byte{
	// 400 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xc1, 0xaa, 0x9b, 0x40,
	0x14, 0xad, 0x71, 0x7c, 0xfa, 0x6e, 0xe1, 0xf1, 0x98, 0x94, 0x22, 0x42, 0x41, 0x24, 0x0b, 0x29,
	0x44, 0x4b, 0xb2, 0xea, 0x32, 0x86, 0x6e, 0x0b, 0x9d, 0xee, 0xba, 0x33, 0x3a, 0xc4, 0x01, 0x75,
	0xac, 0x63, 0x24, 0xe9, 0x2a, 0x5f, 0xd0, 0x0f, 0xeb, 0x27, 0x65, 0x55, 0x66, 0xd4, 0x6a, 0x68,
	0x9b, 0x97, 0x95, 0x77, 0x3c, 0xe7, 0x9e, 0x73, 0xee, 0x9d, 0x81, 0x79, 0x55, 0xf3, 0x86, 0x87,
	0x82, 0xd6, 0x2d, 0x4b, 0x68, 0xa0, 0x4e, 0x18, 0x12, 0x5e, 0x04, 0xd9, 0x2e, 0xc9, 0x68, 0xe9,
	0x6c, 0xf6, 0xac, 0xc9, 0x0e, 0xbb, 0x20, 0xe1, 0x45, 0x48, 0xcb, 0x96, 0x9f, 0xaa, 0x9a, 0x1f,
	0x4f, 0xa1, 0x22, 0x26, 0xcb, 0x3d, 0x2d, 0x97, 0x6d, 0x9c, 0xb3, 0x34, 0x6e, 0x68, 0xf8, 0x57,
	0xd1, 0xc9, 0x79, 0x5f, 0x00, 0x36, 0x15, 0x23, 0xf4, 0xfb, 0x81, 0x8a, 0x06, 0xbf, 0x03, 0x54,
	0xc6, 0x05, 0xb5, 0x35, 0x57, 0xf3, 0x1f, 0xa3, 0xc7, 0x4b, 0xf4, 0x50, 0x23, 0x17, 0xf9, 0x2e,
	0x51, 0xbf, 0xf1, 0x02, 0xac, 0x3e, 0x8c, 0xb0, 0x67, 0x8a, 0x62, 0x5d, 0x22, 0xa3, 0xd6, 0xfd,
	0xb3, 0x45, 0xfe, 0x20, 0x9e, 0x00, 0xf3, 0xa6, 0x9e, 0xe6, 0xa3, 0x5e, 0xcf, 0x06, 0xb3, 0xa5,
	0xb5, 0x60, 0xbc, 0xec, 0xe4, 0xc8, 0x70, 0xc4, 0xe1, 0xc4, 0x49, 0x77, 0x75, 0xff, 0xf5, 0x6a,
	0x1e, 0x8c, 0x83, 0x07, 0x5f, 0x3b, 0x6c, 0x62, 0xfa, 0x4b, 0x03, 0x8b, 0x50, 0x51, 0xf1, 0x52,
	0x50, 0xfc, 0x0c, 0x7a, 0x21, 0xf6, 0x9d, 0x2b, 0x91, 0x25, 0xfe, 0x00, 0x46, 0x92, 0xc5, 0x4c,
	0xfa, 0x48, 0x31, 0x67, 0x2a, 0x36, 0xb4, 0x05, 0x5b, 0xc9, 0x20, 0x1d, 0xd1, 0xf9, 0x01, 0x86,
	0x3a, 0xe3, 0x25, 0x98, 0xbd, 0x8b, 0x12, 0xfc, 0x4f, 0x92, 0x81, 0x23, 0xbd, 0x93, 0xe6, 0xd8,
	0xcf, 0x23, 0xcb, 0xd1, 0x5b, 0xbf, 0xd3, 0xdb, 0x7b, 0x0f, 0xe8, 0x33, 0x4f, 0x29, 0x7e, 0x82,
	0x19, 0x4b, 0xfb, 0x31, 0x66, 0x2c, 0xc5, 0x18, 0x50, 0xc6, 0x45, 0xd3, 0x8b, 0xab, 0xda, 0xfb,
	0xa9, 0x81, 0xd9, 0x87, 0x90, 0xf8, 0xb8, 0xee, 0x17, 0x77, 0xbc, 0x00, 0x54, 0xf2, 0x94, 0xda,
	0xba, 0x9a, 0xea, 0x79, 0x1a, 0x4b, 0xba, 0x13, 0x85, 0x5e, 0xdd, 0x04, 0xba, 0xe3, 0x26, 0x56,
	0x67, 0x0d, 0x9e, 0x3e, 0x1d, 0xe3, 0xa2, 0xca, 0xe9, 0x90, 0xeb, 0x23, 0x98, 0x9b, 0x8a, 0x6d,
	0xe3, 0x3c, 0xc7, 0x6f, 0xa7, 0xcd, 0xe3, 0xcb, 0x73, 0xde, 0xfc, 0x6b, 0x2b, 0xde, 0x2b, 0xbc,
	0x06, 0xa4, 0xfa, 0xe6, 0xd7, 0xf8, 0xcd, 0xa6, 0xc8, 0xfc, 0x66, 0xa8, 0xd7, 0xbd, 0x7b, 0x50,
	0x9f, 0xf5, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa0, 0xd0, 0xed, 0xf1, 0x4a, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ExampleServiceClient is the client API for ExampleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ExampleServiceClient interface {
	ApiCall(ctx context.Context, in *ApiRequest, opts ...grpc.CallOption) (*Response, error)
	Call(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type exampleServiceClient struct {
	cc *grpc.ClientConn
}

func NewExampleServiceClient(cc *grpc.ClientConn) ExampleServiceClient {
	return &exampleServiceClient{cc}
}

func (c *exampleServiceClient) ApiCall(ctx context.Context, in *ApiRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/com.hbchen.ExampleService/ApiCall", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exampleServiceClient) Call(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/com.hbchen.ExampleService/Call", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExampleServiceServer is the server API for ExampleService service.
type ExampleServiceServer interface {
	ApiCall(context.Context, *ApiRequest) (*Response, error)
	Call(context.Context, *Request) (*Response, error)
}

// UnimplementedExampleServiceServer can be embedded to have forward compatible implementations.
type UnimplementedExampleServiceServer struct {
}

func (*UnimplementedExampleServiceServer) ApiCall(ctx context.Context, req *ApiRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApiCall not implemented")
}
func (*UnimplementedExampleServiceServer) Call(ctx context.Context, req *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Call not implemented")
}

func RegisterExampleServiceServer(s *grpc.Server, srv ExampleServiceServer) {
	s.RegisterService(&_ExampleService_serviceDesc, srv)
}

func _ExampleService_ApiCall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApiRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExampleServiceServer).ApiCall(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.hbchen.ExampleService/ApiCall",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExampleServiceServer).ApiCall(ctx, req.(*ApiRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExampleService_Call_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExampleServiceServer).Call(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.hbchen.ExampleService/Call",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExampleServiceServer).Call(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _ExampleService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "com.hbchen.ExampleService",
	HandlerType: (*ExampleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ApiCall",
			Handler:    _ExampleService_ApiCall_Handler,
		},
		{
			MethodName: "Call",
			Handler:    _ExampleService_Call_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/service.proto",
}
