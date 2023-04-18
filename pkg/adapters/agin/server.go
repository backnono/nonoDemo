package agin

import (
	"context"
	"nonoDemo/pkg/framework"
	"github.com/gin-gonic/gin"
	stdlog "log"
	"net/http"
)

type Server struct {
	logger      framework.Logger
	controllers []Controller
	server      *http.Server
	router      *gin.Engine
	options     Options
}

func NewServer(controllers []Controller, logger framework.Logger) *Server {
	return &Server{
		controllers: controllers,
		logger:      logger,
	}
}

func (s *Server) WithOptions(opt Options) *Server {
	s.options = opt
	return s
}

func (s *Server) Serve(ctx context.Context) error {
	go func() {
		<-ctx.Done()
		_ = s.logger.Log("msg", "ready to shutdown http server...")
		_ = s.server.Shutdown(ctx)
		_ = s.logger.Log("msg", "http server shutdown")
	}()
	if s.options.ListenAddr == "" {
		panic("server listen address can not be empty.")
	}
	// 使用logger覆盖默认的日志输出，保证日志输出格式一致
	_logger := newGinLoggerAdapter(s.logger)
	stdlog.SetOutput(_logger)
	stdlog.Default().SetOutput(_logger)
	gin.DefaultWriter = _logger
	gin.SetMode(gin.ReleaseMode)
	s.router = gin.Default()
	// 注册中间件
	for _, middleware := range s.options.middlewares {
		s.router.Use(middleware)
	}
	// 注册路由
	for _, ctrl := range s.controllers {
		ctrl.InitRouter(s.router)
	}
	for _, handler := range s.options.handlers {
		s.router.Handle(handler.method, handler.path, handler.handlerFunc)
	}

	s.server = &http.Server{
		Addr:    s.options.ListenAddr,
		Handler: s.router,
	}
	s.logger.Debug("start listen http on " + s.options.ListenAddr)
	// 启动服务监听
	return s.server.ListenAndServe()
}
