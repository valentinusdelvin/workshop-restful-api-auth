package usecase

import "workshop-restful-api-backend/internal/repository"

type Usecase struct {
	RestaurantUsecase IRestaurantUsecase
}

func NewUsecase(repository *repository.Repository) *Usecase {
	return &Usecase{
		RestaurantUsecase: NewRestaurantUsecase(repository.RestaurantRepository),
	}
}
