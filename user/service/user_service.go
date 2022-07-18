package service

import (
	"context"
	"user/model/proto"
)

type UserService interface {
	FindOneByID(ctx context.Context, userID *proto.UserID) (*proto.User, error)
	FindOneByEmail(ctx context.Context, email *proto.UserEmail) (*proto.User, error)
	FindAll(ctx context.Context) ([]proto.User, error)
	Create(ctx context.Context, request *proto.UserCreateReq) (*proto.User, error)
	Update(ctx context.Context, request *proto.UserUpdateReq) (*proto.User, error)
	Delete(ctx context.Context, userID *proto.UserID) error
}
