package post_services

import (
	"flower-backend/libs"
	"flower-backend/models"
	"fmt"
	"io"
	"mime/multipart"
	"time"

	"go.uber.org/zap"
)

// UpdatePostByIDWithSelect
func (s *postService) UpdatePostByID(postId uint, userId uint, imageFile *multipart.FileHeader, updates map[string]any, selectFields []string) (*models.Post, error) {
	post, err := s.repo.UpdateByIDWithSelect(postId, updates, selectFields)
	if err != nil {
		s.logger.Error("failed to update post", zap.Error(err))
		return nil, err
	}

	if imageFile != nil {
		cld, err := libs.NewCloudinary(s.cfg)
		if err != nil {
			s.logger.Error("failed to create cloudinary client", zap.Error(err))
			return nil, err
		}
		oldPublicId := libs.ExtractPublicId(post.ImageURL)
		if err := libs.DeleteFromCloudinary(cld, oldPublicId); err != nil {
			s.logger.Error("failed to delete old image from cloudinary", zap.Error(err))
			return nil, err
		}

		src, err := imageFile.Open()
		if err != nil {
			s.logger.Error("failed to open image file", zap.Error(err))
			return nil, err
		}
		defer src.Close()

		buffer, err := io.ReadAll(src)
		if err != nil {
			s.logger.Error("failed to read image file", zap.Error(err))
			return nil, err
		}

		newPublicId := fmt.Sprintf("post_image_%d_%d", postId, time.Now().Unix())
		uploadResult, err := libs.UploadToCloudinary(cld, buffer, newPublicId)
		if err != nil {
			s.logger.Error("failed to upload image to cloudinary", zap.Error(err))
			return nil, err
		}
		post.ImageURL = uploadResult.SecureURL
		post.UpdatedAt = time.Now()

		if err := s.repo.Update(post); err != nil {
			s.logger.Error("failed to update post", zap.Error(err))
			return nil, err
		}
		s.logger.Info("post updated successfully", zap.Uint("id", postId))
		return post, nil
	}

	s.logger.Info("post updated successfully", zap.Uint("id", postId))
	return post, nil
}
