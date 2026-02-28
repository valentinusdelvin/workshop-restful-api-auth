package usecase

import (
	"time"
	"workshop-restful-api-backend/internal/entity"
	"workshop-restful-api-backend/internal/model"
	"workshop-restful-api-backend/internal/repository"

	"github.com/google/uuid"
)

type IRestaurantUsecase interface {
	CreateRestaurant(createRestaurant model.CreateRestaurant) (*model.RestaurantResponse, error)
	GetRestaurants() ([]entity.Restaurant, error)
}

type RestaurantUsecase struct {
	restaurantRepository repository.IRestaurantRepository
}

func NewRestaurantUsecase(restaurantRepository repository.IRestaurantRepository) *RestaurantUsecase {
	return &RestaurantUsecase{restaurantRepository}
}

func (r *RestaurantUsecase) CreateRestaurant(createRestaurant model.CreateRestaurant) (*model.RestaurantResponse, error) {
	restaurant := entity.Restaurant{
		Id:        uuid.New(),
		Name:      createRestaurant.Name,
		Location:  createRestaurant.Location,
		CreatedAt: time.Now(),
	}

	err := r.restaurantRepository.CreateRestaurant(restaurant)
	if err != nil {
		return nil, err
	}

	response := model.ToRestourantResponse(restaurant)

	return &response, nil
}

func (r *RestaurantUsecase) GetRestaurants() ([]entity.Restaurant, error) {
	return r.restaurantRepository.GetRestaurants()
}
