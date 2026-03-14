package usecase

import (
	"workshop-restful-api-backend/internal/repository"
	"workshop-restful-api-backend/pkg/bcrypt"
	"workshop-restful-api-backend/pkg/jwt"

	"golang.org/x/oauth2"
)

type Usecase struct {
	AuthUsecase       IAuthUsecase
	RestaurantUsecase IRestaurantUsecase
	ItemUsecase       IItemUsecase
}

func NewUsecase(jwt jwt.JWT, bcrypt bcrypt.IBcrypt, oauth *oauth2.Config, repository *repository.Repository) *Usecase {
	return &Usecase{
		AuthUsecase:       NewAuthUsecase(jwt, bcrypt, oauth, repository.UserRepository),
		RestaurantUsecase: NewRestaurantUsecase(repository.RestaurantRepository),
		ItemUsecase:       NewItemUsecase(repository.ItemRepository),
	}
}
