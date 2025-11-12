package v1_routes

import (
	"flower-backend/config"
	post_controller "flower-backend/controllers/v1/post"
	"flower-backend/database"
	"flower-backend/log"

	"github.com/gin-gonic/gin"
)

func PostRoutes(r *gin.RouterGroup) {
	cfg := config.LoadConfig()
	logger := log.InitLog().Sugar()
	postCtrl := post_controller.NewPostController(database.DB, cfg, logger)

	post := r.Group("/post")
	{
		// Create routes
		post.POST("", postCtrl.CreatePost)
		// Get routes
		post.GET("/:id", postCtrl.GetPostByID)
		post.GET("/user/:user_id/all", postCtrl.GetPostAllByUserID)
		post.GET("/all", postCtrl.GetPostAll)
		post.GET("/search", postCtrl.SearchPosts)
		post.GET("/pagination", postCtrl.GetPostWithPagination)
		// Delete routes
		post.DELETE("/:id", postCtrl.DeletePostByID)
		// Update routes
		post.PUT("/:id", postCtrl.UpdatePostByIDWithSelect)
		// Like routes
		post.POST("/:id/like", postCtrl.LikePost)
		post.DELETE("/:id/dislike", postCtrl.DislikePost)
		post.GET("/:id/likes", postCtrl.GetPostLikes)
		post.GET("/user/:user_id/liked", postCtrl.GetUserLikedPosts)
	}
}
