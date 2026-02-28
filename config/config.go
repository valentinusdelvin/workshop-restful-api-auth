package config

import (
	"log"

	"github.com/joho/godotenv"
)

func NewConfig() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
