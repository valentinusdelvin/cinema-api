package rest

import (
	"Cinema_API/internal/usecase"
	"Cinema_API/pkg/middleware"
	"github.com/gin-gonic/gin"
	"log"
)

type Rest struct {
	Usecase    *usecase.Usecase
	Router     *gin.Engine
	Middleware middleware.MiddlewareInterface
}

func RestInit(usecase *usecase.Usecase, middleware middleware.MiddlewareInterface) *Rest {
	router := gin.Default()

	return &Rest{
		Usecase:    usecase,
		Router:     router,
		Middleware: middleware,
	}
}

func (r *Rest) FinalEndpoint() {
	routergroup := r.Router.Group("/api/v1")

	routergroup.POST("/register", r.Register)
	routergroup.POST("/login", r.Login)
}

func (r *Rest) Run() {
	err := r.Router.Run(":8080")
	if err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
