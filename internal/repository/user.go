package repository

import (
	"context"
	"workshop-restful-api-backend/internal/entity"

	"gorm.io/gorm"
)

type IUserRepository interface {
	CreateUser(ctx context.Context, user entity.User) error
	GetUserByUsername(ctx context.Context, username string) (*entity.User, error)
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{db: db}
}

func (u *UserRepository) CreateUser(ctx context.Context, user entity.User) error {
	err := gorm.G[entity.User](u.db).Create(ctx, &user)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserRepository) GetUserByUsername(ctx context.Context, username string) (*entity.User, error) {
	var user entity.User
	user, err := gorm.G[entity.User](u.db).
		Where("username = ?", username).
		First(ctx)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
