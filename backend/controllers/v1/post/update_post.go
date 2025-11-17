package post_controller

import (
	"flower-backend/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// PUT /api/v1/post/:id?select=field1,field2,field3
func (pc *postController) UpdatePostByIDWithSelect(c *gin.Context) {

	postId := c.Param("id")
	postIdUint, err := utils.ParseUint(postId, pc.logger)
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

	selectQuery := c.Query("select")
	if selectQuery == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Select fields are required"})
		return
	}
	selectFields := strings.Split(selectQuery, ",")

	title := c.PostForm("title")
	content := c.PostForm("content")
	imageFile, err := c.FormFile("image")
	if err != nil {
		pc.logger.Error("failed to get image file", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get image file"})
		return
	}

	updates := make(map[string]any)
	if title != "" {
		updates["title"] = title
	}
	if content != "" {
		updates["content"] = content
	}

	updatedPost, err := pc.svc.UpdatePostByID(uint(postIdUint), userId, imageFile, updates, selectFields)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"post": updatedPost})
	pc.logger.Info("post updated successfully", zap.Uint("post_id", uint(postIdUint)))
}
