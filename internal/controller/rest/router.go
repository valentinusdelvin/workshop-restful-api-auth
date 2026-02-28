package rest

import (
	"github.com/gin-gonic/gin"
)

func NewRouter(app *gin.Engine, v1 *V1) {
	api := app.Group("/api/v1")
	{
		restaurants := api.Group("/restaurants")
		{
			restaurants.GET("/", v1.GetRestaurant)
			restaurants.POST("/", v1.CreateRestaurant)
		}
	}

	app.Run()
}
