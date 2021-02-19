package main

import (
	"context"
	"log"
	"time"

	pb "go-micro-sentinel/getuserinfo/proto"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/web"

	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
)

type UserInfo struct {
}

var (
	cl pb.UserInfoService
)

func (g *UserInfo) Anything(c *gin.Context) {
	log.Print("Received Say.Anything API request")
	c.JSON(200, map[string]string{
		"message": "Hi, this is the Greeter API",
	})
}

func (g *UserInfo) GetInfo(c *gin.Context) {
	log.Print("Received Say.Hello API request")
	name := c.Param("name")

	response, err := cl.GetInfo(context.TODO(), &pb.GetRequest{
		Username: name,
	})

	if err != nil {
		c.JSON(500, err)
	}

	c.JSON(200, response)
}

func main() {
	url := []string{"47.115.20.3:8500"}
	cr := consul.NewRegistry(func(op *registry.Options) {
		op.Addrs = url
	})
	service := web.NewService(
		web.Name("go.micro.api.sendmessage"),
		web.Registry(cr),
		web.RegisterTTL(time.Second*3),
		web.RegisterInterval(time.Second*3),
	)

	service.Init()

	cl = pb.NewUserInfoService("go.micro.srv.getuserinfo", client.DefaultClient)

	start := new(UserInfo)

	router := gin.Default()
	router.GET("/sendmessage", start.Anything)
	router.GET("/sendmessage/:name", start.GetInfo)

	service.Handle("/", router)

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}

}
