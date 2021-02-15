// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/UserInfo.proto

package go_micro_srv_getuserinfo

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

// Client API for UserInfo service

type UserInfoService interface {
	GetInfo(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*GetResponse, error)
}

type userInfoService struct {
	c    client.Client
	name string
}

func NewUserInfoService(name string, c client.Client) UserInfoService {
	return &userInfoService{
		c:    c,
		name: name,
	}
}

func (c *userInfoService) GetInfo(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*GetResponse, error) {
	req := c.c.NewRequest(c.name, "UserInfo.GetInfo", in)
	out := new(GetResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserInfo service

type UserInfoHandler interface {
	GetInfo(context.Context, *GetRequest, *GetResponse) error
}

func RegisterUserInfoHandler(s server.Server, hdlr UserInfoHandler, opts ...server.HandlerOption) error {
	type userInfo interface {
		GetInfo(ctx context.Context, in *GetRequest, out *GetResponse) error
	}
	type UserInfo struct {
		userInfo
	}
	h := &userInfoHandler{hdlr}
	return s.Handle(s.NewHandler(&UserInfo{h}, opts...))
}

type userInfoHandler struct {
	UserInfoHandler
}

func (h *userInfoHandler) GetInfo(ctx context.Context, in *GetRequest, out *GetResponse) error {
	return h.UserInfoHandler.GetInfo(ctx, in, out)
}