package auth_controller

import (
	"flower-backend/config"
	user_services "flower-backend/services/v1/user"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type AuthController interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
	Logout(c *gin.Context)
	RefreshToken(c *gin.Context)
}

type authController struct {
	svc    user_services.UserService
	cfg    *config.Config
	logger *zap.SugaredLogger
}

func NewAuthController(db *gorm.DB, cfg *config.Config, logger *zap.SugaredLogger) AuthController {
	svc := user_services.NewUserService(db, cfg, logger)
	return &authController{svc: svc, logger: logger, cfg: cfg}
}
