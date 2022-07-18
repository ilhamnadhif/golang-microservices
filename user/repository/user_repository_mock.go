package repository

import (
	"context"
	"gorm.io/gorm"
	"time"
	"user/model/schema"
)

func NewUserRepositoryMock() UserRepository {
	return &userRepositoryMock{}
}

type userRepositoryMock struct{} // mock database

func (repository *userRepositoryMock) FindOneByID(ctx context.Context, db *gorm.DB, userID int) (schema.User, error) {
	return schema.User{
		ID:        1,
		Name:      "ilham",
		Email:     "ilham@gmail.com",
		Password:  "12345",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

func (repository *userRepositoryMock) FindOneByEmail(ctx context.Context, db *gorm.DB, email string) (schema.User, error) {
	return schema.User{
		ID:        1,
		Name:      "ilham",
		Email:     "ilham@gmail.com",
		Password:  "12345",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

func (repository *userRepositoryMock) FindAll(ctx context.Context, db *gorm.DB) ([]schema.User, error) {
	return []schema.User{
		{
			ID:        1,
			Name:      "ilham",
			Email:     "ilham@gmail.com",
			Password:  "12345",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        2,
			Name:      "ivana",
			Email:     "ivana@gmail.com",
			Password:  "12345",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        3,
			Name:      "zuhail",
			Email:     "zuhail@gmail.com",
			Password:  "12345",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}, nil
}

func (repository *userRepositoryMock) Create(ctx context.Context, db *gorm.DB, request schema.User) (schema.User, error) {
	return schema.User{
		ID:        1,
		Name:      "ilham",
		Email:     "ilham@gmail.com",
		Password:  "12345",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

func (repository *userRepositoryMock) Update(ctx context.Context, db *gorm.DB, request schema.User) (schema.User, error) {
	return schema.User{
		ID:        1,
		Name:      "ilham",
		Email:     "ilham@gmail.com",
		Password:  "12345",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

func (repository *userRepositoryMock) Delete(ctx context.Context, db *gorm.DB, userID int) error {
	return nil
}
