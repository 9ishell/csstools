package main

import (
	"context"
	"log"

	proto "github.com/9ishell/csstools/api/greeter"
	"github.com/asim/go-micro/v3"
)

func main() {
	// 创建服务
	service := micro.NewService(micro.Name("greeter.client"))
	service.Init()

	// 创建客户端
	client := proto.NewGreeterService("greeter", service.Client())

	rsp, err := client.SayHello(context.Background(), &proto.HelloRequest{Name: "World"})
	if err != nil {
		log.Fatal(err)
	}

	log.Println(rsp.Greeting)
}
