# csstools
# 配置加速器
```
export GOPROXY=https://goproxy.cn,direct
```

# snap安装protobuf 
```
sudo snap install protobuf --classic

```
# 云开发环境没有systemctl的，需要使用下面命令
# 安装protobuf
```
sudo apt install protobuf-compiler
```

# protobuf 新版本需要指定go_package
```
syntax = "proto3";

package hello;

option go_package = "heibai/proto"; // 指定生成的 Go 代码的包路径

// ...


```

# 安装protoc-gen-micro
```
go install github.com/asim/go-micro/cmd/protoc-gen-micro/v3@latest

```

# 添加环境变量
```
export PATH=$PATH:$(go env GOPATH)/bin

```
# 生成proto到api目录
```
protoc --proto_path=api \
    --go_out=api --go_opt=paths=source_relative \
    --micro_out=api --micro_opt=paths=source_relative \
    api/greeter/greeter.proto
```

# 启用 Go Modules: 在您的终端中运行以下命令来显式启用 Go Modules:
```
go env -w GO111MODULE=on

```
# 清理并更新依赖: 在项目根目录下运行以下命令:
```
go clean -modcache
go mod tidy
```
