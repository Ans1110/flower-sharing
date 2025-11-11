package main

import (
	"context"
	"flower-backend/config"
	"flower-backend/database"
	"flower-backend/log"
	"flower-backend/middlewares"
	"flower-backend/models"
	v1Routes "flower-backend/routes/v1"
	"net/http"
	"os"
	"os/signal"
	"syscall"
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

	if err := db.AutoMigrate(&models.User{}, &models.Post{}, &models.Token{}); err != nil {
		logger.Error("failed to migrate database", zap.Error(err))
		os.Exit(1)
	}
	logger.Info("database migrated")
	// gin setup
	r := gin.New()
	r.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(logger, true))
	if cfg.GO_ENV == "production" {
		r.Use(middlewares.HttpLogger)
	}
	r.Use(func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				logger.Error("panic recovered", zap.Any("error", err))
				c.AbortWithStatusJSON(500, gin.H{"error": "internal server error"})
			}
		}()
		c.Next()
	})

	// helmet
	r.Use(middlewares.Helmet())

	// cors
	r.Use(cors.New(cors.Config{
		AllowOrigins:     cfg.AllowOrigins,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Rate limiter: 60 requests per minute per IP
	r.Use(middlewares.RateLimiter())

	// Setup routes
	v1Routes.Routes(r)

	port := cfg.Port
	if port == "" {
		port = "8080"
	}

	// Create HTTP server
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}

	// Start server in a goroutine
	go func() {
		logger.Info("server running on port http://localhost:" + port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error("failed to start server", zap.Error(err))
			if cfg.GO_ENV == "production" {
				os.Exit(1)
			}
		}
	}()

	logger.Info("server started")

	/**
	 * Listens for termination signals (`SIGTERM` and `SIGINT`).
	 *
	 * - `SIGTERM` is typically sent when stopping a process (e.g., `kill` command or container shutdown).
	 * - `SIGINT` is triggered when the user interrupts the process (e.g., pressing `Ctrl + C`).
	 * - When either signal is received, `handleServerShutdown` is executed to ensure proper cleanup.
	 */
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Info("shutting down server...")

	// Graceful shutdown with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Error("server forced to shutdown", zap.Error(err))
	}

	handleServerShutdown(logger)
}

/**
 * Handles server shutdown gracefully by disconnecting from the database.
 *
 * - Attempts to disconnect from the database before shutting down the server.
 * - Logs a success message if the disconnection is successful.
 * - If an error occurs during disconnection, it is logged to the console.
 * - Exits the process with status code `0` (indicating a successful shutdown).
 */
func handleServerShutdown(logger *zap.Logger) {
	logger.Info("disconnecting from database...")
	if err := database.DisconnectDB(); err != nil {
		logger.Error("failed to disconnect from database", zap.Error(err))
	} else {
		logger.Info("database disconnected successfully")
	}
	logger.Warn("server shutdown")
	os.Exit(0)
}
