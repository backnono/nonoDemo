package grpc

import "context"

func DecodeRequest(_ context.Context, req interface{}) (interface{}, error) {
	return req, nil
}
func EncodeResponse(_ context.Context, rsp interface{}) (interface{}, error) {
	return rsp, nil
}
