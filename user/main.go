package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"user/app"
	"user/config"
	"user/handler"
	"user/model/proto"
	"user/repository"
	"user/service"
)

func main() {
	initConfig := config.InitConfig()
	db := app.InitGorm(initConfig.Database)
	userRepository := repository.NewUserRepositoryMock()
	userService := service.NewUserServiceImpl(db, userRepository)
	userHandler := handler.NewUserHandler(userService)

	lis, err := net.Listen("tcp", initConfig.Server.HostPort)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("server run on port %v", initConfig.Server.HostPort)

	s := grpc.NewServer()
	proto.RegisterUserServiceServer(s, &userHandler)

	if err := s.Serve(lis); err != nil {
		log.Fatalln(err)
	}
}
