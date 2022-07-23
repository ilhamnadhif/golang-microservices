package service

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
	"time"
	"user/model/proto"
	"user/model/schema"
	"user/repository"
)

func NewUserServiceImpl(db *gorm.DB, userRepository repository.UserRepository) UserService {
	return &userServiceImpl{
		DB:             db,
		UserRepository: userRepository,
	}
}

type userServiceImpl struct {
	DB             *gorm.DB
	UserRepository repository.UserRepository
}

func (service *userServiceImpl) FindOneByID(ctx context.Context, userID *proto.UserID) (*proto.User, error) {
	tx := service.DB.Begin()
	user, err := service.UserRepository.FindOneByID(ctx, tx, int(userID.ID))
	if err != nil {
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	tx.Commit()
	return &proto.User{
		ID:        int64(user.ID),
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: timestamppb.New(user.UpdatedAt),
	}, nil
}

func (service *userServiceImpl) FindOneByEmail(ctx context.Context, email *proto.UserEmail) (*proto.User, error) {
	tx := service.DB.Begin()
	user, err := service.UserRepository.FindOneByEmail(ctx, tx, email.Email)
	if err != nil {
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	tx.Commit()
	return &proto.User{
		ID:        int64(user.ID),
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: timestamppb.New(user.UpdatedAt),
	}, nil
}

func (service *userServiceImpl) FindAll(ctx context.Context) ([]proto.User, error) {
	tx := service.DB.Begin()
	users, err := service.UserRepository.FindAll(ctx, tx)
	if err != nil {
		tx.Rollback()
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	tx.Commit()
	var usersResponse = make([]proto.User, 0)
	for _, user := range users {
		usersResponse = append(usersResponse, proto.User{
			ID:        int64(user.ID),
			Name:      user.Name,
			Email:     user.Email,
			Password:  user.Password,
			CreatedAt: timestamppb.New(user.CreatedAt),
			UpdatedAt: timestamppb.New(user.UpdatedAt),
		})
	}
	return usersResponse, nil
}

func (service *userServiceImpl) Create(ctx context.Context, request *proto.UserCreateReq) (*proto.User, error) {
	tx := service.DB.Begin()
	user, err := service.UserRepository.Create(ctx, tx, schema.User{
		Name:      request.Name,
		Email:     request.Email,
		Password:  request.Password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		tx.Rollback()
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	tx.Commit()
	return &proto.User{
		ID:        int64(user.ID),
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: timestamppb.New(user.UpdatedAt),
	}, nil
}

func (service *userServiceImpl) Update(ctx context.Context, request *proto.UserUpdateReq) (*proto.User, error) {
	fmt.Println(request)
	tx := service.DB.Begin()
	findUser, err := service.UserRepository.FindOneByID(ctx, tx, int(request.ID))
	if err != nil {
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	user, err := service.UserRepository.Update(ctx, tx, schema.User{
		ID:        int(request.ID),
		Name:      request.Name,
		Email:     request.Email,
		Password:  request.Password,
		CreatedAt: findUser.CreatedAt,
		UpdatedAt: time.Now(),
	})
	if err != nil {
		tx.Rollback()
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	tx.Commit()
	return &proto.User{
		ID:        int64(user.ID),
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: timestamppb.New(user.UpdatedAt),
	}, nil
}

func (service *userServiceImpl) Delete(ctx context.Context, userID *proto.UserID) error {
	tx := service.DB.Begin()
	_, err := service.UserRepository.FindOneByID(ctx, tx, int(userID.ID))
	if err != nil {
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return status.Error(codes.NotFound, err.Error())
		}
		return status.Error(codes.InvalidArgument, err.Error())
	}
	err = service.UserRepository.Delete(ctx, tx, int(userID.ID))
	if err != nil {
		tx.Rollback()
		return status.Error(codes.InvalidArgument, err.Error())
	}
	tx.Commit()
	return nil
}
