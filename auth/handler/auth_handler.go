package handler

import (
	"auth/model/proto"
	"auth/service"
	"context"
)

func NewAuthHandler(authService service.AuthService) AuthServer {
	return AuthServer{
		AuthService: authService,
	}
}

type AuthServer struct {
	proto.UnimplementedAuthServiceServer
	AuthService service.AuthService
}

func (server *AuthServer) Login(ctx context.Context, request *proto.LoginReq) (*proto.LoginResponse, error) {
	return server.AuthService.Login(ctx, request)
}
