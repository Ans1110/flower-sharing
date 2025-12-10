package auth_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// Me godoc
//
//	@Summary		Get current user
//	@Description	Get the authenticated user's information
//	@Tags			auth
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}	"User fetched successfully"
//	@Failure		401	{object}	map[string]interface{}	"Unauthorized"
//	@Failure		500	{object}	map[string]interface{}	"Internal server error"
//	@Security		BearerAuth
//	@Router			/auth/me [get]
func (ac *authController) Me(c *gin.Context) {
	// Get user ID from context (set by auth middleware as "user_id")
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	user, err := ac.svc.GetUserByID(userID.(uint))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			ac.logger.Error("user not found", zap.Any("user_id", userID))
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		ac.logger.Error("failed to get user", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": gin.H{
			"id":         user.ID,
			"username":   user.Username,
			"email":      user.Email,
			"role":       user.Role,
			"avatar":     user.Avatar,
			"created_at": user.CreatedAt,
			"provider":   user.Provider,
		},
	})

	ac.logger.Info("User fetched successfully", zap.Uint("user_id", user.ID))
}
