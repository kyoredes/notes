package main

import (
	"auth/internal/config"
	"auth/internal/handler"
	"auth/internal/logging"
	"auth/internal/server"
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/subosito/gotenv"
	"go.uber.org/zap"
)

func init() {
	logger := logging.Logger
	err := gotenv.Load("/auth/.env")
	if err != nil {
		logger.Fatal("Error loading .env file", zap.Error(err))
	}
}

func main() {
	config.Init()
	cfg := config.NewConfig()

	logging.InitLogger(cfg.LoggingMode)
	logger := logging.Logger

	logger.Info("Starting server...")

	h := handler.NewHandler(cfg)

	srv, err := server.NewServer(cfg, h)

	if err != nil {
		logger.Fatal("Error while starting server", zap.Error(err))
	}
	go func() {
		if err := srv.Start(); err != nil {
			logger.Fatal("Error while starting server", zap.Error(err))
		}
	}()

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
