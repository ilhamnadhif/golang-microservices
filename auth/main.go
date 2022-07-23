package main

import (
	"auth/app"
	"auth/config"
	"auth/handler"
	"auth/model/proto"
	"auth/service"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	initConfig := config.InitConfig()
	userService, connClose := app.InitUserService(initConfig.Service[config.User])
	defer connClose()
	authService := service.NewAuthService(userService)
	authHandler := handler.NewAuthHandler(authService)

	lis, err := net.Listen("tcp", initConfig.Server.HostPort)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("server run on port %v", initConfig.Server.HostPort)

	s := grpc.NewServer()
	proto.RegisterAuthServiceServer(s, &authHandler)

	if err := s.Serve(lis); err != nil {
		log.Fatalln(err)
	}
}
