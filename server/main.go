package main

import (
	"context"
	pb "go-micro-sentinel/server/proto"
	"log"
	"time"

	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/consul"

	"github.com/micro/go-micro"
)

type Start struct{}

func (g *Start) SendMessage(ctx context.Context, req *pb.CallRequest, rsp *pb.CallResponse) error {
	rsp.Msg = "this Respons " + req.Name
	return nil
}

func main() {
	cr := consul.NewRegistry(registry.Addrs("47.115.20.3:8500"))
	service := micro.NewService(
		micro.Name("go.micro.srv.send"),
		micro.RegisterTTL(time.Second*3),
		micro.RegisterInterval(time.Second*3),
		micro.Registry(cr),
	)

	service.Init()

	pb.RegisterStartHandler(service.Server(), new(Start))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}

}
