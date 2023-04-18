package grpc

import (
	"context"
	"nonoDemo/pkg/framework"
	"nonoDemo/pkg/utils/observability/tracing"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/tracing/opentracing"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	stdopentracing "github.com/opentracing/opentracing-go"
	"go.opentelemetry.io/otel/trace"
	stdgrpc "google.golang.org/grpc"
)

type Client struct {
	conn        *stdgrpc.ClientConn
	logger      framework.Logger
	tracer      stdopentracing.Tracer
	tp          trace.TracerProvider
	options     []func(string) grpctransport.ClientOption
	middlewares []func(endpoint.Endpoint) endpoint.Endpoint
}

func NewClient(conn *stdgrpc.ClientConn, logger framework.Logger) *Client {
	return &Client{conn: conn, logger: logger}
}

func (c *Client) WithTracing(tracer stdopentracing.Tracer) *Client {
	c.tracer = tracer
	return c
}

func (c *Client) WithTracerProvider(tp trace.TracerProvider) *Client {
	c.tp = tp
	return c
}

func (c *Client) WithOption(options ...func(string) grpctransport.ClientOption) *Client {
	c.options = options
	return c
}

func (c *Client) WithMiddleware(middlewares ...func(endpoint.Endpoint) endpoint.Endpoint) *Client {
	c.middlewares = middlewares
	return c
}

func (c *Client) Invoke(ctx context.Context, service, method string, request, reply interface{}) (interface{}, error) {
	var options []grpctransport.ClientOption
	for _, option := range c.options {
		options = append(options, option(service+"/"+method))
	}
	if c.tracer != nil {
		options = append(options, grpctransport.ClientBefore(opentracing.ContextToGRPC(c.tracer, c.logger)))
	}
	if c.tp != nil {
		options = append(options, grpctransport.ClientBefore(tracing.ContextToGRPC(c.tp, c.logger)))
	}
	ep := grpctransport.NewClient(c.conn,
		service,
		method,
		encodeRequest,
		decodeResponse,
		reply,
		options...).Endpoint()
	for _, middleware := range c.middlewares {
		ep = middleware(ep)
	}

	if c.tracer != nil {
		ep = opentracing.TraceClient(c.tracer, service+"/"+method)(ep)
	}
	if c.tp != nil {
		ep = TraceClientMiddleware(c.tp, service+"/"+method)(ep)
	}
	return ep(ctx, request)
}

func (c *Client) Close() error {
	return c.conn.Close()
}

func encodeRequest(ctx context.Context, request interface{}) (interface{}, error) {
	return request, nil
}

func decodeResponse(ctx context.Context, response interface{}) (interface{}, error) {
	return response, nil
}
