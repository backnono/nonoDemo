package internal

import (
	"github.com/google/wire"
	"nonoDemo/internal/adapters/controllers"
	"nonoDemo/internal/adapters/grpcsvc"
	"nonoDemo/proto_gen/api/hello"
)

//nolint:all
var Provider = wire.NewSet(
	grpcsvc.NewUserUsecase,
	controllers.NewHelloController,
	wire.Bind(new(hello.UserServiceServer), new(*grpcsvc.UserUsecase)))
