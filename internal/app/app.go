package app

import (
	"log"
	"os"
	"workshop-restful-api-backend/internal/controller/rest"
	"workshop-restful-api-backend/internal/repository"
	"workshop-restful-api-backend/internal/usecase"
	"workshop-restful-api-backend/pkg/bcrypt"
	httpserver "workshop-restful-api-backend/pkg/gin"
	"workshop-restful-api-backend/pkg/jwt"
	"workshop-restful-api-backend/pkg/middleware"
	"workshop-restful-api-backend/pkg/postgres"
)

func Run() {
	db := postgres.StartPostgres()
	app := httpserver.Start()
	jwtInit := *jwt.NewJWT()
	bcryptInit := bcrypt.NewBcrypt()
	middleware := middleware.NewMiddleware(&jwtInit)

	repo := repository.NewRepository(db)
	uc := usecase.NewUsecase(jwtInit, bcryptInit, repo)
	v1 := rest.NewV1(middleware, uc)

	rest.NewRouter(app, v1)

	if err := app.Run(":" + os.Getenv("APP_PORT")); err != nil {
		log.Fatalf("Failed to start server: %s", err.Error())
	}
}
