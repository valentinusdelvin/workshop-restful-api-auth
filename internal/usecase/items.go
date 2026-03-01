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
	DeleteItem(ctx context.Context, id uuid.UUID) error
	EditItem(ctx context.Context, id uuid.UUID, edit model.EditItem) error
}

type ItemUsecase struct {
	itemRepository repository.IItemRepository
}

func NewItemUsecase(itemRepository repository.IItemRepository) *ItemUsecase {
	return &ItemUsecase{
		itemRepository: itemRepository,
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

func (u *ItemUsecase) DeleteItem(ctx context.Context, id uuid.UUID) error {
	return u.itemRepository.DeleteItem(ctx, id)
}

func (u *ItemUsecase) EditItem(ctx context.Context, id uuid.UUID, edit model.EditItem) error {
	return u.itemRepository.EditItem(ctx, id, edit)
}
