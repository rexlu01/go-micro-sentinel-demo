package main

import (
	"context"
	pb "go-micro-sentinel/server/proto"
	"log"
	"time"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
)

type Start struct{}

func (g *Start) SendMessage(ctx context.Context, req *pb.CallRequest, rsp *pb.CallResponse) error {
	rsp.Msg = "this Respons " + req.Name
	return nil
}

func main() {
	url := []string{"47.115.20.3:8500"}
	cr := consul.NewRegistry(func(op *registry.Options) {
		op.Addrs = url
	})
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
