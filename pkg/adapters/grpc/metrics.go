package grpc

import (
	"context"
	"nonoDemo/pkg/utils/observability"
	"github.com/go-kit/kit/endpoint"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"google.golang.org/grpc/status"
	"strconv"
	"time"
)

var InvokeCounter = promauto.NewCounterVec(prometheus.CounterOpts{
	Name: "grpc_request_total",
	Help: "grpc request total",
}, []string{"uri", "error", "error_code"})

var LatencySummary = promauto.NewSummaryVec(prometheus.SummaryOpts{
	Name:       "grpc_request_latency",
	Help:       "the summary of grpc request latency",
	Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.95: 0.005, 0.99: 0.001},
}, []string{"uri"})

func EndpointInvokeCounterMiddleware() endpoint.Middleware {
	return func(e endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			start := time.Now()
			response, err = e(ctx, request)
			end := time.Now()
			latency := float64(end.Sub(start).Nanoseconds()) / 1000000
			observability.AppendEvents(ctx, EventKeyGrpcLatency, latency)
			errExist := "false"
			code := "UNKNOWN"
			if err != nil {
				errExist = "true"
				code = getBizErrCode(err)
			}
			InvokeCounter.WithLabelValues(observability.GetValue(ctx, EventKeyGrpcMethod).(string), errExist, code).Inc()
			LatencySummary.WithLabelValues(observability.GetValue(ctx, EventKeyGrpcMethod).(string)).
				Observe(observability.GetValue(ctx, EventKeyGrpcLatency).(float64))
			return response, err
		}
	}
}

func getBizErrCode(err error) string {
	if s, ok := status.FromError(err); ok {
		return strconv.Itoa((int)(s.Code()))
	}
	return "UNKNOWN"
}
