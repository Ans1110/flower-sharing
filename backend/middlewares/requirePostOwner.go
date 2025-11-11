package middlewares

import (
	post_services "flower-backend/services/v1/post"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RequirePostOwner(postService *post_services.PostService) gin.HandlerFunc {
	return func(c *gin.Context) {
		userId, exists := c.Get("user_id")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    "Unauthorized",
				"message": "User not authenticated",
			})
			c.Abort()
			return
		}

		userIdUint, ok := userId.(uint)
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    "BadRequest",
				"message": "Invalid user id",
			})
			c.Abort()
			return
		}
		postId := c.Param("id")
		postIdUint, err := strconv.ParseUint(postId, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    "BadRequest",
				"message": "Invalid post id",
			})
			c.Abort()
			return
		}

		ownership, err := postService.CheckPostOwnership(uint(postIdUint), userIdUint)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    "ServerError",
				"message": "Internal server error",
			})
			c.Abort()
			return
		}
		if !ownership {
			c.JSON(http.StatusForbidden, gin.H{
				"code":    "Forbidden",
				"message": "You are not the owner of this post",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
