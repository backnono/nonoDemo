package main

import (
	"nonoDemo/pkg/adapters/agin"
	"nonoDemo/pkg/adapters/grpc"
	"nonoDemo/pkg/framework"
	"nonoDemo/pkg/utils/observability/log"
	"nonoDemo/pkg/utils/observability/metrics"
	"nonoDemo/pkg/utils/observability/tracing"
)

func main() {
	// 加载配置文件
	cfg := Config{}
	err := framework.LoadConfiguration("./configs/config.yaml", "", &cfg)
	if err != nil {
		panic(err)
	}
	// 初始化logger
	logger := log.NewLogger(cfg.Log)
	// 创建metrics
	meter := metrics.NewMetrics(logger)
	// 初始化tracing
	tracing.InitTracer(&cfg, logger)
	
	// 初始化gRPC-server
	server := NewGrpcServer(logger)
	server.WithOptions(grpc.NewOptionsFromConfig(&cfg).
		OptionTracing().
		OptionMetrics().
		OptionMiddleware(grpc.LogAccessMiddleware(cfg.Trace.ServiceName, logger)).
		OptionListenAddr(":9999"),
	)
    // 初始化ginOption
	ginServer := NewGinServer(logger).
		WithOptions(agin.NewOptionsFromConfig(&cfg).
			OptionListenAddr(":80").
			OptionsMiddlewares(agin.LogAccessMiddleware(cfg.Trace.ServiceName, logger)).
			OptionMetrics().
			OptionTrace(),
		)
	// Important: 需要在所有得server option 完成之后进行meter的初始化，
	// 初始换完成之后将会返回一个http.Handler，使用该Handler处理请求
	meter.Init()
	app := framework.NewApp(logger)
	app.WithServer(
		server,
		ginServer,
		meter,
	)
	err = app.Start()
	if err != nil {
		logger.Error("app exited", err)
	}
}
    