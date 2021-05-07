package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/zemags/gym_app/pkg/service"
)

type Handler struct {
	// dependency injection
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	// Register and authorization endpoints
	auth := router.Group("/auth")
	{
		auth.POST("sign-up", h.signUp)
		auth.POST("sign-in", h.signIn)
	}

	api := router.Group("/api", h.userIdentity)
	{
		exercises := api.Group("/exercises")
		{
			exercises.POST("/", h.createExerciseList)
		}
	}

	return router
}
