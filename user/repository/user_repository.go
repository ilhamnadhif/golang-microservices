package repository

import (
	"context"
	"gorm.io/gorm"
	"user/model/schema"
)

type UserRepository interface {
	FindOneByID(ctx context.Context, db *gorm.DB, userID int) (schema.User, error)
	FindOneByEmail(ctx context.Context, db *gorm.DB, email string) (schema.User, error)
	FindAll(ctx context.Context, db *gorm.DB) ([]schema.User, error)
	Create(ctx context.Context, db *gorm.DB, request schema.User) (schema.User, error)
	Update(ctx context.Context, db *gorm.DB, request schema.User) (schema.User, error)
	Delete(ctx context.Context, db *gorm.DB, userID int) error
}
