package main

import (
	"context"
	"fmt"
	pb "go-micro-sentinel/getuserinfo/proto"
	start "go-micro-sentinel/server/proto"


	sentinel "go-micro-sentinel/sentinal"

	"log"
	"time"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/consul"
)

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
	sentinel.InitSentinel()
	cr := consul.NewRegistry(registry.Addrs("47.115.20.3:8500"))
	service := micro.NewService(
		micro.Name("go.micro.srv.getuserinfo"),
		micro.RegisterTTL(time.Second*3),
		micro.RegisterInterval(time.Second*3),
		micro.Registry(cr),
		micro.WrapHandler(sentinel.NewSentinelHandlerWrapper()),
	)
	service.Init()

	cl = start.NewStartService("go.micro.srv.send", client.DefaultClient)

	pb.RegisterUserInfoHandler(service.Server(), new(UserInfo))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}

}
