package rest

import (
	"workshop-restful-api-backend/internal/usecase"
	"workshop-restful-api-backend/pkg/middleware"
)

type V1 struct {
	middleware.IMiddleware
	usecase *usecase.Usecase
}

func NewV1(middleware middleware.IMiddleware, usecase *usecase.Usecase) *V1 {
	return &V1{middleware, usecase}
}
