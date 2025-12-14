package admin_user_controller

import (
	"flower-backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// DELETE /api/v1/admin/user/:id
func (uc *adminUserController) DeleteUserByID(c *gin.Context) {
	userId := c.Param("id")
	userIdUint, err := utils.ParseUint(userId, uc.logger)
	if err != nil {
		utils.JSONError(c, http.StatusBadRequest, "ValidationError", err.Error())
		return
	}
	if err := uc.svc.DeleteUserByID(uint(userIdUint)); err != nil {
		if err == gorm.ErrRecordNotFound {
			uc.logger.Error("user not found", zap.String("user_id", userId))
			utils.JSONError(c, http.StatusNotFound, "NotFound", "User not found")
			return
		}
		utils.JSONError(c, http.StatusInternalServerError, "ServerError", "Internal server error")
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
	uc.logger.Info("user deleted successfully", zap.String("user_id", userId))
}
