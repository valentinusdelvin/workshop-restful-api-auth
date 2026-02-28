package rest

import (
	"net/http"
	"workshop-restful-api-backend/internal/model"

	"github.com/gin-gonic/gin"
)

func (r *Controller) GetRestaurant(c *gin.Context) {
	ctx := c.Request.Context()

	restaurants, err := r.usecase.RestaurantUsecase.GetRestaurants(ctx)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, restaurants)
}

func (r *Controller) CreateRestaurant(c *gin.Context) {
	var createRestaurant model.CreateRestaurant

	err := c.ShouldBindBodyWithJSON(&createRestaurant)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	ctx := c.Request.Context()

	restaurant, err := r.usecase.RestaurantUsecase.CreateRestaurant(ctx, createRestaurant)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, restaurant)
}
