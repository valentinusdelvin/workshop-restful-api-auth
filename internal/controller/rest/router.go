package rest

import (
	"github.com/gin-gonic/gin"
)

func NewRouter(app *gin.Engine, v1 *V1) {
	api := app.Group("/api/v1")
	{
		restaurants := api.Group("/restaurants")
		{
			restaurants.GET("/", v1.GetRestaurants)
			restaurants.POST("/", v1.CreateRestaurant)
			restaurants.DELETE("/:id", v1.DeleteRestaurant)
			restaurants.PATCH("/:id", v1.EditRestaurant)

			restaurantItems := restaurants.Group("/:id/items")
			{
				restaurantItems.GET("/", v1.GetRestaurantItems)
			}
		}

		items := api.Group("/items")
		{
			items.POST("/", v1.CreateItem)
			items.DELETE("/:id", v1.DeleteItem)
			items.PATCH("/:id", v1.EditItem)
		}
	}
}
