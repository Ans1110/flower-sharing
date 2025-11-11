package post_services

import (
	"flower-backend/libs"
	"flower-backend/models"
	"fmt"
	"time"

	"go.uber.org/zap"
)

// CreatePost
func (s *PostService) CreatePost(post models.Post) (*models.Post, error) {
	if err := s.repo.Create(&post); err != nil {
		s.logger.Error("failed to create post", zap.Error(err))
		return nil, err
	}
	s.logger.Info("post created successfully", zap.String("title", post.Title))
	return &post, nil
}

// upload image
func (s *PostService) UploadImage(buffer []byte, postID uint) (string, error) {
	cld, err := libs.NewCloudinary(s.cfg)
	if err != nil {
		s.logger.Error("failed to create cloudinary client", zap.Error(err))
		return "", err
	}

	publicId := fmt.Sprintf("post_image_%d_%d", postID, time.Now().Unix())
	uploadResult, err := libs.UploadToCloudinary(cld, buffer, publicId)
	if err != nil {
		s.logger.Error("failed to upload image to cloudinary", zap.Error(err))
		return "", err
	}
	return uploadResult.SecureURL, nil
}
