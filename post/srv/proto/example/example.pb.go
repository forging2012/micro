// Code generated by protoc-gen-go. DO NOT EDIT.
// source: post/srv/proto/example/example.proto

/*
Package go_micro_srv_post is a generated protocol buffer package.

It is generated from these files:
	post/srv/proto/example/example.proto

It has these top-level messages:
	Message
	Request
	Response
	StreamingRequest
	StreamingResponse
	Ping
	Pong
*/
package go_micro_srv_post

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
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

type Message struct {
	Say string `protobuf:"bytes,1,opt,name=say" json:"say,omitempty"`
}

func (m *Message) Reset()                    { *m = Message{} }
func (m *Message) String() string            { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()               {}
func (*Message) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Message) GetSay() string {
	if m != nil {
		return m.Say
	}
	return ""
}

type Request struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Request) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Response struct {
	Msg string `protobuf:"bytes,1,opt,name=msg" json:"msg,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Response) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type StreamingRequest struct {
	Count int64 `protobuf:"varint,1,opt,name=count" json:"count,omitempty"`
}

func (m *StreamingRequest) Reset()                    { *m = StreamingRequest{} }
func (m *StreamingRequest) String() string            { return proto.CompactTextString(m) }
func (*StreamingRequest) ProtoMessage()               {}
func (*StreamingRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *StreamingRequest) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type StreamingResponse struct {
	Count int64 `protobuf:"varint,1,opt,name=count" json:"count,omitempty"`
}

func (m *StreamingResponse) Reset()                    { *m = StreamingResponse{} }
func (m *StreamingResponse) String() string            { return proto.CompactTextString(m) }
func (*StreamingResponse) ProtoMessage()               {}
func (*StreamingResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *StreamingResponse) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type Ping struct {
	Stroke int64 `protobuf:"varint,1,opt,name=stroke" json:"stroke,omitempty"`
}

func (m *Ping) Reset()                    { *m = Ping{} }
func (m *Ping) String() string            { return proto.CompactTextString(m) }
func (*Ping) ProtoMessage()               {}
func (*Ping) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Ping) GetStroke() int64 {
	if m != nil {
		return m.Stroke
	}
	return 0
}

type Pong struct {
	Stroke int64 `protobuf:"varint,1,opt,name=stroke" json:"stroke,omitempty"`
}

func (m *Pong) Reset()                    { *m = Pong{} }
func (m *Pong) String() string            { return proto.CompactTextString(m) }
func (*Pong) ProtoMessage()               {}
func (*Pong) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Pong) GetStroke() int64 {
	if m != nil {
		return m.Stroke
	}
	return 0
}

func init() {
	proto.RegisterType((*Message)(nil), "go.micro.srv.post.Message")
	proto.RegisterType((*Request)(nil), "go.micro.srv.post.Request")
	proto.RegisterType((*Response)(nil), "go.micro.srv.post.Response")
	proto.RegisterType((*StreamingRequest)(nil), "go.micro.srv.post.StreamingRequest")
	proto.RegisterType((*StreamingResponse)(nil), "go.micro.srv.post.StreamingResponse")
	proto.RegisterType((*Ping)(nil), "go.micro.srv.post.Ping")
	proto.RegisterType((*Pong)(nil), "go.micro.srv.post.Pong")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Example service

type ExampleClient interface {
	Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	Stream(ctx context.Context, in *StreamingRequest, opts ...client.CallOption) (Example_StreamClient, error)
	PingPong(ctx context.Context, opts ...client.CallOption) (Example_PingPongClient, error)
}

type exampleClient struct {
	c           client.Client
	serviceName string
}

func NewExampleClient(serviceName string, c client.Client) ExampleClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "go.micro.srv.post"
	}
	return &exampleClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *exampleClient) Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "Example.Call", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exampleClient) Stream(ctx context.Context, in *StreamingRequest, opts ...client.CallOption) (Example_StreamClient, error) {
	req := c.c.NewRequest(c.serviceName, "Example.Stream", &StreamingRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &exampleStreamClient{stream}, nil
}

type Example_StreamClient interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*StreamingResponse, error)
}

type exampleStreamClient struct {
	stream client.Streamer
}

func (x *exampleStreamClient) Close() error {
	return x.stream.Close()
}

func (x *exampleStreamClient) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *exampleStreamClient) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *exampleStreamClient) Recv() (*StreamingResponse, error) {
	m := new(StreamingResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *exampleClient) PingPong(ctx context.Context, opts ...client.CallOption) (Example_PingPongClient, error) {
	req := c.c.NewRequest(c.serviceName, "Example.PingPong", &Ping{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &examplePingPongClient{stream}, nil
}

type Example_PingPongClient interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Ping) error
	Recv() (*Pong, error)
}

type examplePingPongClient struct {
	stream client.Streamer
}

func (x *examplePingPongClient) Close() error {
	return x.stream.Close()
}

func (x *examplePingPongClient) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *examplePingPongClient) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *examplePingPongClient) Send(m *Ping) error {
	return x.stream.Send(m)
}

func (x *examplePingPongClient) Recv() (*Pong, error) {
	m := new(Pong)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Example service

type ExampleHandler interface {
	Call(context.Context, *Request, *Response) error
	Stream(context.Context, *StreamingRequest, Example_StreamStream) error
	PingPong(context.Context, Example_PingPongStream) error
}

func RegisterExampleHandler(s server.Server, hdlr ExampleHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&Example{hdlr}, opts...))
}

type Example struct {
	ExampleHandler
}

func (h *Example) Call(ctx context.Context, in *Request, out *Response) error {
	return h.ExampleHandler.Call(ctx, in, out)
}

func (h *Example) Stream(ctx context.Context, stream server.Streamer) error {
	m := new(StreamingRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.ExampleHandler.Stream(ctx, m, &exampleStreamStream{stream})
}

type Example_StreamStream interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*StreamingResponse) error
}

type exampleStreamStream struct {
	stream server.Streamer
}

func (x *exampleStreamStream) Close() error {
	return x.stream.Close()
}

func (x *exampleStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *exampleStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *exampleStreamStream) Send(m *StreamingResponse) error {
	return x.stream.Send(m)
}

func (h *Example) PingPong(ctx context.Context, stream server.Streamer) error {
	return h.ExampleHandler.PingPong(ctx, &examplePingPongStream{stream})
}

type Example_PingPongStream interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Pong) error
	Recv() (*Ping, error)
}

type examplePingPongStream struct {
	stream server.Streamer
}

func (x *examplePingPongStream) Close() error {
	return x.stream.Close()
}

func (x *examplePingPongStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *examplePingPongStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *examplePingPongStream) Send(m *Pong) error {
	return x.stream.Send(m)
}

func (x *examplePingPongStream) Recv() (*Ping, error) {
	m := new(Ping)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}

func init() { proto.RegisterFile("post/srv/proto/example/example.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 279 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xc1, 0x4a, 0xc4, 0x30,
	0x10, 0x86, 0x37, 0x6c, 0x6d, 0xd7, 0x39, 0xed, 0x06, 0x51, 0xe9, 0xaa, 0x48, 0xdc, 0x43, 0xbd,
	0xa4, 0x8b, 0x3e, 0x81, 0x8a, 0x47, 0x41, 0xea, 0xc1, 0x73, 0x5c, 0x42, 0x28, 0x36, 0x49, 0xcd,
	0x64, 0x17, 0x7d, 0x64, 0xdf, 0x42, 0x9a, 0xb6, 0x20, 0x1a, 0x3d, 0x75, 0xa6, 0xdf, 0xff, 0x0f,
	0xf3, 0x4f, 0x60, 0xd5, 0x5a, 0xf4, 0x25, 0xba, 0x5d, 0xd9, 0x3a, 0xeb, 0x6d, 0x29, 0xdf, 0x85,
	0x6e, 0x1b, 0x39, 0x7e, 0x79, 0xf8, 0x4b, 0x17, 0xca, 0x72, 0x5d, 0x6f, 0x9c, 0xe5, 0xe8, 0x76,
	0xbc, 0xb3, 0xb0, 0x25, 0x64, 0x0f, 0x12, 0x51, 0x28, 0x49, 0xe7, 0x30, 0x45, 0xf1, 0x71, 0x4c,
	0xce, 0x49, 0xb1, 0x5f, 0x75, 0x25, 0x3b, 0x85, 0xac, 0x92, 0x6f, 0x5b, 0x89, 0x9e, 0x52, 0x48,
	0x8c, 0xd0, 0x72, 0xa0, 0xa1, 0x66, 0x27, 0x30, 0xab, 0x24, 0xb6, 0xd6, 0x60, 0x30, 0x6b, 0x54,
	0xa3, 0x59, 0xa3, 0x62, 0x05, 0xcc, 0x9f, 0xbc, 0x93, 0x42, 0xd7, 0x46, 0x8d, 0x53, 0x0e, 0x60,
	0x6f, 0x63, 0xb7, 0xc6, 0x07, 0xdd, 0xb4, 0xea, 0x1b, 0x76, 0x09, 0x8b, 0x6f, 0xca, 0x61, 0x60,
	0x5c, 0x7a, 0x06, 0xc9, 0x63, 0x6d, 0x14, 0x3d, 0x84, 0x14, 0xbd, 0xb3, 0xaf, 0x72, 0xc0, 0x43,
	0x17, 0xb8, 0xfd, 0x9b, 0x5f, 0x7d, 0x12, 0xc8, 0xee, 0xfb, 0x9b, 0xd0, 0x1b, 0x48, 0xee, 0x44,
	0xd3, 0xd0, 0x9c, 0xff, 0x3a, 0x0b, 0x1f, 0x16, 0xce, 0x97, 0x51, 0xd6, 0xaf, 0xc8, 0x26, 0xf4,
	0x19, 0xd2, 0x7e, 0x73, 0x7a, 0x11, 0x11, 0xfe, 0x8c, 0x9f, 0xaf, 0xfe, 0x17, 0x8d, 0x63, 0xd7,
	0x84, 0xde, 0xc2, 0xac, 0xcb, 0x19, 0xb2, 0x1c, 0x45, 0x5c, 0x1d, 0xcc, 0xa3, 0xc0, 0x1a, 0xc5,
	0x26, 0x05, 0x59, 0x93, 0x97, 0x34, 0x3c, 0xfa, 0xf5, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0xdd,
	0xc6, 0xf4, 0xd3, 0x1c, 0x02, 0x00, 0x00,
}
