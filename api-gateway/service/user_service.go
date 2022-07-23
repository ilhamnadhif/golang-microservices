package service

import (
	"api-gateway/model/dto"
	"context"
)

type UserService interface {
	FindOneByID(ctx context.Context, userID int) (dto.UserResponse, error)
	FindOneByEmail(ctx context.Context, email string) (dto.UserResponse, error)
	FindAll(ctx context.Context) ([]dto.UserResponse, error)
	Create(ctx context.Context, req dto.UserCreateReq) (dto.UserResponse, error)
	Update(ctx context.Context, req dto.UserUpdateReq) (dto.UserResponse, error)
	Delete(ctx context.Context, userID int) error
}
