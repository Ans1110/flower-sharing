package public_user_controller

import (
	public_user_dto "flower-backend/dto/public"

	"flower-backend/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// UpdateUserByIDWithSelect godoc
//
//	@Summary		Update user profile
//	@Description	Update specific fields of user profile
//	@Tags			users
//	@Accept			multipart/form-data
//	@Produce		json
//	@Param			id			path		int						true	"User ID"
//	@Param			select		query		string					true	"Fields to update (comma-separated)"
//	@Param			username	formData	string					false	"Username"
//	@Param			email		formData	string					false	"Email"
//	@Param			image		formData	file					false	"Avatar image"
//	@Success		200			{object}	map[string]interface{}	"User updated successfully"
//	@Failure		400			{object}	map[string]interface{}	"Bad request - invalid input"
//	@Failure		403			{object}	map[string]interface{}	"Forbidden - you are not the owner of this user"
//	@Failure		500			{object}	map[string]interface{}	"Internal server error"
//	@Securuty		BearerAuth
//	@Router			/user/id/{id}/select [put]
func (uc *userController) UpdateUserByIDWithSelect(c *gin.Context) {

	userId := c.Param("id")
	userIdUint, err := utils.ParseUint(userId, uc.logger)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ownership, err := uc.svc.CheckUserOwnership(uint(userIdUint), c.GetUint("user_id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check user ownership"})
		return
	}
	if !ownership {
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not the owner of this user"})
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
		uc.logger.Error("failed to get image file", zap.Error(err))
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
