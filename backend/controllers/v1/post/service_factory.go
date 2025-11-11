package post

import (
	"errors"
	"flower-backend/config"
	"flower-backend/database"
	post_services "flower-backend/services/v1/post"
	"sync"
)

const (
	LogErrInitPostService = "failed to initialize post service"
	RespErrInternalServer = "Internal server error"
	LogErrPostNotFound    = "post not found"
	RespErrPostNotFound   = "Post not found"
	LogMsgPostFetched     = "post fetched successfully"
	LogErrBindPost        = "failed to bind post"
	LogMsgPostUpdated     = "post updated successfully"
	RespErrSelectRequired = "Select fields are required"
)

var (
	postServiceInstance *post_services.PostService
	postServiceOnce     sync.Once
)

func GetPostService() (*post_services.PostService, error) {
	if database.DB == nil {
		return nil, errors.New("database not initialized")
	}

	postServiceOnce.Do(func() {
		postServiceInstance = post_services.NewPostService(database.DB, config.LoadConfig())
	})

	if postServiceInstance == nil {
		return nil, errors.New("failed to initialize post service")
	}

	return postServiceInstance, nil
}
