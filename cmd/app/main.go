package main

import (
	"flag"
	"os"
	"path"

	"nonoDemo/internal/infrastructure/client"
	"nonoDemo/internal/infrastructure/database/entity"
	"nonoDemo/pkg/adapters/agin"
	"nonoDemo/pkg/adapters/grpc"
	"nonoDemo/pkg/config"
	"nonoDemo/pkg/framework"
	"nonoDemo/pkg/utils/observability/log"
	"nonoDemo/pkg/utils/observability/metrics"
	"nonoDemo/pkg/utils/observability/tracing"
)

var pwd, _ = os.Getwd()
var rootdir = path.Dir(path.Dir(pwd))
var configFile = flag.String("f", path.Join(pwd, "configs/config.yaml"), "set config file which viper will loading.")
var envPrefix = "nonoDemo"

func main() {
	// 加载配置文件
	cfg := config.Config{}
	viperOption := &framework.ViperOption{CfPath: *configFile, EnvPrefix: ""}
	_, err := framework.LoadConfiguration(viperOption, &cfg)
	if err != nil {
		panic(err)
	}
	// 初始化logger
	logger := log.NewLogger(cfg.Log)
	// 创建metrics
	meter := metrics.NewMetrics(logger)
	// 初始化tracing
	tracing.InitTracer(&cfg, logger)
	// 初始化xorm
	dbEngine := client.NewXorm(cfg)
	// TODO 初始化redis 需要时注入
	//redisConn := cache.NewConnection(&cfg)
	// 初始化实现的数据持久化
	entity.Sync(dbEngine, logger)

	// 初始化gRPC-server
	server := NewGrpcServer(logger, dbEngine)
	server.WithOptions(grpc.NewOptionsFromConfig(&cfg).
		OptionTracing().
		OptionMetrics().
		OptionMiddleware(grpc.LogAccessMiddleware(cfg.Trace.ServiceName, logger)).
		OptionListenAddr(":9999"),
	)
	// 初始化ginOption
	ginServer := NewGinServer(logger, dbEngine).
		WithOptions(agin.NewOptionsFromConfig(&cfg).
			OptionListenAddr(":8080").
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
