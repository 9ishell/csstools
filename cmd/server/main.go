package main

import (
	"log"

	"github.com/9ishell/csstools/service/greeter"

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
		micro.Name("greeter"),
		micro.Registry(reg), // 设置注册中心为 Consul
	)
	service.Init()

	// 注册处理器
	proto.RegisterGreeterHandler(service.Server(), new(greeter.GreeterHandler))

	// 启动服务
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
