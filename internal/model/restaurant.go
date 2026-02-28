package model

import (
	"workshop-restful-api-backend/internal/entity"

	"github.com/google/uuid"
)

type CreateRestaurant struct {
	Name     string `json:"name"`
	Location string `json:"location"`
}

type EditRestaurant struct {
	Name     string `json:"name"`
	Location string `json:"location"`
}

type RestaurantResponse struct {
	Id       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Location string    `json:"location"`
}

func ToRestaurantResponse(restaurant entity.Restaurant) RestaurantResponse {
	return RestaurantResponse{
		Id:       restaurant.Id,
		Name:     restaurant.Name,
		Location: restaurant.Location,
	}
}

func ToRestaurantResponses(restaurants []entity.Restaurant) []RestaurantResponse {
	var responses []RestaurantResponse
	for _, restaurant := range restaurants {
		responses = append(responses, ToRestaurantResponse(restaurant))
	}

	return responses
}
