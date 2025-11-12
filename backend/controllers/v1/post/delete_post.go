package post_controller

import (
	"flower-backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// DELETE /api/v1/post/:id
func (pc *postController) DeletePostByID(c *gin.Context) {

	postId := c.Param("id")
	postIdUint, err := utils.ParseUint(postId, zap.L().Sugar())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userId := c.GetUint("user_id")
	ownership, err := pc.svc.CheckPostOwnership(uint(postIdUint), userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check post ownership"})
		return
	}
	if !ownership {
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not the owner of this post"})
		return
	}
	if err := pc.svc.DeletePostByID(uint(postIdUint), userId); err != nil {
		if err == gorm.ErrRecordNotFound {
			zap.L().Error("post not found", zap.String("post_id", postId))
			c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete post"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Post deleted successfully"})
	zap.L().Info("post deleted successfully", zap.String("post_id", postId), zap.Uint("user_id", userId))
}
