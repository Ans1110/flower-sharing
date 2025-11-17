package post_controller

import (
	"flower-backend/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// UpdatePostByIDWithSelect godoc
//
//	@Summary		Update post by ID with select
//	@Description	Update a post by ID with select fields
//	@Tags			posts
//	@Accept			multipart/form-data
//	@Produce		json
//	@Param			id		path		int						true	"Post ID"
//	@Param			select	query		string					true	"Fields to update (comma-separated)"
//	@Param			title	formData	string					false	"Post title"
//	@Param			content	formData	string					false	"Post content"
//	@Param			image	formData	file					false	"Post image"
//	@Success		200		{object}	map[string]interface{}	"Post updated successfully"
//	@Failure		400		{object}	map[string]interface{}	"Bad request - invalid input"
//	@Failure		403		{object}	map[string]interface{}	"Forbidden - you are not the owner of this post"
//	@Failure		500		{object}	map[string]interface{}	"Internal server error"
//	@Securuty		BearerAuth
//	@Router			/post/{id} [put]
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
