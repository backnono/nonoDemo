package metrics

import (
	"context"
	"nonoDemo/pkg/framework"
	"nonoDemo/pkg/utils/observability/log"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	stdhttp "net/http"
	"regexp"
)

type Metrics struct {
	collector []prometheus.Collector
	server    *stdhttp.Server
	logger    framework.Logger
	handler   stdhttp.Handler
}

type OptionMetrics func(config *Metrics)

func OptionsCollectors(collectors ...prometheus.Collector) OptionMetrics {
	return func(config *Metrics) {
		config.collector = append(config.collector, collectors...)
	}
}

func (metrics *Metrics) Init(opts ...OptionMetrics) stdhttp.Handler {
	for _, opt := range opts {
		opt(metrics)
	}
	metrics.collector = append(metrics.collector,
		collectors.NewBuildInfoCollector(),

		collectors.NewGoCollector(
			collectors.WithGoCollectorRuntimeMetrics(collectors.GoRuntimeMetricsRule{Matcher: regexp.MustCompile(".*")}),
			// collectors.WithGoCollections(
			// 	collectors.GoRuntimeMetricsCollection,
			// )
		),
		collectors.NewProcessCollector(
			collectors.ProcessCollectorOpts{
				PidFn:        nil,
				Namespace:    "",
				ReportErrors: false,
			},
		),
	)
	set := map[prometheus.Collector]struct{}{}
	for _, collector := range metrics.collector {
		set[collector] = struct{}{}
	}
	metrics.collector = nil
	for k, _ := range set {
		metrics.collector = append(metrics.collector, k)
	}
	registry := prometheus.NewRegistry()
	registry.MustRegister(metrics.collector...)
	metrics.handler = promhttp.HandlerFor(registry, promhttp.HandlerOpts{})
	return metrics.handler
}

func (metrics *Metrics) Serve(ctx context.Context) error {
	go func() {
		<-ctx.Done()
		_ = metrics.logger.Log("msg", "ready to shutdown service server...")
		_ = metrics.server.Shutdown(ctx)
		_ = metrics.logger.Log("msg", "service server shutdown")
	}()
	metrics.server = &stdhttp.Server{
		Addr:    ":8888",
		Handler: metrics.handler,
	}
	metrics.logger.Debug("metrics server listen on :8888")
	return metrics.server.ListenAndServe()
}

func NewMetrics(logger framework.Logger) *Metrics {
	_meter = &Metrics{
		logger: logger,
	}
	return _meter
}

var _meter *Metrics

func GlobalMeter() *Metrics {
	if _meter != nil {
		return _meter
	}
	return NewMetrics(log.GlobalLogger())
}
