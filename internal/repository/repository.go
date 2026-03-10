package repository

import "gorm.io/gorm"

type Repository struct {
	UserRepository       IUserRepository
	RestaurantRepository IRestaurantRepository
	ItemRepository       IItemRepository
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		UserRepository:       NewUserRepository(db),
		RestaurantRepository: NewRestaurantRepository(db),
		ItemRepository:       NewItemRepository(db),
	}
}
