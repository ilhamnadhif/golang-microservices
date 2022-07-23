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
	var user schema.User
	err := db.WithContext(ctx).First(&user, "id = ?", userID).Error
	if err != nil {
		return schema.User{}, err
	}
	return user, nil
}

func (repository *userRepositoryImpl) FindOneByEmail(ctx context.Context, db *gorm.DB, email string) (schema.User, error) {
	var user schema.User
	err := db.WithContext(ctx).First(&user, "email = ?", email).Error
	if err != nil {
		return schema.User{}, err
	}
	return user, nil
}

func (repository *userRepositoryImpl) FindAll(ctx context.Context, db *gorm.DB) ([]schema.User, error) {
	var users []schema.User
	err := db.WithContext(ctx).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, err
}

func (repository *userRepositoryImpl) Create(ctx context.Context, db *gorm.DB, request schema.User) (schema.User, error) {
	err := db.WithContext(ctx).Create(&request).Error
	if err != nil {
		return schema.User{}, err
	}
	return request, nil
}

func (repository *userRepositoryImpl) Update(ctx context.Context, db *gorm.DB, request schema.User) (schema.User, error) {
	err := db.WithContext(ctx).Where(&schema.User{ID: request.ID}).Updates(&request).Error
	if err != nil {
		return schema.User{}, err
	}
	return request, nil
}

func (repository *userRepositoryImpl) Delete(ctx context.Context, db *gorm.DB, userID int) error {
	err := db.WithContext(ctx).Where(&schema.User{ID: userID}).Delete(&schema.User{}).Error
	if err != nil {
		return err
	}
	return nil
}
