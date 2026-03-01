package repository

import "gorm.io/gorm"

type Repository struct {
	RestaurantRepository IRestaurantRepository
	ItemRepository       IItemRepository
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		RestaurantRepository: NewRestaurantRepository(db),
		ItemRepository:       NewItemRepository(db),
	}
}
