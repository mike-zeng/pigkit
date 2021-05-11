// Code generated by protoc-gen-go. DO NOT EDIT.
// source: examples/greeter/greeter.proto

package greeter

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	gorpc "github.com/lubanproj/gorpc"
	client "github.com/lubanproj/gorpc/client"
	interceptor "github.com/lubanproj/gorpc/interceptor"
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

type Request struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_3cc58d5671cc8265, []int{0}
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

type Response struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_3cc58d5671cc8265, []int{1}
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

func init() {
	proto.RegisterType((*Request)(nil), "Request")
	proto.RegisterType((*Response)(nil), "Response")
}

func init() { proto.RegisterFile("examples/greeter/greeter.proto", fileDescriptor_3cc58d5671cc8265) }

var fileDescriptor_3cc58d5671cc8265 = []byte{
	// 133 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4b, 0xad, 0x48, 0xcc,
	0x2d, 0xc8, 0x49, 0x2d, 0xd6, 0x4f, 0x2f, 0x4a, 0x4d, 0x2d, 0x49, 0x2d, 0x82, 0xd1, 0x7a, 0x05,
	0x45, 0xf9, 0x25, 0xf9, 0x4a, 0xb2, 0x5c, 0xec, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42,
	0x42, 0x5c, 0x2c, 0x79, 0x89, 0xb9, 0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x60, 0xb6,
	0x92, 0x0c, 0x17, 0x47, 0x50, 0x6a, 0x71, 0x41, 0x7e, 0x5e, 0x71, 0xaa, 0x90, 0x00, 0x17, 0x73,
	0x6e, 0x71, 0x3a, 0x54, 0x1a, 0xc4, 0x34, 0xd2, 0xe4, 0x62, 0x77, 0x87, 0x98, 0x26, 0x24, 0xc7,
	0xc5, 0xea, 0x91, 0x9a, 0x93, 0x93, 0x2f, 0xc4, 0xa1, 0x07, 0x35, 0x4f, 0x8a, 0x53, 0x0f, 0xa6,
	0x55, 0x89, 0x21, 0x89, 0x0d, 0x6c, 0x9d, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xee, 0x1e, 0x1c,
	0x2a, 0x90, 0x00, 0x00, 0x00,
}

// This following code was generated by protoc-gen-gorpc, DO NOT EDIT!!!

//================== server skeleton ===================
type GreeterService interface {
	Hello(ctx context.Context, req *Request) (*Response, error)
}

var _Greeter_serviceDesc = &gorpc.ServiceDesc{
	ServiceName: ".Greeter",
	HandlerType: (*GreeterService)(nil),
	Methods: []*gorpc.MethodDesc{

		{
			MethodName: "Hello",
			Handler:    GreeterService_Hello_Handler,
		},
	},
}

func GreeterService_Hello_Handler(ctx context.Context, svr interface{}, dec func(interface{}) error, ceps []interceptor.ServerInterceptor) (interface{}, error) {

	req := new(Request)
	if err := dec(req); err != nil {
		return nil, err
	}

	if len(ceps) == 0 {
		return svr.(GreeterService).Hello(ctx, req)
	}

	handler := func(ctx context.Context, reqbody interface{}) (interface{}, error) {
		return svr.(GreeterService).Hello(ctx, reqbody.(*Request))
	}

	return interceptor.ServerIntercept(ctx, req, ceps, handler)
}

func RegisterService(s *gorpc.Server, svr interface{}) {
	s.Register(_Greeter_serviceDesc, svr)
}

//================== client stub===================
//GreeterClientProxy is a client proxy for service Greeter.
type GreeterClientProxy interface {
	Hello(ctx context.Context, req *Request, opts ...client.Option) (*Response, error)
}

// Hello is server rpc method as defined
func (c *GreeterImpl) Hello(ctx context.Context, req *Request, opts ...client.Option) (*Response, error) {

	callopts := make([]client.Option, 0, len(c.opts)+len(opts))
	callopts = append(callopts, c.opts...)
	callopts = append(callopts, opts...)

	rsp := &Response{}
	err := c.client.Invoke(ctx, req, rsp, "/.Greeter/Hello", callopts...)
	if err != nil {
		return nil, err
	}

	return rsp, nil
}
