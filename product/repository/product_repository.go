package repository

import (
	"context"
	"gorm.io/gorm"
	"product/model/schema"
)

type ProductRepository interface {
	FindOneByID(ctx context.Context, db *gorm.DB, userID int) (schema.Product, error)
	FindAll(ctx context.Context, db *gorm.DB) ([]schema.Product, error)
	Create(ctx context.Context, db *gorm.DB, request schema.Product) (schema.Product, error)
	Update(ctx context.Context, db *gorm.DB, request schema.Product) (schema.Product, error)
	Delete(ctx context.Context, db *gorm.DB, userID int) error
}
