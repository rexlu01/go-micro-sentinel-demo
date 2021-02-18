package main

import (
	"context"
	"errors"
	"fmt"
	pb "go-micro-sentinel/getuserinfo/proto"
	start "go-micro-sentinel/server/proto"

	sentinelGo "go-micro-sentinel/sentinel"

	"log"
	"time"

	"github.com/alibaba/sentinel-golang/core/base"
	"github.com/micro/go-grpc/client"

	"github.com/micro/go-micro/v2/server"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
)

const FakeErrorMsg = "fake error for testing"

type UserInfo struct {
}

var (
	cl start.StartService
)

func (u *UserInfo) GetInfo(ctx context.Context, req *pb.GetRequest, rsp *pb.GetResponse) error {
	res, err := cl.SendMessage(context.TODO(), &start.CallRequest{
		Name: req.Username,
	})
	if err != nil {
		fmt.Println(err)
	}

	rsp.Msg = "this GetInfo respones " + res.Msg
	return nil
}

func main() {
	url := []string{"47.115.20.3:8500"}
	cr := consul.NewRegistry(func(op *registry.Options) {
		op.Addrs = url
	})
	service := micro.NewService(
		micro.Name("go.micro.srv.getuserinfo"),
		micro.RegisterTTL(time.Second*3),
		micro.RegisterInterval(time.Second*3),
		micro.Registry(cr),
		micro.WrapHandler(sentinelGo.NewHandlerWrapper(
			// add custom fallback function to return a fake error for assertion
			sentinelGo.WithServerBlockFallback(
				func(ctx context.Context, request server.Request, blockError *base.BlockError) error {
					return errors.New(FakeErrorMsg)
				}),
		)),
	)

	service.Init()

	cl = start.NewStartService("go.micro.srv.send", client.DefaultClient)

	pb.RegisterUserInfoHandler(service.Server(), new(UserInfo))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}

}
