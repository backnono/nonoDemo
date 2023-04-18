package grpc

import (
	"context"
	"google.golang.org/grpc"
)

func getTransmissionMethod(ctx context.Context) string {
	stream := grpc.ServerTransportStreamFromContext(ctx)
	method := stream.Method()
	return method
}
