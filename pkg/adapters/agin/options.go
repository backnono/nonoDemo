package agin

import (
	"nonoDemo/pkg/framework"
	"nonoDemo/pkg/utils"
	"nonoDemo/pkg/utils/observability/log"
	"nonoDemo/pkg/utils/observability/metrics"
	"github.com/gin-gonic/gin"
)

type customHandler struct {
	method      string
	path        string
	handlerFunc gin.HandlerFunc
}

type Options struct {
	ListenAddr     string `json:"listen_addr" yaml:"listen_addr" mapstructure:"listen_addr"`
	TraceEnabled   bool   `json:"trace_enabled" yaml:"trace_enabled" mapstructure:"trace_enabled"`
	MetricsEnabled bool   `json:"metrics_enabled" yaml:"metrics_enabled" mapstructure:"metrics_enabled"`
	middlewares    []gin.HandlerFunc
	handlers       []customHandler
}

func NewOptions() Options {
	return Options{}
}

func NewOptionsFromConfig(config framework.Configuration) Options {
	options := utils.LoadConfig(config, Options{}).(Options)
	if options.TraceEnabled {
		options = options.OptionTrace()
	}
	if options.MetricsEnabled {
		options = options.OptionMetrics()
	}
	return options
}

func (opt Options) OptionListenAddr(addr string) Options {
	opt.ListenAddr = addr
	return opt
}

func (opt Options) OptionTrace() Options {
	opt.TraceEnabled = true
	opt.middlewares = append(opt.middlewares, TraceMiddleware(log.GlobalLogger()))
	return opt
}

func (opt Options) OptionMetrics() Options {
	opt.MetricsEnabled = true
	metrics.OptionsCollectors(HTTPEndpointInvokeCounter)(metrics.GlobalMeter())
	metrics.OptionsCollectors(HTTPEndpointLatency)(metrics.GlobalMeter())
	opt.middlewares = append(opt.middlewares, MetricsMiddleware())
	return opt
}

func (opt Options) OptionsMiddlewares(middlewares ...gin.HandlerFunc) Options {
	opt.middlewares = append(opt.middlewares, middlewares...)
	return opt
}

func (opt Options) OptionCustomHandler(method, path string, handlerFunc gin.HandlerFunc) Options {
	opt.handlers = append(opt.handlers, customHandler{
		method:      method,
		path:        path,
		handlerFunc: handlerFunc,
	})
	return opt
}
