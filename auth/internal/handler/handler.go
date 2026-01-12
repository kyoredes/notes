package handler

import (
	"auth/internal/config"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	cfg *config.Config
}

func NewHandler(cfg *config.Config) *Handler {
	return &Handler{
		cfg: cfg,
	}
}

func (h *Handler) CreateUser(c *gin.Context) {
	c.JSON(201, gin.H{
		"message": "User created",
	})
}
