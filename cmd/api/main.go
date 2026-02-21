package main

import (
	"workshop-restful-api-backend/config"
	"workshop-restful-api-backend/internal/app"
)

func main() {
	config.NewCofig()

	app.Run()
}
