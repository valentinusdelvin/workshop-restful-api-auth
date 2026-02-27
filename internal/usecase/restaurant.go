package usecase

import (
	"workshop-restful-api-backend/internal/entity"
	"workshop-restful-api-backend/internal/repository"
)

type IRestaurantUsecase interface {
	CreateRestaurant(restaurant entity.Restaurant) error
	GetRestaurants() ([]entity.Restaurant, error)
}

type RestaurantUsecase struct {
	restaurantRepository repository.IRestaurantRepository
}

func NewRestaurantUsecase(restaurantRepository repository.IRestaurantRepository) *RestaurantUsecase {
	return &RestaurantUsecase{restaurantRepository}
}

func (r *RestaurantUsecase) CreateRestaurant(restaurant entity.Restaurant) error {
	return r.restaurantRepository.CreateRestaurant(restaurant)
}

func (r *RestaurantUsecase) GetRestaurants() ([]entity.Restaurant, error) {
	return r.restaurantRepository.GetRestaurants()
}
