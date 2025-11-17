package public_user_controller

import (
	"flower-backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// DeleteUserByID godoc
//
//	@Summary		Delete user
//	@Description	Delete a user by ID
//	@Tags			users
//	@Produce		json
//	@Param			id	path		int						true	"User ID"
//	@Success		200	{object}	map[string]interface{}	"User deleted successfully"
//	@Failure		400	{object}	map[string]interface{}	"Bad request - invalid input"
//	@Failure		403	{object}	map[string]interface{}	"Forbidden - you are not the owner of this user"
//	@Failure		404	{object}	map[string]interface{}	"User not found"
//	@Failure		500	{object}	map[string]interface{}	"Internal server error"
//	@Securuty		BearerAuth
//	@Router			/user/{id} [delete]
func (uc *userController) DeleteUserByID(c *gin.Context) {
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
	if err := uc.svc.DeleteUserByID(uint(userIdUint)); err != nil {
		if err == gorm.ErrRecordNotFound {
			uc.logger.Error("user not found", zap.String("user_id", userId))
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
	uc.logger.Info("user deleted successfully", zap.String("user_id", userId))
}
