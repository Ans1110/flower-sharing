package routes

import (
	"flower-backend/controllers"
	"flower-backend/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine) {
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	protected := r.Group("/api")
	protected.Use(middlewares.AuthMiddleware())
	{
		protected.GET("/me", func(ctx *gin.Context) {
			userId := ctx.MustGet("userId").(uint)
			role := ctx.MustGet("role").(string)
			ctx.JSON(http.StatusOK, gin.H{"userId": userId, "role": role})
		})
	}
}
