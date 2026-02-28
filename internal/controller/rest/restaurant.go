package rest

import (
	"errors"
	"net/http"
	"workshop-restful-api-backend/internal/model"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (r *V1) GetRestaurant(c *gin.Context) {
	ctx := c.Request.Context()

	restaurants, err := r.usecase.RestaurantUsecase.GetRestaurants(ctx)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, restaurants)
}

func (r *V1) CreateRestaurant(c *gin.Context) {
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

func (r *V1) DeleteRestaurant(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	ctx := c.Request.Context()

	err = r.usecase.RestaurantUsecase.DeleteRestaurants(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, nil)
}
