// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/second.proto

package second

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

// Api Endpoints for Second service

func NewSecondEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Second service

type SecondService interface {
	Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	Stream(ctx context.Context, in *StreamingRequest, opts ...client.CallOption) (Second_StreamService, error)
	PingPong(ctx context.Context, opts ...client.CallOption) (Second_PingPongService, error)
}

type secondService struct {
	c    client.Client
	name string
}

func NewSecondService(name string, c client.Client) SecondService {
	return &secondService{
		c:    c,
		name: name,
	}
}

func (c *secondService) Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Second.Call", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secondService) Stream(ctx context.Context, in *StreamingRequest, opts ...client.CallOption) (Second_StreamService, error) {
	req := c.c.NewRequest(c.name, "Second.Stream", &StreamingRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &secondServiceStream{stream}, nil
}

type Second_StreamService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*StreamingResponse, error)
}

type secondServiceStream struct {
	stream client.Stream
}

func (x *secondServiceStream) Close() error {
	return x.stream.Close()
}

func (x *secondServiceStream) Context() context.Context {
	return x.stream.Context()
}

func (x *secondServiceStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *secondServiceStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *secondServiceStream) Recv() (*StreamingResponse, error) {
	m := new(StreamingResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *secondService) PingPong(ctx context.Context, opts ...client.CallOption) (Second_PingPongService, error) {
	req := c.c.NewRequest(c.name, "Second.PingPong", &Ping{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &secondServicePingPong{stream}, nil
}

type Second_PingPongService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Ping) error
	Recv() (*Pong, error)
}

type secondServicePingPong struct {
	stream client.Stream
}

func (x *secondServicePingPong) Close() error {
	return x.stream.Close()
}

func (x *secondServicePingPong) Context() context.Context {
	return x.stream.Context()
}

func (x *secondServicePingPong) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *secondServicePingPong) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *secondServicePingPong) Send(m *Ping) error {
	return x.stream.Send(m)
}

func (x *secondServicePingPong) Recv() (*Pong, error) {
	m := new(Pong)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Second service

type SecondHandler interface {
	Call(context.Context, *Request, *Response) error
	Stream(context.Context, *StreamingRequest, Second_StreamStream) error
	PingPong(context.Context, Second_PingPongStream) error
}

func RegisterSecondHandler(s server.Server, hdlr SecondHandler, opts ...server.HandlerOption) error {
	type second interface {
		Call(ctx context.Context, in *Request, out *Response) error
		Stream(ctx context.Context, stream server.Stream) error
		PingPong(ctx context.Context, stream server.Stream) error
	}
	type Second struct {
		second
	}
	h := &secondHandler{hdlr}
	return s.Handle(s.NewHandler(&Second{h}, opts...))
}

type secondHandler struct {
	SecondHandler
}

func (h *secondHandler) Call(ctx context.Context, in *Request, out *Response) error {
	return h.SecondHandler.Call(ctx, in, out)
}

func (h *secondHandler) Stream(ctx context.Context, stream server.Stream) error {
	m := new(StreamingRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.SecondHandler.Stream(ctx, m, &secondStreamStream{stream})
}

type Second_StreamStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*StreamingResponse) error
}

type secondStreamStream struct {
	stream server.Stream
}

func (x *secondStreamStream) Close() error {
	return x.stream.Close()
}

func (x *secondStreamStream) Context() context.Context {
	return x.stream.Context()
}

func (x *secondStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *secondStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *secondStreamStream) Send(m *StreamingResponse) error {
	return x.stream.Send(m)
}

func (h *secondHandler) PingPong(ctx context.Context, stream server.Stream) error {
	return h.SecondHandler.PingPong(ctx, &secondPingPongStream{stream})
}

type Second_PingPongStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Pong) error
	Recv() (*Ping, error)
}

type secondPingPongStream struct {
	stream server.Stream
}

func (x *secondPingPongStream) Close() error {
	return x.stream.Close()
}

func (x *secondPingPongStream) Context() context.Context {
	return x.stream.Context()
}

func (x *secondPingPongStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *secondPingPongStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *secondPingPongStream) Send(m *Pong) error {
	return x.stream.Send(m)
}

func (x *secondPingPongStream) Recv() (*Ping, error) {
	m := new(Ping)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}
