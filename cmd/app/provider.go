package main

import (
	"github.com/google/wire"
	"nonoDemo/internal/adapters/controllers"
	"nonoDemo/pkg/adapters/agin"
	"nonoDemo/pkg/adapters/grpc"
	"nonoDemo/proto_gen/api/hello"
)

var Provider = wire.NewSet(
	hello.NewUserService,
)

func ProvideGrpcServices(userService *hello.UserService) []grpc.Instance {
	return []grpc.Instance{userService}
}

func ProvideController(helloController *controllers.HelloController) []agin.Controller {
	return []agin.Controller{
		helloController,
	}
}
