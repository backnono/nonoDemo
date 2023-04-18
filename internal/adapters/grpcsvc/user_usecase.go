package grpcsvc

import (
	"context"
	"nonoDemo/proto_gen/api/hello"
)

type UserUsecase struct {
	hello.UnimplementedUserServiceServer
}

func NewUserUsecase() *UserUsecase {
	return &UserUsecase{}
}

func (u *UserUsecase) Login(ctx context.Context, request *hello.LoginRequest) (*hello.LoginResponse, error) {
	return &hello.LoginResponse{Token: request.Username + request.Password}, nil
}

func (u *UserUsecase) CreateUser(ctx context.Context, request *hello.CreateUserRequest) (*hello.CreateUserResponse, error) {
	// TODO implement me
	panic("implement me")
}

func (u *UserUsecase) GetUser(ctx context.Context, request *hello.GetUserRequest) (*hello.GetUserResponse, error) {
	// TODO implement me
	panic("implement me")
}
