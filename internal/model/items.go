package model

import (
	"workshop-restful-api-backend/internal/entity"

	"github.com/google/uuid"
)

type ItemResponse struct {
	Id           uuid.UUID `json:"id"`
	RestaurantId uuid.UUID `json:"restaurant_id"`
	Name         string    `json:"name"`
	Price        float64   `json:"price"`
	Available    bool      `json:"available"`
}

func ToItemResponse(item entity.Item) ItemResponse {
	return ItemResponse{
		Id:           item.Id,
		RestaurantId: item.RestaurantId,
		Name:         item.Name,
		Price:        item.Price,
		Available:    item.Available,
	}
}

func ToItemResponses(items []entity.Item) []ItemResponse {
	var responses []ItemResponse
	for _, item := range items {
		responses = append(responses, ToItemResponse(item))
	}

	return responses
}

type CreateItem struct {
	Name         string    `json:"name"`
	RestaurantId uuid.UUID `json:"restaurant_id"`
	Price        float64   `json:"price"`
	Available    bool      `json:"available"`
}
