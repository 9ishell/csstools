package greeter

import (
	"context"

	"github.com/9ishell/csstools/api/greeter"
)

type GreeterHandler struct{}

func (g *GreeterHandler) SayHello(ctx context.Context, req *greeter.HelloRequest, rsp *greeter.HelloResponse) error {
	rsp.Greeting = "Hello222 " + req.Name
	return nil
}
