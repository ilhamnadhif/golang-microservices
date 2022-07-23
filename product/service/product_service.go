package service

import (
	"context"
	"product/model/proto"
)

type ProductService interface {
	FindOneByID(ctx context.Context, userID *proto.ProductID) (*proto.Product, error)
	FindAll(ctx context.Context) ([]proto.Product, error)
	Create(ctx context.Context, request *proto.ProductCreateReq) (*proto.Product, error)
	Update(ctx context.Context, request *proto.ProductUpdateReq) (*proto.Product, error)
	Delete(ctx context.Context, userID *proto.ProductID) error
}
