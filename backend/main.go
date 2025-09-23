package main

import (
	"flower-backend/database"
	"flower-backend/models"
	"flower-backend/routes"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	//db
	database.ConnectDB()
	database.DB.AutoMigrate(&models.User{}, &models.Post{})

	//cors
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	//routes
	routes.AuthRoutes(r)
	routes.FlowerRoutes(r)

	r.Run(":8080")
}
