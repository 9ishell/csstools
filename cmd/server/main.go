package main

import (
    "log"
    "github.com/9ishell/csstools/service/greeter"

    "github.com/asim/go-micro/v3"
    proto "github.com/9ishell/csstools/api/greeter"
)

func main() {
    // 创建服务
    service := micro.NewService(
        micro.Name("greeter"),
    )
    service.Init()

    // 注册处理器
    proto.RegisterGreeterHandler(service.Server(), new(greeter.GreeterHandler))

    // 启动服务
    if err := service.Run(); err != nil {
        log.Fatal(err)
    }
}

