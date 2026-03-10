package middleware

import "workshop-restful-api-backend/pkg/jwt"

type IMiddleware interface{}

type Middleware struct {
	jwt jwt.IJWT
}

func NewMiddleware(jwt jwt.IJWT) IMiddleware {
	return &Middleware{
		jwt: jwt,
	}
}
