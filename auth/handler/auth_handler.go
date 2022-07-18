package handler

import (
	"auth/model/proto"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func NewAuthHandler(userService proto.UserServiceClient) AuthServer {
	return AuthServer{
		UserService: userService,
	}
}

type AuthServer struct {
	proto.UnimplementedAuthServiceServer
	UserService proto.UserServiceClient
}

func (server *AuthServer) Login(ctx context.Context, request *proto.LoginReq) (*proto.LoginResponse, error) {
	user, err := server.UserService.FindOneByEmail(ctx, &proto.UserEmail{
		Email: request.Email,
	})
	if err != nil {
		return nil, err
	}
	if request.Password != user.Password {
		return nil, status.Errorf(codes.Unauthenticated, "password not match")
	}
	return &proto.LoginResponse{
		Token: "ajshdjkahsdkjashdkjskjads",
	}, nil
}
