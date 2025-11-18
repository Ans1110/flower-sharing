package database

import (
	"flower-backend/config"
	"fmt"
	"os"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

var DB *gorm.DB

// ConnectDB connects to the database with connection pooling
func ConnectDB(cfg *config.Config, logger *zap.Logger) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)

	// Configure GORM with custom logger settings
	gormConfig := &gorm.Config{}

	// Use different log levels based on environment
	if cfg.GO_ENV == "production" {
		gormConfig.Logger = gormlogger.Default.LogMode(gormlogger.Silent)
	} else {
		gormConfig.Logger = gormlogger.Default.LogMode(gormlogger.Warn)
	}

	db, err := gorm.Open(mysql.Open(dsn), gormConfig)
	if err != nil {
		logger.Error("failed to connect to database", zap.Error(err))
		os.Exit(1)
	}

	// Get the underlying sql.DB to configure connection pooling
	sqlDB, err := db.DB()
	if err != nil {
		logger.Error("failed to get database instance", zap.Error(err))
		os.Exit(1)
	}

	// Configure connection pool settings
	// SetMaxOpenConns sets the maximum number of open connections to the database
	sqlDB.SetMaxOpenConns(cfg.DBMaxOpenConns)

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool
	sqlDB.SetMaxIdleConns(cfg.DBMaxIdleConns)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused
	sqlDB.SetConnMaxLifetime(cfg.DBConnMaxLifetime)

	// SetConnMaxIdleTime sets the maximum amount of time a connection may be idle
	sqlDB.SetConnMaxIdleTime(cfg.DBConnMaxIdleTime)

	// Verify the connection
	if err := sqlDB.Ping(); err != nil {
		logger.Error("failed to ping database", zap.Error(err))
		os.Exit(1)
	}

	DB = db

	// Log connection pool configuration
	logger.Info("database connection pool configured",
		zap.Int("max_open_connections", cfg.DBMaxOpenConns),
		zap.Int("max_idle_connections", cfg.DBMaxIdleConns),
		zap.Duration("connection_max_lifetime", cfg.DBConnMaxLifetime),
		zap.Duration("connection_max_idle_time", cfg.DBConnMaxIdleTime),
	)
}

// DisconnectDB closes the database connection gracefully
func DisconnectDB(logger *zap.Logger) error {
	if DB != nil {
		sqlDB, err := DB.DB()
		if err != nil {
			return err
		}

		// Log final connection pool stats before closing
		stats := sqlDB.Stats()
		logger.Info("closing database connection",
			zap.Int("open_connections", stats.OpenConnections),
			zap.Int("in_use", stats.InUse),
			zap.Int("idle", stats.Idle),
		)

		return sqlDB.Close()
	}
	return nil
}
