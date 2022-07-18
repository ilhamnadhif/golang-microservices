package repository

import (
	"context"
	"gorm.io/gorm"
	"user/model/schema"
)

func NewUserRepositoryImpl() UserRepository {
	return &userRepositoryImpl{}
}

type userRepositoryImpl struct{} // store in database

func (repository *userRepositoryImpl) FindOneByID(ctx context.Context, db *gorm.DB, userID int) (schema.User, error) {
	//TODO implement me
	panic("implement me")
}

func (repository *userRepositoryImpl) FindOneByEmail(ctx context.Context, db *gorm.DB, email string) (schema.User, error) {
	//TODO implement me
	panic("implement me")
}

func (repository *userRepositoryImpl) FindAll(ctx context.Context, db *gorm.DB) ([]schema.User, error) {
	//TODO implement me
	panic("implement me")
}

func (repository *userRepositoryImpl) Create(ctx context.Context, db *gorm.DB, request schema.User) (schema.User, error) {
	//TODO implement me
	panic("implement me")
}

func (repository *userRepositoryImpl) Update(ctx context.Context, db *gorm.DB, request schema.User) (schema.User, error) {
	//TODO implement me
	panic("implement me")
}

func (repository *userRepositoryImpl) Delete(ctx context.Context, db *gorm.DB, userID int) error {
	//TODO implement me
	panic("implement me")
}
