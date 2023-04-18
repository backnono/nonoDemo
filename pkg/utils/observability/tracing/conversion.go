package tracing

import (
	"context"
	"nonoDemo/pkg/framework"
	"nonoDemo/pkg/utils/observability"
	"fmt"
	httptransport "github.com/go-kit/kit/transport/http"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc/metadata"
	"net/http"
	"strings"
)

func ContextToGRPC(tp trace.TracerProvider, logger framework.Logger) func(ctx context.Context, md *metadata.MD) context.Context {
	return func(ctx context.Context, md *metadata.MD) context.Context {
		if span := trace.SpanFromContext(ctx); span != nil {
			spanContext := span.SpanContext()
			traceID := spanContext.TraceID
			spanID := spanContext.SpanID
			md.Set(TraceIDKey, fmt.Sprintf("%s:%s:%s:%s", traceID().String(), spanID().String(), "00000000", "1"))
		}
		return ctx
	}
}

func GRPCToContext(logger framework.Logger) func(ctx context.Context, md metadata.MD) context.Context {
	return func(ctx context.Context, md metadata.MD) context.Context {
		traceID := md.Get(TraceIDKey)
		if traceID == nil {
			return ctx
		}
		return extractTraceID(ctx, traceID[0], logger)
	}
}

func HTTPToContext(log framework.Logger) func(context.Context, *http.Request) context.Context {
	return func(ctx context.Context, request *http.Request) context.Context {
		header := request.Header.Get(TraceIDKey)
		if header == "" {
			return ctx
		}
		return extractTraceID(ctx, header, log)
	}
}

func ContextToHTTP(tp trace.TracerProvider, logger framework.Logger) httptransport.RequestFunc {
	return func(ctx context.Context, request *http.Request) context.Context {
		if span := trace.SpanFromContext(ctx); span != nil {
			spanContext := span.SpanContext()
			traceID := spanContext.TraceID
			spanID := spanContext.SpanID
			request.Header.Set(TraceIDKey, fmt.Sprintf("%s:%s:%s:%s", traceID().String(), spanID().String(), "00000000", "1"))
		}
		return ctx
	}
}

func extractTraceID(ctx context.Context, fullID string, logger framework.Logger) context.Context {
	ids := strings.Split(fullID, ":")
	traceID, spanID := ids[0], ids[1]
	if len(traceID) < 32 {
		traceID = "0000000000000000" + traceID
	}

	tid, err := trace.TraceIDFromHex(traceID)
	if err != nil {
		_ = logger.Log("error", err)
		return ctx
	}
	sid, err := trace.SpanIDFromHex(spanID)
	if err != nil {
		_ = logger.Log("error", err)
		return ctx
	}
	spanContext := trace.NewSpanContext(trace.SpanContextConfig{
		SpanID:     sid,
		TraceID:    tid,
		TraceFlags: 1,
		TraceState: trace.TraceState{},
		Remote:     true,
	})
	ctx = observability.AppendEvent(ctx, "trace.id", traceID)
	ctx = trace.ContextWithSpanContext(ctx, spanContext)
	return ctx
}
