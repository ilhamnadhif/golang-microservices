package app

import (
	"api-gateway/config"
	"api-gateway/model/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func InitProductService(config config.ServiceConfig) (proto.ProductServiceClient, func() error) {
	conn, err := grpc.Dial(config.HostPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln(err)
	}
	//defer conn.Close()
	client := proto.NewProductServiceClient(conn)
	return client, conn.Close
}
