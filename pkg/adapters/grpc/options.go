package grpc

import (
	"nonoDemo/pkg/framework"
	"nonoDemo/pkg/utils"
	"nonoDemo/pkg/utils/observability/metrics"
	"github.com/go-kit/kit/endpoint"
)

type Options struct {
	TraceEnable   bool   `yaml:"trace_enabled" json:"trace_enable" mapstructure:"trace_enabled"`
	MetricsEnable bool   `yaml:"metrics_enable" json:"metrics_enable" mapstructure:"metrics_enabled"`
	ListenAddr    string `yaml:"listen_addr" json:"listen_addr" mapstructure:"listen_addr"`
	middlewares   []func(endpoint.Endpoint) endpoint.Endpoint
}

func NewOptions() Options {
	return Options{}
}

func NewOptionsFromConfig(config framework.Configuration) Options {
	options := utils.LoadConfig(config, Options{}).(Options)
	if options.MetricsEnable {
		options = options.OptionMetrics()
	}
	return options
}

func (opt Options) OptionTracing() Options {
	opt.TraceEnable = true
	return opt
}

func (opt Options) OptionMetrics() Options {
	opt.MetricsEnable = true
	metrics.OptionsCollectors(InvokeCounter)(metrics.GlobalMeter())
	metrics.OptionsCollectors(LatencySummary)(metrics.GlobalMeter())
	opt.middlewares = append(opt.middlewares, EndpointInvokeCounterMiddleware())
	return opt
}

func (opt Options) OptionMiddleware(middlewares ...func(endpoint.Endpoint) endpoint.Endpoint) Options {
	opt.middlewares = append(opt.middlewares, middlewares...)
	return opt
}

func (opt Options) OptionListenAddr(addr string) Options {
	opt.ListenAddr = addr
	return opt
}

func (opt Options) Trace() bool {
	return opt.TraceEnable
}

func (opt Options) Metrics() bool {
	return opt.MetricsEnable
}

func (opt Options) Middlewares() []func(endpoint.Endpoint) endpoint.Endpoint {
	return opt.middlewares
}
