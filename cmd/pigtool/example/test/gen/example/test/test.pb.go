// Code generated by protoc-gen-go. DO NOT EDIT.
// source: test.proto

package test

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	codec "github.com/mike-zeng/pigkit/rpc/codec"
	server "github.com/mike-zeng/pigkit/rpc/server"
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

type TestRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TestRequest) Reset()         { *m = TestRequest{} }
func (m *TestRequest) String() string { return proto.CompactTextString(m) }
func (*TestRequest) ProtoMessage()    {}
func (*TestRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c161fcfdc0c3ff1e, []int{0}
}

func (m *TestRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestRequest.Unmarshal(m, b)
}
func (m *TestRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestRequest.Marshal(b, m, deterministic)
}
func (m *TestRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestRequest.Merge(m, src)
}
func (m *TestRequest) XXX_Size() int {
	return xxx_messageInfo_TestRequest.Size(m)
}
func (m *TestRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TestRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TestRequest proto.InternalMessageInfo

func (m *TestRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type TestResponse struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TestResponse) Reset()         { *m = TestResponse{} }
func (m *TestResponse) String() string { return proto.CompactTextString(m) }
func (*TestResponse) ProtoMessage()    {}
func (*TestResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c161fcfdc0c3ff1e, []int{1}
}

func (m *TestResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestResponse.Unmarshal(m, b)
}
func (m *TestResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestResponse.Marshal(b, m, deterministic)
}
func (m *TestResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestResponse.Merge(m, src)
}
func (m *TestResponse) XXX_Size() int {
	return xxx_messageInfo_TestResponse.Size(m)
}
func (m *TestResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TestResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TestResponse proto.InternalMessageInfo

func (m *TestResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*TestRequest)(nil), "TestRequest")
	proto.RegisterType((*TestResponse)(nil), "TestResponse")
}

func init() { proto.RegisterFile("test.proto", fileDescriptor_c161fcfdc0c3ff1e) }

var fileDescriptor_c161fcfdc0c3ff1e = []byte{
	// 131 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x49, 0x2d, 0x2e,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x52, 0xe4, 0xe2, 0x0e, 0x49, 0x2d, 0x2e, 0x09, 0x4a,
	0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x12, 0xe2, 0x62, 0xc9, 0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x54,
	0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb3, 0x95, 0x14, 0xb8, 0x78, 0x20, 0x4a, 0x8a, 0x0b, 0xf2, 0xf3,
	0x8a, 0x53, 0x85, 0x04, 0xb8, 0x98, 0x73, 0x8b, 0xd3, 0xa1, 0x4a, 0x40, 0x4c, 0x23, 0x5d, 0x2e,
	0x16, 0x90, 0x0a, 0x21, 0x55, 0x28, 0xcd, 0xa3, 0x87, 0x64, 0xa6, 0x14, 0xaf, 0x1e, 0xb2, 0x76,
	0x25, 0x06, 0x27, 0xbe, 0x28, 0x9e, 0xd4, 0x8a, 0xc4, 0xdc, 0x82, 0x9c, 0x54, 0x7d, 0x90, 0x4b,
	0x92, 0xd8, 0xc0, 0x4e, 0x31, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x50, 0x03, 0xac, 0x19, 0x98,
	0x00, 0x00, 0x00,
}

// This following code was generated by protoc-gen-pigrpc, DO NOT EDIT!!!

//================== server skeleton ===================
type TestService interface {
	Test(ctx context.Context, req *TestRequest) (*TestResponse, error)
}

var TestServiceDesc = &server.ServiceDesc{
	ServiceName: ".Test",
	HandlerType: (*TestService)(nil),
	Methods: map[string]server.Handler{

		"Test": TestService_Test_Handler,
	},
}

func TestService_Test_Handler(ctx context.Context, s interface{}, pigReq *codec.PigReq, fillReq func(interface{}, *codec.PigReq) error) (interface{}, error) {
	req := new(TestRequest)
	if err := fillReq(req, pigReq); err != nil {
		return nil, err
	}
	return s.(TestService).Test(ctx, req)
}
