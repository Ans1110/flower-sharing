package middlewares

import (
	"flower-backend/database"
	"flower-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func Authorize(roles []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userId, exists := c.Get("userId")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    "Unauthorized",
				"message": "User not authenticated",
			})
			c.Abort()
			return
		}

		userIdUint, ok := userId.(uint)
		if !ok {
			zap.L().Error("Error while authorizing user: invalid userId type")
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    "ServerError",
				"message": "Internal server error",
			})
			c.Abort()
			return
		}

		// Fetch user from database
		var user models.User
		if err := database.DB.Select("role").Where("id = ?", userIdUint).First(&user).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				c.JSON(http.StatusNotFound, gin.H{
					"code":    "NotFound",
					"message": "User not found",
				})
				c.Abort()
				return
			}

			// Catch-all for other database errors
			zap.L().Error("Error while authorizing user", zap.Error(err))
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    "ServerError",
				"message": "Internal server error",
				"error":   err.Error(),
			})
			c.Abort()
			return
		}

		// Check if user's role is in allowed roles
		roleMap := make(map[string]bool, len(roles))
		for _, role := range roles {
			roleMap[role] = true
		}

		if !roleMap[user.Role] {
			c.JSON(http.StatusForbidden, gin.H{
				"code":    "AuthorizationError",
				"message": "Access denied, insufficient permissions",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
