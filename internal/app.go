package internal

import (
	"github.com/google/wire"
	"nonoDemo/internal/adapters/controllers"
	"nonoDemo/internal/adapters/grpcsvc"
	"nonoDemo/internal/application/executor"
	"nonoDemo/internal/domain/gateway"
	"nonoDemo/internal/infrastructure/database"
	"nonoDemo/proto_gen/api/hello"
)

//nolint:all
var Provider = wire.NewSet(
	userProvider,
	helloProvider,
	blogProvider,
)

var userProvider = wire.NewSet(
	grpcsvc.NewUserUsecase,
	wire.Bind(new(hello.UserServiceServer), new(*grpcsvc.UserUsecase)),
)

var helloProvider = wire.NewSet(
	controllers.NewHelloController,
)

var blogProvider = wire.NewSet(
	controllers.NewBlogController,
	database.NewBlogRepository,
	executor.NewBlogOperator,
	wire.Bind(new(gateway.IBlogManager), new(*database.BlogRepository)),
)
