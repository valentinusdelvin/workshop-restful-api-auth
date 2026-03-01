package usecase

import "workshop-restful-api-backend/internal/repository"

type Usecase struct {
	RestaurantUsecase IRestaurantUsecase
	ItemUsecase       IItemUsecase
}

func NewUsecase(repository *repository.Repository) *Usecase {
	return &Usecase{
		RestaurantUsecase: NewRestaurantUsecase(repository.RestaurantRepository),
		ItemUsecase:       NewItemUsecase(repository.ItemRepository, repository.RestaurantRepository),
	}
}
