// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/first.proto

package first

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v3/api"
	client "github.com/micro/go-micro/v3/client"
	server "github.com/micro/go-micro/v3/server"
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

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for First service

func NewFirstEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for First service

type FirstService interface {
	Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	Stream(ctx context.Context, in *StreamingRequest, opts ...client.CallOption) (First_StreamService, error)
	PingPong(ctx context.Context, opts ...client.CallOption) (First_PingPongService, error)
}

type firstService struct {
	c    client.Client
	name string
}

func NewFirstService(name string, c client.Client) FirstService {
	return &firstService{
		c:    c,
		name: name,
	}
}

func (c *firstService) Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "First.Call", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *firstService) Stream(ctx context.Context, in *StreamingRequest, opts ...client.CallOption) (First_StreamService, error) {
	req := c.c.NewRequest(c.name, "First.Stream", &StreamingRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &firstServiceStream{stream}, nil
}

type First_StreamService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*StreamingResponse, error)
}

type firstServiceStream struct {
	stream client.Stream
}

func (x *firstServiceStream) Close() error {
	return x.stream.Close()
}

func (x *firstServiceStream) Context() context.Context {
	return x.stream.Context()
}

func (x *firstServiceStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *firstServiceStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *firstServiceStream) Recv() (*StreamingResponse, error) {
	m := new(StreamingResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *firstService) PingPong(ctx context.Context, opts ...client.CallOption) (First_PingPongService, error) {
	req := c.c.NewRequest(c.name, "First.PingPong", &Ping{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &firstServicePingPong{stream}, nil
}

type First_PingPongService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Ping) error
	Recv() (*Pong, error)
}

type firstServicePingPong struct {
	stream client.Stream
}

func (x *firstServicePingPong) Close() error {
	return x.stream.Close()
}

func (x *firstServicePingPong) Context() context.Context {
	return x.stream.Context()
}

func (x *firstServicePingPong) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *firstServicePingPong) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *firstServicePingPong) Send(m *Ping) error {
	return x.stream.Send(m)
}

func (x *firstServicePingPong) Recv() (*Pong, error) {
	m := new(Pong)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for First service

type FirstHandler interface {
	Call(context.Context, *Request, *Response) error
	Stream(context.Context, *StreamingRequest, First_StreamStream) error
	PingPong(context.Context, First_PingPongStream) error
}

func RegisterFirstHandler(s server.Server, hdlr FirstHandler, opts ...server.HandlerOption) error {
	type first interface {
		Call(ctx context.Context, in *Request, out *Response) error
		Stream(ctx context.Context, stream server.Stream) error
		PingPong(ctx context.Context, stream server.Stream) error
	}
	type First struct {
		first
	}
	h := &firstHandler{hdlr}
	return s.Handle(s.NewHandler(&First{h}, opts...))
}

type firstHandler struct {
	FirstHandler
}

func (h *firstHandler) Call(ctx context.Context, in *Request, out *Response) error {
	return h.FirstHandler.Call(ctx, in, out)
}

func (h *firstHandler) Stream(ctx context.Context, stream server.Stream) error {
	m := new(StreamingRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.FirstHandler.Stream(ctx, m, &firstStreamStream{stream})
}

type First_StreamStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*StreamingResponse) error
}

type firstStreamStream struct {
	stream server.Stream
}

func (x *firstStreamStream) Close() error {
	return x.stream.Close()
}

func (x *firstStreamStream) Context() context.Context {
	return x.stream.Context()
}

func (x *firstStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *firstStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *firstStreamStream) Send(m *StreamingResponse) error {
	return x.stream.Send(m)
}

func (h *firstHandler) PingPong(ctx context.Context, stream server.Stream) error {
	return h.FirstHandler.PingPong(ctx, &firstPingPongStream{stream})
}

type First_PingPongStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Pong) error
	Recv() (*Ping, error)
}

type firstPingPongStream struct {
	stream server.Stream
}

func (x *firstPingPongStream) Close() error {
	return x.stream.Close()
}

func (x *firstPingPongStream) Context() context.Context {
	return x.stream.Context()
}

func (x *firstPingPongStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *firstPingPongStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *firstPingPongStream) Send(m *Pong) error {
	return x.stream.Send(m)
}

func (x *firstPingPongStream) Recv() (*Ping, error) {
	m := new(Ping)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}
