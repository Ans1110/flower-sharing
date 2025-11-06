package database

import (
	"flower-backend/config"
	"os"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB(config *config.Config) {
	dsn := config.DatabaseURL
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		zap.L().Error("failed to connect to database", zap.Error(err))
		os.Exit(1)
	}
	DB = db
}
