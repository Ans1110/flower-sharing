package admin_user_controller

import (
	"flower-backend/config"
	user_services "flower-backend/services/v1/user"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type AdminUserController interface {
	// Get user operations
	GetUserByID(c *gin.Context)
	GetUserByEmail(c *gin.Context)
	GetUserByUsername(c *gin.Context)
	GetUserAll(c *gin.Context)
	GetUserByIDWithSelect(c *gin.Context)
	// Update user operations
	UpdateUserByIDWithSelect(c *gin.Context)
	// Delete user operations
	DeleteUserByID(c *gin.Context)
}

type adminUserController struct {
	svc    user_services.UserService
	cfg    *config.Config
	logger *zap.SugaredLogger
}

func NewAdminUserController(db *gorm.DB, cfg *config.Config, logger *zap.SugaredLogger) AdminUserController {
	svc := user_services.NewUserService(db, cfg, logger)
	return &adminUserController{svc: svc, logger: logger, cfg: cfg}
}
