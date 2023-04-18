package grpc

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"nonoDemo/pkg/utils/observability"
	"nonoDemo/pkg/utils/observability/tracing"
)

func TraceServerMiddleware(tp trace.TracerProvider) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			var opName string
			if observability.GetValue(ctx, EventKeyGrpcMethod) == nil {
				opName = getTransmissionMethod(ctx)
				ctx = observability.AppendEvents(ctx, EventKeyGrpcMethod, opName)
			} else {
				opName = observability.GetValue(ctx, EventKeyGrpcMethod).(string)
			}
			opName = fmt.Sprintf("%s : %s", "GRPC", opName)
			tracer := tp.Tracer(tracing.TracerName, trace.WithInstrumentationVersion(tracing.Version))
			opts := []trace.SpanStartOption{
				trace.WithSpanKind(trace.SpanKindServer),
			}
			ctx, span := tracer.Start(ctx, opName, opts...)
			defer span.End()
			defer func() {
				if err != nil {
					span.RecordError(err)
					span.SetStatus(codes.Error, err.Error())
					return
				}
				if res, ok := response.(endpoint.Failer); ok && res.Failed() != nil {
					span.RecordError(res.Failed())
					span.SetStatus(codes.Error, res.Failed().Error())
					return
				}
			}()
			response, err = next(ctx, request)
			return
		}
	}
}

func TraceClientMiddleware(tp trace.TracerProvider, operation string) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			tracer := tp.Tracer(tracing.TracerName, trace.WithInstrumentationVersion(tracing.Version))
			opts := []trace.SpanStartOption{
				trace.WithSpanKind(trace.SpanKindClient),
			}
			ctx, span := tracer.Start(ctx, operation, opts...)
			defer span.End()
			defer func() {
				if err != nil {
					// generic error
					span.RecordError(err)
					span.SetStatus(codes.Error, err.Error())

					return
				}

				// Test for business error. Business errors are often
				// successful requests carrying a business failure that
				// the client can act upon and therefore do not count
				// as failed requests.
				if res, ok := response.(endpoint.Failer); ok && res.Failed() != nil {
					span.RecordError(res.Failed())
					span.SetStatus(codes.Error, res.Failed().Error())
					return
				}

				// no errors identified
			}()
			ctx = context.WithValue(ctx, "span", span)
			response, err = next(ctx, request)
			return
		}
	}
}
