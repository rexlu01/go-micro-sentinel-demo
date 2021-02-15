package sentinel

import (
	"context"
	"log"

	sentinel_api "github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/base"
	"github.com/alibaba/sentinel-golang/core/flow"
	"github.com/micro/go-micro/errors"
	"github.com/micro/go-micro/server"
)

/*
time: 2020/6/10-10:38
author: waiwen
*/
//化sentinal
func InitSentinel() {
	err := sentinel_api.InitDefault()
	if err != nil {
		log.Fatalf("Unexpected error: %+v", err)
	}

	_, err = flow.LoadRules([]*flow.Rule{
		{
			Resource:         "list-limit",
			ControlBehavior:  flow.Reject,
			Threshold:        1000,
			StatIntervalInMs: 1000,
		},
	})
	if err != nil {
		log.Fatalf("Unexpected error: %+v", err)
		return
	}
}

//构建HandlerWrapper中间件
func NewSentinelHandlerWrapper() server.HandlerWrapper {
	return func(handlerFunc server.HandlerFunc) server.HandlerFunc {
		return func(ctx context.Context, req server.Request, rsp interface{}) error {
			//name := fmt.Sprintf("%s.%s", req.Service(), req.Endpoint())
			resourceName := "list-limit"
			entry, err := sentinel_api.Entry(
				resourceName,
				sentinel_api.WithResourceType(base.ResTypeRPC),
				sentinel_api.WithTrafficType(base.Inbound),
			)
			if err != nil {
				return errors.BadRequest("gk.micro.srv.circle", "too many requests")
			}
			defer entry.Exit()
			return handlerFunc(ctx, req, rsp)
		}
	}
}
