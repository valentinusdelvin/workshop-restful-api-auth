package middleware

import (
	"workshop-restful-api-backend/pkg/jwt"

	"github.com/gin-gonic/gin"
)

type IMiddleware interface {
	Authentication(c *gin.Context)
}

type Middleware struct {
	jwt jwt.IJWT
}

func NewMiddleware(jwt jwt.IJWT) IMiddleware {
	return &Middleware{
		jwt: jwt,
	}
}
