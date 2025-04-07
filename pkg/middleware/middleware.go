package middleware

import (
	"Cinema_API/internal/usecase"
	"Cinema_API/pkg/jwt"
	"github.com/gin-gonic/gin"
)

type MiddlewareInterface interface {
	Authentication(ctx *gin.Context)
}

type Middleware struct {
	usecase usecase.Usecase
	jwt     jwt.JwtService
}

func Init(usecase usecase.Usecase) *Middleware {
	return &Middleware{
		usecase: usecase,
		jwt:     jwt.JWTInit(),
	}
}
