package database

import (
	"auth/internal/config"
	"context"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabase(DBConfig *config.DBConfig, models ...any) (*gorm.DB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), DBConfig.DBTimeout)
	defer cancel()

	time.Sleep(1 * time.Second)

	db, err := gorm.Open(postgres.Open(DBConfig.DBDSN), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db.WithContext(ctx), nil

}
