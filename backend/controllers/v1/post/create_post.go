package post_controller

import (
	"flower-backend/models"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// POST /api/v1/post
func (pc *postController) CreatePost(c *gin.Context) {

	userId := c.GetUint("user_id")
	title := c.PostForm("title")
	content := c.PostForm("content")
	imageFile, err := c.FormFile("image")
	if err != nil {
		pc.logger.Error("failed to get image file", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get image file"})
		return
	}

	if title == "" || content == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Title and content are required"})
		return
	}

	var imageURL string

	if imageFile != nil {
		src, err := imageFile.Open()
		if err != nil {
			pc.logger.Error("failed to open image file", zap.Error(err))
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open image file"})
			return
		}
		defer src.Close()

		buffer, err := io.ReadAll(src)
		if err != nil {
			pc.logger.Error("failed to read image file", zap.Error(err))
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read image file"})
			return
		}

		imageURL, err = pc.svc.UploadImage(buffer, userId)
		if err != nil {
			pc.logger.Error("failed to upload image", zap.Error(err))
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload image"})
			return
		}
	}

	post, err := pc.svc.CreatePost(models.Post{
		Title:    title,
		Content:  content,
		ImageURL: imageURL,
		UserID:   userId,
	})
	if err != nil {
		pc.logger.Error("failed to create post", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create post"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"post": post})
	pc.logger.Info("post created successfully", zap.String("title", post.Title))
}
