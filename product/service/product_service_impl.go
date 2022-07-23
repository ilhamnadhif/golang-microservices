package service

import (
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
	"product/model/proto"
	"product/model/schema"
	"product/repository"
	"time"
)

func NewProductServiceImpl(db *gorm.DB, productRepository repository.ProductRepository) ProductService {
	return &productServiceImpl{
		DB:                db,
		ProductRepository: productRepository,
	}
}

type productServiceImpl struct {
	DB                *gorm.DB
	ProductRepository repository.ProductRepository
}

func (service *productServiceImpl) FindOneByID(ctx context.Context, productID *proto.ProductID) (*proto.Product, error) {
	tx := service.DB.Begin()
	product, err := service.ProductRepository.FindOneByID(ctx, tx, int(productID.ID))
	if err != nil {
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	tx.Commit()
	return &proto.Product{
		ID:          int64(product.ID),
		Name:        product.Name,
		Description: product.Description,
		Price:       int64(product.Price),
		CreatedAt:   timestamppb.New(product.CreatedAt),
		UpdatedAt:   timestamppb.New(product.UpdatedAt),
	}, nil
}

func (service *productServiceImpl) FindAll(ctx context.Context) ([]proto.Product, error) {
	tx := service.DB.Begin()
	products, err := service.ProductRepository.FindAll(ctx, tx)
	if err != nil {
		tx.Rollback()
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	tx.Commit()
	var productsResponse = make([]proto.Product, 0)
	for _, product := range products {
		productsResponse = append(productsResponse, proto.Product{
			ID:          int64(product.ID),
			Name:        product.Name,
			Description: product.Description,
			Price:       int64(product.Price),
			CreatedAt:   timestamppb.New(product.CreatedAt),
			UpdatedAt:   timestamppb.New(product.UpdatedAt),
		})
	}
	return productsResponse, nil
}

func (service *productServiceImpl) Create(ctx context.Context, request *proto.ProductCreateReq) (*proto.Product, error) {
	tx := service.DB.Begin()
	product, err := service.ProductRepository.Create(ctx, tx, schema.Product{
		Name:        request.Name,
		Description: request.Description,
		Price:       int(request.Price),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	})
	if err != nil {
		tx.Rollback()
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	tx.Commit()
	return &proto.Product{
		ID:          int64(product.ID),
		Name:        product.Name,
		Description: product.Description,
		Price:       int64(product.Price),
		CreatedAt:   timestamppb.New(product.CreatedAt),
		UpdatedAt:   timestamppb.New(product.UpdatedAt),
	}, nil
}

func (service *productServiceImpl) Update(ctx context.Context, request *proto.ProductUpdateReq) (*proto.Product, error) {
	tx := service.DB.Begin()
	findProduct, err := service.ProductRepository.FindOneByID(ctx, tx, int(request.ID))
	if err != nil {
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	product, err := service.ProductRepository.Create(ctx, tx, schema.Product{
		ID:          int(request.ID),
		Name:        request.Name,
		Description: request.Description,
		Price:       int(request.ID),
		CreatedAt:   findProduct.CreatedAt,
		UpdatedAt:   time.Now(),
	})
	if err != nil {
		tx.Rollback()
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	tx.Commit()
	return &proto.Product{
		ID:          int64(product.ID),
		Name:        product.Name,
		Description: product.Description,
		Price:       int64(product.Price),
		CreatedAt:   timestamppb.New(product.CreatedAt),
		UpdatedAt:   timestamppb.New(product.UpdatedAt),
	}, nil
}

func (service *productServiceImpl) Delete(ctx context.Context, productID *proto.ProductID) error {
	tx := service.DB.Begin()
	_, err := service.ProductRepository.FindOneByID(ctx, tx, int(productID.ID))
	if err != nil {
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return status.Error(codes.NotFound, err.Error())
		}
		return status.Error(codes.InvalidArgument, err.Error())
	}
	err = service.ProductRepository.Delete(ctx, tx, int(productID.ID))
	if err != nil {
		tx.Rollback()
		return status.Error(codes.InvalidArgument, err.Error())
	}
	tx.Commit()
	return nil
}
