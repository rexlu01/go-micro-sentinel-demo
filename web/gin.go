package main

import (
	"context"
	"log"
	"time"
	pb "twomicroexercise/03/server/proto"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/web"

	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/consul"
)

type Start struct{}

var (
	cl pb.StartService
)

func (g *Start) Anything(c *gin.Context) {
	log.Print("Received Say.Anything API request")
	c.JSON(200, map[string]string{
		"message": "Hi, this is the Greeter API",
	})
}

func (g *Start) SendMessage(c *gin.Context) {
	log.Print("Received Say.Hello API request")
	name := c.Param("name")

	response, err := cl.SendMessage(context.TODO(), &pb.CallRequest{
		Name: name,
	})

	if err != nil {
		c.JSON(500, err)
	}

	c.JSON(200, response)
}

func main() {
	cr := consul.NewRegistry(registry.Addrs("47.115.20.3:8500"))
	service := web.NewService(
		web.Name("go.micro.api.sendmessage"),
		web.Registry(cr),
		web.RegisterTTL(time.Second*3),
		web.RegisterInterval(time.Second*3),
	)

	service.Init()

	cl = pb.NewStartService("go.micro.srv.send", client.DefaultClient)

	start := new(Start)

	router := gin.Default()
	router.GET("/sendmessage", start.Anything)
	router.GET("/sendmessage/:name", start.SendMessage)

	service.Handle("/", router)

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}

}
