package service

import (
	"api-gateway/model"
	"api-gateway/model/dto"
	"api-gateway/model/proto"
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
)

type userServiceImpl struct {
	UserService proto.UserServiceClient
}

func NewUserService(userService proto.UserServiceClient) UserService {
	return &userServiceImpl{
		UserService: userService,
	}
}

func (service *userServiceImpl) FindOneByID(ctx context.Context, userID int) (dto.UserResponse, error) {
	user, err := service.UserService.FindOneByID(ctx, &proto.UserID{
		ID: int64(userID),
	})
	if err != nil {
		return dto.UserResponse{}, echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	return dto.UserResponse{
		ID:        int(user.ID),
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: model.ToTime(user.CreatedAt.AsTime()),
		UpdatedAt: model.ToTime(user.UpdatedAt.AsTime()),
	}, nil
}

func (service *userServiceImpl) FindOneByEmail(ctx context.Context, email string) (dto.UserResponse, error) {
	user, err := service.UserService.FindOneByEmail(ctx, &proto.UserEmail{
		Email: email,
	})
	if err != nil {
		return dto.UserResponse{}, echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	return dto.UserResponse{
		ID:        int(user.ID),
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: model.ToTime(user.CreatedAt.AsTime()),
		UpdatedAt: model.ToTime(user.UpdatedAt.AsTime()),
	}, nil
}

func (service *userServiceImpl) FindAll(ctx context.Context) ([]dto.UserResponse, error) {
	stream, err := service.UserService.FindAll(ctx, &empty.Empty{})
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	var usersResponse = make([]dto.UserResponse, 0)
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		usersResponse = append(usersResponse, dto.UserResponse{
			ID:        int(msg.ID),
			Name:      msg.Name,
			Email:     msg.Email,
			Password:  msg.Password,
			CreatedAt: model.ToTime(msg.CreatedAt.AsTime()),
			UpdatedAt: model.ToTime(msg.UpdatedAt.AsTime()),
		})
	}
	return usersResponse, nil
}

func (service *userServiceImpl) Create(ctx context.Context, req dto.UserCreateReq) (dto.UserResponse, error) {
	user, err := service.UserService.Create(ctx, &proto.UserCreateReq{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return dto.UserResponse{}, echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return dto.UserResponse{
		ID:        int(user.ID),
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: model.ToTime(user.CreatedAt.AsTime()),
		UpdatedAt: model.ToTime(user.UpdatedAt.AsTime()),
	}, nil
}

func (service *userServiceImpl) Update(ctx context.Context, req dto.UserUpdateReq) (dto.UserResponse, error) {
	user, err := service.UserService.Update(ctx, &proto.UserUpdateReq{
		ID:       int64(req.ID),
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return dto.UserResponse{}, echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return dto.UserResponse{
		ID:        int(user.ID),
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: model.ToTime(user.CreatedAt.AsTime()),
		UpdatedAt: model.ToTime(user.UpdatedAt.AsTime()),
	}, nil
}

func (service *userServiceImpl) Delete(ctx context.Context, userID int) error {
	_, err := service.UserService.Delete(ctx, &proto.UserID{
		ID: int64(userID),
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}
