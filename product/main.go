package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"product/app"
	"product/config"
	"product/handler"
	"product/model/proto"
	"product/repository"
	"product/service"
)

func main() {
	initConfig := config.InitConfig()
	db := app.InitGorm(initConfig.Database)
	productRepository := repository.NewProductRepositoryImpl()
	productService := service.NewProductServiceImpl(db, productRepository)
	productHandler := handler.NewProductHandler(productService)

	lis, err := net.Listen("tcp", initConfig.Server.HostPort)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("server run on port %v", initConfig.Server.HostPort)

	s := grpc.NewServer()
	proto.RegisterProductServiceServer(s, &productHandler)

	if err := s.Serve(lis); err != nil {
		log.Fatalln(err)
	}
}
