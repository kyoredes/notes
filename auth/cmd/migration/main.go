package main

import (
	"auth/cmd/migration/script"
	"auth/internal/config"
	"auth/internal/database"
	"auth/internal/logging"
	"auth/internal/models"
	"fmt"

	"github.com/subosito/gotenv"
	"go.uber.org/zap"
)

var ModelsList = []any{
	&models.User{},
}

func main() {
	config.Init()
	cfg := config.NewConfig()

	logging.InitLogger(cfg.LoggingMode)
	logger := logging.Logger
	err := gotenv.Load(".env")

	if err != nil {
		logger.Fatal("Error loading .env file", zap.Error(err))
	}
	db, err := database.NewDatabase(config.NewDBConfig(), ModelsList)
	if err != nil {
		fmt.Println(err)
		return
	}
	if err := script.RunMigrations(db, ModelsList); err != nil {
		fmt.Println(err)
		return
	}
}
