// Package grpcsvc @Author nono.he 2023/4/14 14:35:00
package grpcsvc

import (
	"context"
	"fmt"
	"google.golang.org/grpc/metadata"
	v1 "nonoDemo/proto_gen/api/metadata/v1"
)

type MetadataUsecase struct {
	v1.UnimplementedMetaTestServiceServer
}

func NewMetadataUsecase() *MetadataUsecase {
	return &MetadataUsecase{}
}

func (u *MetadataUsecase) Create(ctx context.Context, in *v1.Req) (*v1.Response, error) {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		fmt.Println(md)
	}

	return &v1.Response{}, nil
}
