package client

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/nacos/v2"
	helloworld "user/grpc/pb"
)

const serverName = "my.micro.service"

func C() {
	addrs := make([]string, 1)
	addrs[0] = "139.159.182.159:8850"
	r := nacos.NewRegistry(func(options *registry.Options) {
		options.Addrs = addrs
	})
	// 定义服务，可以传入其它可选参数
	service := micro.NewService(
		micro.Name("my.micro.service.client"),
		micro.Registry(r))
	// 创建新的客户端
	greeter := helloworld.NewGreeterClient(serverName, service.Client())
	// 调用greeter
	rsp, err := greeter.Hello(context.TODO(), &helloworld.HelloRequest{Name: "John"})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(rsp.Greeting)

	go service.Run()

}
