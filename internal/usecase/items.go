package usecase

import (
	"context"
	"workshop-restful-api-backend/internal/entity"
	"workshop-restful-api-backend/internal/model"
	"workshop-restful-api-backend/internal/repository"

	"github.com/google/uuid"
)

type IItemUsecase interface {
	GetRestaurantItems(ctx context.Context, restaurantId uuid.UUID) ([]model.ItemResponse, error)
	CreateItem(ctx context.Context, creteItem model.CreateItem) (*model.ItemResponse, error)
}

type ItemUsecase struct {
	itemRepository       repository.IItemRepository
	restaurantRepository repository.IRestaurantRepository
}

func NewItemUsecase(itemRepository repository.IItemRepository, restaurantRepository repository.IRestaurantRepository) *ItemUsecase {
	return &ItemUsecase{
		itemRepository:       itemRepository,
		restaurantRepository: restaurantRepository,
	}
}

func (u *ItemUsecase) GetRestaurantItems(ctx context.Context, restaurantId uuid.UUID) ([]model.ItemResponse, error) {
	items, err := u.itemRepository.GetRestaurantItems(ctx, restaurantId)
	if err != nil {
		return nil, err
	}

	responses := model.ToItemResponses(items)
	return responses, nil
}

func (u *ItemUsecase) CreateItem(ctx context.Context, creteItem model.CreateItem) (*model.ItemResponse, error) {
	item := entity.Item{
		Id:           uuid.New(),
		RestaurantId: creteItem.RestaurantId,
		Name:         creteItem.Name,
		Price:        creteItem.Price,
		Available:    creteItem.Available,
	}

	err := u.itemRepository.CreateItem(ctx, item)
	if err != nil {
		return nil, err
	}

	response := model.ToItemResponse(item)
	return &response, nil
}
