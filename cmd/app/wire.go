//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"xorm.io/xorm"

	"nonoDemo/internal"
	"nonoDemo/pkg/adapters/agin"
	"nonoDemo/pkg/adapters/grpc"
	"nonoDemo/pkg/framework"
)

func NewGrpcServer(logger framework.Logger, dbEngine *xorm.Engine) *grpc.Server {
	wire.Build(ProvideGrpcServices,
		internal.Provider,
		Provider,
		grpc.Provider)
	return &grpc.Server{}
}

func NewGinServer(logger framework.Logger, dbEngine *xorm.Engine) *agin.Server {
	wire.Build(
		//ViperProviderSet,
		//DevkitProvider,
		ProvideController,
		internal.Provider,
		agin.Provider)
	return &agin.Server{}
}

/*
var DevkitProvider = wire.NewSet(
	config.CfgProviderSet,
	client.NewXorm, //ORM 框架的相关组件
	cache.Provider,
)*/
