package public_user_controller

import (
	public_user_dto "flower-backend/dto/public"

	"flower-backend/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// PUT /api/v1/user/id/:id/select?select=field1,field2,field3
func (uc *userController) UpdateUserByIDWithSelect(c *gin.Context) {

	userId := c.Param("id")
	userIdUint, err := utils.ParseUint(userId, zap.L().Sugar())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	selectQuery := c.Query("select")
	if selectQuery == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Select fields are required"})
		return
	}
	selectFields := strings.Split(selectQuery, ",")

	username := c.PostForm("username")
	email := c.PostForm("email")
	imageFile, err := c.FormFile("image")
	if err != nil {
		zap.L().Error("failed to get image file", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get image file"})
		return
	}

	updates := make(map[string]any)
	if username != "" {
		updates["username"] = username
	}
	if email != "" {
		updates["email"] = email
	}

	updatedUser, err := uc.svc.UpdateUserByIDWithSelect(uint(userIdUint), updates, imageFile, selectFields)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": public_user_dto.ToPublicUser(updatedUser)})
	uc.logger.Info("user updated successfully", zap.Uint("user_id", uint(userIdUint)))
}
