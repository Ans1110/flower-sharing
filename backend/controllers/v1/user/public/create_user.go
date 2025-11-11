package public_user_controller

import (
	user_service_factory "flower-backend/controllers/v1/user"
	"flower-backend/models"
	"flower-backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

// POST /api/v1/user
func CreateUser(c *gin.Context) {
	userService, err := user_service_factory.GetUserService()
	if err != nil {
		zap.L().Error(user_service_factory.LogErrInitUserService, zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": user_service_factory.RespErrInternalServer})
		return
	}

	username := c.PostForm("username")
	email := c.PostForm("email")
	password := c.PostForm("password")
	file, err := c.FormFile("avatar")
	if err != nil {
		zap.L().Error("failed to get avatar file", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get avatar file"})
		return
	}

	if username == "" || email == "" || password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username, email and password are required"})
		return
	}

	if !utils.ValidateUsername(username) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid username"})
		return
	}

	if !utils.ValidateEmail(email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		zap.L().Error("failed to hash password", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	var avatarURL string
	if file != nil {
		f, err := file.Open()
		if err != nil {
			zap.L().Error("failed to open avatar file", zap.Error(err))
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open avatar file"})
			return
		}
		defer f.Close()

		buffer := make([]byte, file.Size)
		_, err = f.Read(buffer)
		if err != nil {
			zap.L().Error("failed to read avatar file", zap.Error(err))
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read avatar file"})
			return
		}
		avatarURL, err = userService.UploadAvatar(buffer)
		if err != nil {
			zap.L().Error("failed to upload avatar", zap.Error(err))
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload avatar"})
			return
		}
	}

	user, err := userService.CreateUser(models.User{
		Username: username,
		Email:    email,
		Password: string(hashedPassword),
		Avatar:   avatarURL,
	})
	if err != nil {
		zap.L().Error("failed to create user", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
	zap.L().Info("user created successfully", zap.String("username", username))
}
