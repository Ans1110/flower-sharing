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
		utils.JSONError(c, http.StatusBadRequest, "ValidationError", err.Error())
		return
	}
	ownership, err := uc.svc.CheckUserOwnership(uint(userIdUint), c.GetUint("user_id"))
	if err != nil {
		utils.JSONError(c, http.StatusInternalServerError, "ServerError", "Failed to check user ownership")
		return
	}
	if !ownership {
		utils.JSONError(c, http.StatusForbidden, "Forbidden", "You are not the owner of this user")
		return
	}
	if err := uc.svc.DeleteUserByID(uint(userIdUint)); err != nil {
		if err == gorm.ErrRecordNotFound {
			uc.logger.Error("user not found", zap.String("user_id", userId))
			utils.JSONError(c, http.StatusNotFound, "NotFound", "User not found")
			return
		}
		utils.JSONError(c, http.StatusInternalServerError, "ServerError", "Failed to delete user")
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
	uc.logger.Info("user deleted successfully", zap.String("user_id", userId))
}
