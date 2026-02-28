package main

import (
	"workshop-restful-api-backend/config"
	"workshop-restful-api-backend/internal/app"
)

func main() {
	config.NewConfig()

	app.Run()
}
