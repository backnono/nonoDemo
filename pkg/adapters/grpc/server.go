package grpc

import (
	"context"
	"fmt"
	"github.com/google/wire"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"nonoDemo/pkg/framework"
)

type ServerOpt func(s *grpc.Server)

type Server struct {
	listener   net.Listener
	grpcServer *grpc.Server
	logger     framework.Logger
	servers    []Instance
	options    Options
}

func NewServer(logger framework.Logger, servers []Instance) *Server {
	return &Server{
		grpcServer: grpc.NewServer(),
		logger:     logger,
		servers:    servers,
	}
}

func (s *Server) registerServers() {
	for _, inst := range s.servers {
		inst.RegisterService(s.grpcServer)
		inst.WithOptions(s.options)
		inst.Build()
	}
}

func (s *Server) Serve(ctx context.Context) error {
	go func() {
		<-ctx.Done()
		_ = s.logger.Log("msg", "ready to shutdown service server...")
		s.grpcServer.GracefulStop()
		_ = s.logger.Log("msg", "service server shutdown")
	}()
	if s.options.ListenAddr == "" {
		panic("server listen address can not be empty.")
	}
	listener, err := net.Listen("tcp", s.options.ListenAddr)
	if err != nil {
		panic(err)
	}
	s.listener = listener
	reflection.Register(s.grpcServer)
	s.registerServers()
	s.logger.Debug(fmt.Sprintf("start listen grpc on %s", s.options.ListenAddr))
	return s.grpcServer.Serve(listener)
}

func (s *Server) WithOptions(opt Options) *Server {
	s.options = opt
	return s
}

var Provider = wire.NewSet(NewServer)
