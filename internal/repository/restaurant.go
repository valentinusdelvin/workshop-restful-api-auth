package repository

import (
	"context"
	"workshop-restful-api-backend/internal/entity"
	"workshop-restful-api-backend/internal/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IRestaurantRepository interface {
	CreateRestaurant(ctx context.Context, restaurant entity.Restaurant) error
	GetRestaurants(ctx context.Context) ([]entity.Restaurant, error)
	DeleteRestaurant(ctx context.Context, id uuid.UUID) error
	EditRestaurant(ctx context.Context, id uuid.UUID, edit model.EditRestaurant) error
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

func (r *RestaurantRepository) DeleteRestaurant(ctx context.Context, id uuid.UUID) error {
	rows, err := gorm.G[entity.Restaurant](r.db).Where("id = ?", id).Delete(ctx)
	if err != nil {
		return err
	}

	if rows == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (r *RestaurantRepository) EditRestaurant(ctx context.Context, id uuid.UUID, edit model.EditRestaurant) error {
	rows, err := gorm.G[entity.Restaurant](r.db).
		Where("id = ?", id).
		Updates(ctx, entity.Restaurant{Name: edit.Name, Location: edit.Location})

	if err != nil {
		return err
	}

	if rows == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
