package router

import (
	"auth/internal/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter(h *handler.Handler) *gin.Engine {
	router := gin.Default()
	auth := router.Group("/auth")

	auth.POST("/create", h.CreateUser)
	return router
}
