package main

import (
	"context"
	"log"

	proto "github.com/9ishell/csstools/api/greeter"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/registry"
	consul "github.com/go-micro/plugins/v3/registry/consul"
)

func main() {
	// 使用 Consul 作为注册中心
	reg := consul.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{
			"127.0.0.1:8500", // 这里填写您的 Consul 地址
		}
	})

	// 创建服务
	service := micro.NewService(
		micro.Name("greeter.client"),
		micro.Registry(reg),
	)
	service.Init()

	// 创建客户端
	client := proto.NewGreeterService("greeter", service.Client())

	rsp, err := client.SayHello(context.Background(), &proto.HelloRequest{Name: "World"})
	if err != nil {
		log.Fatal(err)
	}

	log.Println(rsp.Greeting)
}
