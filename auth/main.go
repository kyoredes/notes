package main

import (
	"auth/internal/config"
	"auth/internal/logging"
	"auth/internal/server"
	"context"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	config.Init()
	cfg := config.NewConfig()

	logging.InitLogger(cfg.LoggingMode)
	logger := logging.Logger

	logger.Info("Starting server...")

	h := handler.NewHandler()

	srv, err := server.NewServer(cfg, h)

	if err != nil {
		logger.Fatal("Error while starting server")
	}
	go func() {
		if err := srv.Start(); err != nil {
			logger.Fatal("Error while starting server")
		}
	}

	logger.Info("Server started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Stop(ctx); err != nil {
		logger.Fatal("Error while stopping server")
	}

	logger.Info("Server stopped")



}

func pingHandler(context *gin.Context) {
	context.JSON(200, gin.H{"success": true})
}
