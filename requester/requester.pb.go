// Code generated by protoc-gen-go. DO NOT EDIT.
// source: requester.proto

/*
Package requester is a generated protocol buffer package.

It is generated from these files:
	requester.proto

It has these top-level messages:
	Request
	Response
	KeyValue
	Value
*/
package requester

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Request struct {
	KeyId string `protobuf:"bytes,1,opt,name=key_id,json=keyId" json:"key_id,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Request) GetKeyId() string {
	if m != nil {
		return m.KeyId
	}
	return ""
}

type Response struct {
	Kv  *KeyValue `protobuf:"bytes,1,opt,name=kv" json:"kv,omitempty"`
	Val *Value    `protobuf:"bytes,2,opt,name=val" json:"val,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Response) GetKv() *KeyValue {
	if m != nil {
		return m.Kv
	}
	return nil
}

func (m *Response) GetVal() *Value {
	if m != nil {
		return m.Val
	}
	return nil
}

type KeyValue struct {
	Keyvalue map[string]string `protobuf:"bytes,2,rep,name=keyvalue" json:"keyvalue,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *KeyValue) Reset()                    { *m = KeyValue{} }
func (m *KeyValue) String() string            { return proto.CompactTextString(m) }
func (*KeyValue) ProtoMessage()               {}
func (*KeyValue) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *KeyValue) GetKeyvalue() map[string]string {
	if m != nil {
		return m.Keyvalue
	}
	return nil
}

type Value struct {
	Value string `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
}

func (m *Value) Reset()                    { *m = Value{} }
func (m *Value) String() string            { return proto.CompactTextString(m) }
func (*Value) ProtoMessage()               {}
func (*Value) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Value) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "requester.Request")
	proto.RegisterType((*Response)(nil), "requester.Response")
	proto.RegisterType((*KeyValue)(nil), "requester.KeyValue")
	proto.RegisterType((*Value)(nil), "requester.Value")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Requester service

type RequesterClient interface {
	Process(ctx context.Context, in *Request, opts ...grpc.CallOption) (Requester_ProcessClient, error)
}

type requesterClient struct {
	cc *grpc.ClientConn
}

func NewRequesterClient(cc *grpc.ClientConn) RequesterClient {
	return &requesterClient{cc}
}

func (c *requesterClient) Process(ctx context.Context, in *Request, opts ...grpc.CallOption) (Requester_ProcessClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Requester_serviceDesc.Streams[0], c.cc, "/requester.Requester/Process", opts...)
	if err != nil {
		return nil, err
	}
	x := &requesterProcessClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Requester_ProcessClient interface {
	Recv() (*Response, error)
	grpc.ClientStream
}

type requesterProcessClient struct {
	grpc.ClientStream
}

func (x *requesterProcessClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Requester service

type RequesterServer interface {
	Process(*Request, Requester_ProcessServer) error
}

func RegisterRequesterServer(s *grpc.Server, srv RequesterServer) {
	s.RegisterService(&_Requester_serviceDesc, srv)
}

func _Requester_Process_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Request)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RequesterServer).Process(m, &requesterProcessServer{stream})
}

type Requester_ProcessServer interface {
	Send(*Response) error
	grpc.ServerStream
}

type requesterProcessServer struct {
	grpc.ServerStream
}

func (x *requesterProcessServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

var _Requester_serviceDesc = grpc.ServiceDesc{
	ServiceName: "requester.Requester",
	HandlerType: (*RequesterServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Process",
			Handler:       _Requester_Process_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "requester.proto",
}

func init() { proto.RegisterFile("requester.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 303 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0x4f, 0x4b, 0x33, 0x31,
	0x10, 0xc6, 0x49, 0x4a, 0xff, 0x4d, 0x79, 0x79, 0x4b, 0x6a, 0xb1, 0x14, 0x0b, 0x35, 0x5e, 0x7a,
	0x6a, 0x4a, 0xbd, 0x88, 0xe2, 0xd1, 0x83, 0x78, 0x91, 0x14, 0xbc, 0xea, 0xda, 0x0e, 0xa5, 0xa4,
	0x6c, 0xd6, 0x4d, 0xba, 0x10, 0xc4, 0x8b, 0x07, 0xbf, 0x80, 0x1f, 0xcd, 0xaf, 0xe0, 0x07, 0x91,
	0x4d, 0xba, 0xeb, 0x16, 0xbc, 0x3d, 0x3b, 0xcf, 0x6f, 0x66, 0x9f, 0x99, 0xc0, 0xff, 0x14, 0x5f,
	0x76, 0x68, 0x2c, 0xa6, 0xd3, 0x24, 0xd5, 0x56, 0xb3, 0x76, 0x59, 0x18, 0x9e, 0xac, 0xb5, 0x5e,
	0x6f, 0x51, 0x44, 0xc9, 0x46, 0x44, 0x71, 0xac, 0x6d, 0x64, 0x37, 0x3a, 0x36, 0x01, 0xe4, 0x63,
	0x68, 0xca, 0x80, 0xb2, 0x3e, 0x34, 0x14, 0xba, 0xc7, 0xcd, 0x6a, 0x40, 0xc6, 0x64, 0xd2, 0x96,
	0x75, 0x85, 0xee, 0x76, 0xc5, 0x17, 0xd0, 0x92, 0x68, 0x12, 0x1d, 0x1b, 0x64, 0x67, 0x40, 0x55,
	0xe6, 0xed, 0xce, 0xbc, 0x37, 0xfd, 0xfd, 0xe9, 0x1d, 0xba, 0x87, 0x68, 0xbb, 0x43, 0x49, 0x55,
	0xc6, 0x38, 0xd4, 0xb2, 0x68, 0x3b, 0xa0, 0x9e, 0xea, 0x56, 0xa8, 0x80, 0xe4, 0x26, 0xff, 0x20,
	0xd0, 0x2a, 0x9a, 0xd8, 0x35, 0xb4, 0x14, 0xba, 0x2c, 0xd7, 0x03, 0x3a, 0xae, 0x4d, 0x3a, 0xf3,
	0xd3, 0x3f, 0x66, 0xe7, 0xc2, 0x33, 0x37, 0xb1, 0x4d, 0x9d, 0x2c, 0x5b, 0x86, 0x57, 0xf0, 0xef,
	0xc0, 0x62, 0x5d, 0xa8, 0x29, 0x74, 0xfb, 0x2d, 0x72, 0xc9, 0x8e, 0xa0, 0x5e, 0x8c, 0xf7, 0x9b,
	0xf9, 0x8f, 0x4b, 0x7a, 0x41, 0xf8, 0x08, 0xea, 0x21, 0x44, 0x89, 0x90, 0x0a, 0x32, 0x7f, 0x82,
	0xb6, 0x2c, 0x92, 0xb0, 0x05, 0x34, 0xef, 0x53, 0xbd, 0x44, 0x63, 0x18, 0xab, 0x04, 0xdc, 0x03,
	0xc3, 0xde, 0x41, 0x2d, 0x5c, 0x8c, 0x8f, 0xde, 0xbf, 0xbe, 0x3f, 0xe9, 0x31, 0xeb, 0x8b, 0x6c,
	0x26, 0xd0, 0x2e, 0x57, 0x42, 0xa1, 0x13, 0xaf, 0xe1, 0xd0, 0x6f, 0x33, 0xf2, 0xdc, 0xf0, 0xef,
	0x70, 0xfe, 0x13, 0x00, 0x00, 0xff, 0xff, 0x7d, 0xed, 0x55, 0xb6, 0xc3, 0x01, 0x00, 0x00,
}