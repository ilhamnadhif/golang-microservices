package handler

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/protobuf/types/known/emptypb"
	"user/model/proto"
	"user/service"
)

func NewUserHandler(userService service.UserService) UserServer {
	return UserServer{
		UserService: userService,
	}
}

type UserServer struct {
	proto.UnimplementedUserServiceServer
	UserService service.UserService
}

func (server *UserServer) FindOneByID(ctx context.Context, request *proto.UserID) (*proto.User, error) {
	return server.UserService.FindOneByID(ctx, request)
}
func (server *UserServer) FindOneByEmail(ctx context.Context, request *proto.UserEmail) (*proto.User, error) {
	return server.UserService.FindOneByEmail(ctx, request)
}
func (server *UserServer) FindAll(empty *emptypb.Empty, stream proto.UserService_FindAllServer) error {
	users, err := server.UserService.FindAll(context.Background())
	if err != nil {
		return err
	}
	for _, user := range users {
		stream.Send(&user)
	}
	return nil
}
func (server *UserServer) Create(ctx context.Context, request *proto.UserCreateReq) (*proto.User, error) {
	return server.UserService.Create(ctx, request)
}
func (server *UserServer) Update(ctx context.Context, request *proto.UserUpdateReq) (*proto.User, error) {
	return server.UserService.Update(ctx, request)
}
func (server *UserServer) Delete(ctx context.Context, request *proto.UserID) (*empty.Empty, error) {
	err := server.UserService.Delete(ctx, request)
	if err != nil {
		return &empty.Empty{}, err
	}
	return &empty.Empty{}, nil
}
