package service

import (
	"auth/config"
	"auth/model/dto"
	"auth/model/proto"
	"context"
	"github.com/golang-jwt/jwt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

func NewAuthService(userService proto.UserServiceClient) AuthService {
	return &authService{
		UserService: userService,
	}
}

type authService struct {
	UserService proto.UserServiceClient
	Config      config.Config
}

func (service *authService) Login(ctx context.Context, request *proto.LoginReq) (*proto.LoginResponse, error) {
	user, err := service.UserService.FindOneByEmail(ctx, &proto.UserEmail{
		Email: request.Email,
	})
	if err != nil {
		return nil, err
	}
	if request.Password != user.Password {
		return nil, status.Errorf(codes.Unauthenticated, "password not match")
	}
	token, err := service.GenerateToken(dto.JWTCustomClaims{
		ID:    int(user.ID),
		Email: user.Email,
	})
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	return &proto.LoginResponse{
		Token: token,
	}, nil
}

func (service *authService) GenerateToken(claims dto.JWTCustomClaims) (string, error) {
	claims.StandardClaims = jwt.StandardClaims{
		ExpiresAt: time.Now().Add(service.Config.Jwt.ExpiresIn).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(service.Config.Jwt.SigningKey))
	if err != nil {
		return "", err
	}
	return t, nil
}
