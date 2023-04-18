package grpc

import (
	"context"
	"nonoDemo/pkg/framework"
	"nonoDemo/pkg/utils/observability"
	"github.com/go-kit/kit/endpoint"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	EventKeyGrpcMethod  = "grpc.method"
	EventKeyGrpcLatency = "grpc.latency"
)

func LogAccessMiddleware(serviceName string, logger framework.Logger) endpoint.Middleware {
	return func(e endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			method := getTransmissionMethod(ctx)
			ctx = observability.AppendEvents(ctx, "service.name", serviceName, EventKeyGrpcMethod, method)
			defer func() {
				event := observability.GetEvent(ctx)
				if err != nil {
					st, _ := status.FromError(err)
					event.Append(framework.EventKeyBizCode, st.Code(), framework.EventKeyErrMsg, st.Message())
				}
				logger.Info("gRPC Access Log", event.EventToArgList()...)
			}()
			response, err = e(ctx, request)
			// 处理错误
			if err != nil {
				if appErr, ok := err.(framework.AppError); ok {
					err = status.Error(codes.Code(framework.GetType(appErr)), appErr.Error())
				} else if _, ok := status.FromError(err); ok {
					// do nothing
				} else {
					err = status.Error(codes.Code(framework.CommonInternalErr), appErr.Error())
				}
			}
			return
		}
	}
}
