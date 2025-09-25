package routes

import (
	"flower-backend/controllers"
	"flower-backend/middlewares"

	"github.com/gin-gonic/gin"
)

func FlowerRoutes(r *gin.Engine) {
	api := r.Group("/api/flowers")
	{
		api.GET("", controllers.ListFlowers)
		api.GET("/:id", controllers.GetFlower)
	}

	r.GET("/api/search", controllers.SearchPosts)

	auth := r.Group("/api/flowers")
	auth.Use(middlewares.AuthMiddleware())
	{
		auth.GET("/user", controllers.GetUser)
		auth.POST("", controllers.CreateFlower)
		auth.PUT("/:id", controllers.UpdateFlower)
		auth.POST("/:id/like", controllers.LikeFlower)
		auth.DELETE("/:id/unlike", controllers.UnlikeFlower)
		auth.DELETE("/:id", controllers.DeleteFlower)
	}
}
