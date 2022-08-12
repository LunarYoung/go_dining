package server

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	nacos "github.com/micro/go-plugins/registry/nacos/v2"
	helloworld "order/grpc/pb"
)

type Helloworld struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Helloworld) Hello(ctx context.Context, req *helloworld.HelloRequest, rsp *helloworld.HelloResponse) error {
	logger.Info("Received Helloworld.Call request")
	fmt.Println(req.Name)
	rsp.Greeting = "dd"
	return nil
}
func S() {
	addrs := make([]string, 1)
	addrs[0] = "139.159.182.159:8850"
	registry := nacos.NewRegistry(func(options *registry.Options) {
		options.Addrs = addrs
	})
	service := micro.NewService(
		// Set service name
		micro.Name("my.micro.service"),
		// Set service registry
		micro.Registry(registry),
	)
	helloworld.RegisterGreeterHandler(service.Server(), new(Helloworld))
	service.Run()
}
