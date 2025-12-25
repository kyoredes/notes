package main

import (
	"auth/internal"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	cfg := internal.NewConfig()
	DBConfig := internal.NewDBConfig()
	logger := internal.Logger

	logger.Debug("config", zap.Any("cfg", cfg))
	logger.Debug("db config", zap.Any("db config", DBConfig))

	router := gin.Default()

	router.GET("/ping", pingHandler)

	router.Run(":8000")
}

func pingHandler(context *gin.Context) {
	context.JSON(200, gin.H{"success": true})
}
