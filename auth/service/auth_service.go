package service

import (
	"auth/model/proto"
	"context"
)

type AuthService interface {
	Login(ctx context.Context, request *proto.LoginReq) (*proto.LoginResponse, error)
}
