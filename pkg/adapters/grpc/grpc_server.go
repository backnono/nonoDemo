package grpc

import (
	"google.golang.org/grpc"
)

type Instance interface {
	RegisterService(s *grpc.Server)
	WithOptions(options Options)
	Build()
}
