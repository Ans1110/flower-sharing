package user_service_factory

import (
	"errors"
	"sync"

	"flower-backend/config"
	"flower-backend/database"
	user_services "flower-backend/services/v1/user"
)

const (
	LogErrInitUserService = "failed to initialize user service"
	RespErrInternalServer = "Internal server error"
	LogErrUserNotFound    = "user not found"
	RespErrUserNotFound   = "User not found"
	LogMsgUserFetched     = "user fetched successfully"
	LogErrBindUser        = "failed to bind user"
	LogMsgUserUpdated     = "user updated successfully"
	RespErrSelectRequired = "Select fields are required"
)

var (
	userServiceInstance *user_services.UserService
	userServiceOnce     sync.Once
)

func GetUserService() (*user_services.UserService, error) {
	if database.DB == nil {
		return nil, errors.New("database not initialized")
	}

	userServiceOnce.Do(func() {
		userServiceInstance = user_services.NewUserService(database.DB, config.LoadConfig())
	})

	if userServiceInstance == nil {
		return nil, errors.New("failed to initialize user service")
	}

	return userServiceInstance, nil
}
