package agin

import (
	"nonoDemo/pkg/framework"
	"nonoDemo/pkg/utils/observability"
	"nonoDemo/pkg/utils/observability/tracing"
	"nonoDemo/pkg/utils/tools"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"strconv"
	"time"
)

const (
	EventKeyHTTPMethod  = "http.method"
	EventKeyHTTPUrl     = "http.url"
	EventKeyHTTPStatus  = "http.status"
	EventKeyHTTPLatency = "http.latency"
)

// LogAccess [http.method,http.url]
//                                  => metrics []

func LogAccessMiddleware(ServiceName string, logger framework.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Request = c.Request.WithContext(observability.AppendEvents(c.Request.Context(),
			"server.name", ServiceName,
			EventKeyHTTPMethod, c.Request.Method,
			EventKeyHTTPUrl, c.Request.URL.String()))
		defer func() {
			event := observability.GetEvent(c.Request.Context())
			logger.Info("Access Log", event.EventToArgList()...)
		}()
		c.Next()
		_ = observability.AppendEvents(c.Request.Context(), "server.node", tools.GetLocalIP(),
			"server.name", ServiceName,
			EventKeyHTTPMethod, c.Request.Method,
			EventKeyHTTPUrl, c.Request.URL.String(),
			EventKeyHTTPStatus, c.Writer.Status(),
			EventKeyHTTPLatency, float64(time.Now().Sub(start).Nanoseconds())/1000000)
	}
}

func MetricsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		end := time.Now()
		latency := float64(end.Sub(start).Nanoseconds()) / 1000000
		observability.AppendEvent(c.Request.Context(), EventKeyHTTPLatency, latency)
		var url, method, status, bizCode interface{}
		bizCode = "0"
		event := observability.GetEvent(c.Request.Context())
		if event == nil ||
			event.GetValue(EventKeyHTTPUrl) == nil ||
			event.GetValue(EventKeyHTTPMethod) == nil ||
			event.GetValue(EventKeyHTTPStatus) == nil {
			url = c.Request.URL.String()
			method = c.Request.Method
			status = strconv.Itoa(c.Writer.Status())
			observability.AppendEvents(c.Request.Context(),
				EventKeyHTTPStatus, status,
				EventKeyHTTPUrl, url,
				EventKeyHTTPMethod, method)
		} else {
			url = event.GetValue(EventKeyHTTPUrl).(string)
			method = event.GetValue(EventKeyHTTPMethod).(string)
			status = strconv.Itoa(event.GetValue(EventKeyHTTPStatus).(int))
		}
		if event.GetValue(framework.EventKeyBizCode) != nil {
			bizCode = strconv.FormatUint(uint64(event.GetValue(framework.EventKeyBizCode).(framework.ErrorType)), 10)
		}
		HTTPEndpointInvokeCounter.WithLabelValues(url.(string), method.(string), status.(string), bizCode.(string)).Inc()
		HTTPEndpointLatency.WithLabelValues(url.(string), method.(string)).Observe(latency)
	}
}

func TraceMiddleware(logger framework.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := tracing.HTTPToContext(logger)(c.Request.Context(), c.Request)
		event := observability.GetEvent(ctx)
		var httpMethod, httpUrl interface{}
		if event == nil || event.GetValue(EventKeyHTTPMethod) == nil || event.GetValue(EventKeyHTTPUrl) == nil {
			httpMethod = c.Request.Method
			httpUrl = c.Request.URL.String()
			observability.AppendEvents(ctx, EventKeyHTTPMethod, httpMethod, EventKeyHTTPUrl, httpUrl)
		} else {
			httpMethod = event.GetValue(EventKeyHTTPMethod)
			httpUrl = event.GetValue(EventKeyHTTPUrl)
		}
		var err error
		opName := fmt.Sprintf("HTTP %s : %s", httpMethod, httpUrl)
		tracer := tracing.GlobalTraceProvider().Tracer(tracing.TracerName, trace.WithInstrumentationVersion(tracing.Version))
		opts := []trace.SpanStartOption{
			trace.WithSpanKind(trace.SpanKindServer),
		}
		ctx, span := tracer.Start(ctx, opName, opts...)
		c.Request = c.Request.WithContext(ctx)
		defer span.End()
		defer func() {
			if err != nil {
				// generic error
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
				return
			}
		}()
		c.Next()
		observability.AppendEvents(ctx, EventKeyHTTPStatus, c.Writer.Status())
		if c.Writer.Status() > 399 {
			err = errors.New("http response status error: " + strconv.Itoa(c.Writer.Status()))
		}
	}
}

var HTTPEndpointInvokeCounter = promauto.NewCounterVec(prometheus.CounterOpts{
	Name: "http_request_total",
	Help: "http request total",
}, []string{"uri", "method", "status", "biz_code"})

var HTTPEndpointLatency = promauto.NewSummaryVec(prometheus.SummaryOpts{
	Name: "http_request_latency",
	Help: "the summary of http request latency",
	Objectives: map[float64]float64{
		0.5:  0.05,
		0.9:  0.01,
		0.95: 0.005,
		0.99: 0.001},
}, []string{"uri", "method"})
