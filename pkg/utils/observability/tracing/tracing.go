package tracing

import (
	"github.com/google/uuid"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
	"go.opentelemetry.io/otel/trace"
	stdlog "log"
	"nonoDemo/pkg/framework"
	"nonoDemo/pkg/utils"
	"nonoDemo/pkg/utils/tools"
	"os"
)

func init() {
	clientUUID = uuid.New().String()
}

const Version = "0.1.0"

var clientUUID string
var TracerProvider trace.TracerProvider

func initTracing(tracingConfig Config, logger framework.Logger) (trace.TracerProvider, error) {
	hostname, _ := os.Hostname()
	ip := tools.GetLocalIP()
	// Create the Jaeger exporter
	var exp *jaeger.Exporter
	var err error
	exp, err = jaeger.New(jaeger.WithAgentEndpoint(jaeger.WithAgentHost(tracingConfig.AgentHost),
		jaeger.WithAgentPort(tracingConfig.AgentPort),
		jaeger.WithLogger(stdlog.New(newLogger(logger), "", 0)),
	))
	if err != nil {
		return nil, err
	}
	var res *resource.Resource
	if os.Getenv("ENV") != "" {
		res = resource.NewWithAttributes(semconv.SchemaURL,
			semconv.ServiceNameKey.String(tracingConfig.ServiceName),
			attribute.String("hostname", hostname),
			attribute.String("ip", ip),
			attribute.String("client-uuid", clientUUID),
			attribute.String("ENV", os.Getenv("ENV")),
		)
	} else {
		res = resource.NewWithAttributes(semconv.SchemaURL,
			semconv.ServiceNameKey.String(tracingConfig.ServiceName),
			attribute.String("hostname", hostname),
			attribute.String("ip", ip),
			attribute.String("client-uuid", clientUUID),
		)
	}

	TracerProvider = tracesdk.NewTracerProvider(
		// Always be sure to batch in production.
		tracesdk.WithBatcher(exp),
		// Record information about this application in a Resource.
		tracesdk.WithResource(res),
		tracesdk.WithSampler(tracesdk.TraceIDRatioBased(tracingConfig.SamplingRate)),
	)
	return TracerProvider, nil
}

func InitTracer(config framework.Configuration, logger framework.Logger) trace.TracerProvider {
	cfg := utils.LoadConfig(config, Config{}).(Config)
	tp, err := initTracing(cfg, logger)
	if err != nil {
		panic(err)
	}
	return tp
}

func GlobalTraceProvider() trace.TracerProvider {
	return TracerProvider
}
