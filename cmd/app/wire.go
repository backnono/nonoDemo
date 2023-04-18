//go:build wireinject
// +build wireinject

package main

import (
	"nonoDemo/internal"
	"nonoDemo/pkg/adapters/agin"
	"nonoDemo/pkg/adapters/grpc"
	"nonoDemo/pkg/framework" 
	"github.com/google/wire"
)

func NewGrpcServer(logger framework.Logger) *grpc.Server {
	wire.Build(ProvideGrpcServices,
		internal.Provider,
		Provider,
		grpc.Provider)
	return &grpc.Server{}
}

func NewGinServer(logger framework.Logger) *agin.Server {
	wire.Build(ProvideController,
		internal.Provider,
		agin.Provider)
	return &agin.Server{}
}
