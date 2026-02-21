package app

import "workshop-restful-api-backend/pkg/postgres"

func Run() {
	_ = postgres.StartPostgres()

}
