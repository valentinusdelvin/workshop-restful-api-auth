package app

import (
	"workshop-restful-api-backend/internal/controller/rest"
	"workshop-restful-api-backend/pkg/postgres"

	"github.com/gin-gonic/gin"
)

func Run() {
	_ = postgres.StartPostgres()

	app := gin.New()

	rest.NewRouter(app)
}
