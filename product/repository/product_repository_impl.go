package repository

import (
	"context"
	"gorm.io/gorm"
	"product/model/schema"
)

func NewProductRepositoryImpl() ProductRepository {
	return &productRepositoryImpl{}
}

type productRepositoryImpl struct{} // store in database

func (repository *productRepositoryImpl) FindOneByID(ctx context.Context, db *gorm.DB, productID int) (schema.Product, error) {
	var product schema.Product
	err := db.WithContext(ctx).First(&product, "id = ?", productID).Error
	if err != nil {
		return schema.Product{}, err
	}
	return product, nil
}

func (repository *productRepositoryImpl) FindAll(ctx context.Context, db *gorm.DB) ([]schema.Product, error) {
	var products []schema.Product
	err := db.WithContext(ctx).Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, err
}

func (repository *productRepositoryImpl) Create(ctx context.Context, db *gorm.DB, request schema.Product) (schema.Product, error) {
	err := db.WithContext(ctx).Create(&request).Error
	if err != nil {
		return schema.Product{}, err
	}
	return request, nil
}

func (repository *productRepositoryImpl) Update(ctx context.Context, db *gorm.DB, request schema.Product) (schema.Product, error) {
	err := db.WithContext(ctx).Where(&schema.Product{ID: request.ID}).Updates(&request).Error
	if err != nil {
		return schema.Product{}, err
	}
	return request, nil
}

func (repository *productRepositoryImpl) Delete(ctx context.Context, db *gorm.DB, productID int) error {
	err := db.WithContext(ctx).Where(&schema.Product{ID: productID}).Delete(&schema.Product{}).Error
	if err != nil {
		return err
	}
	return nil
}
