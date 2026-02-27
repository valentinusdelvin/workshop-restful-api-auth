package app

import (
	"workshop-restful-api-backend/internal/controller/rest"
	"workshop-restful-api-backend/internal/repository"
	"workshop-restful-api-backend/pkg/postgres"

	"github.com/gin-gonic/gin"
)

func Run() {
	db := postgres.StartPostgres()

	_ = repository.NewRepository(db)

	app := gin.New()

	rest.NewRouter(app)
}
