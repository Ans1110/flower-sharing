package user_services

import (
	"flower-backend/libs"
	"flower-backend/models"
	"flower-backend/utils"
	"fmt"
	"io"
	"mime/multipart"
	"time"

	"go.uber.org/zap"
)

// UpdateUserByIDWithSelect
func (s *userService) UpdateUserByIDWithSelect(id uint, updates map[string]any, imageFile *multipart.FileHeader, selectFields []string) (*models.User, error) {
	if username, ok := updates["username"].(string); ok {
		updates["username"] = utils.SanitizeUsername(username)
	}
	if email, ok := updates["email"].(string); ok {
		updates["email"] = utils.SanitizeEmail(email)
	}

	user, err := s.repo.UpdateByIDWithSelect(id, updates, selectFields)
	if err != nil {
		s.logger.Error("failed to update user", zap.Error(err))
		return nil, err
	}

	if imageFile != nil {
		cld, err := libs.NewCloudinary(s.cfg)
		if err != nil {
			s.logger.Error("failed to create cloudinary client", zap.Error(err))
			return nil, err
		}
		oldPublicId := libs.ExtractPublicId(user.Avatar)
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

		newPublicId := fmt.Sprintf("avatar_%d_%d", id, time.Now().Unix())
		uploadResult, err := libs.UploadToCloudinary(cld, buffer, newPublicId)
		if err != nil {
			s.logger.Error("failed to upload image to cloudinary", zap.Error(err))
			return nil, err
		}
		user.Avatar = uploadResult.SecureURL

		if err := s.repo.Update(user); err != nil {
			s.logger.Error("failed to update user", zap.Error(err))
			return nil, err
		}
		s.logger.Info("user updated successfully", zap.Uint("id", id))
		return user, nil
	}
	s.logger.Info("user updated successfully", zap.Uint("id", id))
	return user, nil
}
