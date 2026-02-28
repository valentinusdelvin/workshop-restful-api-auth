package repository

import (
	"context"
	"workshop-restful-api-backend/internal/entity"

	"gorm.io/gorm"
)

type IRestaurantRepository interface {
	CreateRestaurant(ctx context.Context, trestaurant entity.Restaurant) error
	GetRestaurants(ctx context.Context) ([]entity.Restaurant, error)
}

type RestaurantRepository struct {
	db *gorm.DB
}

func NewRestaurantRepository(db *gorm.DB) *RestaurantRepository {
	return &RestaurantRepository{db}
}

func (r *RestaurantRepository) CreateRestaurant(ctx context.Context, restaurant entity.Restaurant) error {
	err := gorm.G[entity.Restaurant](r.db).Create(ctx, &restaurant)
	if err != nil {
		return err
	}

	return nil
}

func (r *RestaurantRepository) GetRestaurants(ctx context.Context) ([]entity.Restaurant, error) {
	restaurants, err := gorm.G[entity.Restaurant](r.db).Find(ctx)
	if err != nil {
		return nil, err
	}

	return restaurants, nil
}
