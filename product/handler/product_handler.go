package handler

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	"product/model/proto"
	"product/service"
)

func NewProductHandler(productService service.ProductService) ProductServer {
	return ProductServer{
		ProductService: productService,
	}
}

type ProductServer struct {
	proto.UnimplementedProductServiceServer
	ProductService service.ProductService
}

func (server *ProductServer) FindOneByID(ctx context.Context, request *proto.ProductID) (*proto.Product, error) {
	return server.ProductService.FindOneByID(ctx, request)
}
func (server *ProductServer) FindAll(empty *emptypb.Empty, stream proto.ProductService_FindAllServer) error {
	products, err := server.ProductService.FindAll(context.Background())
	if err != nil {
		return err
	}
	for _, product := range products {
		stream.Send(&product)
	}
	return nil
}
func (server *ProductServer) Create(ctx context.Context, request *proto.ProductCreateReq) (*proto.Product, error) {
	return server.ProductService.Create(ctx, request)
}
func (server *ProductServer) Update(ctx context.Context, request *proto.ProductUpdateReq) (*proto.Product, error) {
	return server.ProductService.Update(ctx, request)
}
func (server *ProductServer) Delete(ctx context.Context, request *proto.ProductID) (*emptypb.Empty, error) {
	err := server.ProductService.Delete(ctx, request)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
