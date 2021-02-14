// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/start.proto

package go_micro_srv_send

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
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
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Start service

type StartService interface {
	SendMessage(ctx context.Context, in *CallRequest, opts ...client.CallOption) (*CallResponse, error)
}

type startService struct {
	c    client.Client
	name string
}

func NewStartService(name string, c client.Client) StartService {
	return &startService{
		c:    c,
		name: name,
	}
}

func (c *startService) SendMessage(ctx context.Context, in *CallRequest, opts ...client.CallOption) (*CallResponse, error) {
	req := c.c.NewRequest(c.name, "Start.SendMessage", in)
	out := new(CallResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Start service

type StartHandler interface {
	SendMessage(context.Context, *CallRequest, *CallResponse) error
}

func RegisterStartHandler(s server.Server, hdlr StartHandler, opts ...server.HandlerOption) error {
	type start interface {
		SendMessage(ctx context.Context, in *CallRequest, out *CallResponse) error
	}
	type Start struct {
		start
	}
	h := &startHandler{hdlr}
	return s.Handle(s.NewHandler(&Start{h}, opts...))
}

type startHandler struct {
	StartHandler
}

func (h *startHandler) SendMessage(ctx context.Context, in *CallRequest, out *CallResponse) error {
	return h.StartHandler.SendMessage(ctx, in, out)
}
