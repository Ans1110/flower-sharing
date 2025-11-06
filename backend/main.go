package main

import (
	"flower-backend/config"
	"flower-backend/database"
	"flower-backend/log"
	"flower-backend/models"
	v1Routes "flower-backend/routes/v1"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	// log
	logger := log.InitLog()
	defer logger.Sync()
	// config
	cfg := config.LoadConfig()

	logger.Info("starting server")
	// db
	database.ConnectDB(cfg)
	db := database.DB

	// Clean up orphaned posts before migration
	// Disable foreign key checks temporarily to allow cleanup
	db.Exec("SET FOREIGN_KEY_CHECKS = 0")

	if db.Migrator().HasTable(&models.Post{}) && db.Migrator().HasTable(&models.User{}) {
		// Delete posts with invalid user_id references
		result := db.Exec(`
			DELETE FROM posts 
			WHERE user_id NOT IN (SELECT id FROM users) 
			AND user_id IS NOT NULL
		`)
		if result.Error == nil && result.RowsAffected > 0 {
			logger.Info("cleaned up orphaned posts", zap.Int64("count", result.RowsAffected))
		}
	}

	// Run migration with foreign key checks disabled
	if err := db.AutoMigrate(&models.User{}, &models.Post{}, &models.Token{}); err != nil {
		// Re-enable checks before exiting on error
		db.Exec("SET FOREIGN_KEY_CHECKS = 1")
		logger.Error("failed to migrate database", zap.Error(err))
		os.Exit(1)
	}

	// Re-enable foreign key checks after successful migration
	db.Exec("SET FOREIGN_KEY_CHECKS = 1")
	logger.Info("database migrated")
	// gin setup
	r := gin.New()
	r.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(logger, true))
	r.Use(func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				logger.Error("panic recovered", zap.Any("error", err))
				c.AbortWithStatusJSON(500, gin.H{"error": "internal server error"})
			}
		}()
		c.Next()
	})

	// cors
	r.Use(cors.New(cors.Config{
		AllowOrigins:     cfg.AllowOrigins,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Setup routes
	v1Routes.Routes(r)

	port := cfg.Port
	if port == "" {
		port = "8080"
	}
	logger.Info("server running on port http://localhost:" + port)

	if err := r.Run(":" + port); err != nil {
		logger.Error("failed to start server", zap.Error(err))
		if cfg.GO_ENV == "production" {
			os.Exit(1)
		}
	}

	logger.Info("server started")
}
