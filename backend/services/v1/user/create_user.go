package user_services

import (
	"flower-backend/libs"
	"flower-backend/models"
	"fmt"
	"mime/multipart"
	"time"

	"go.uber.org/zap"
)

// CreateUser
func (s *UserService) CreateUser(user models.User) (*models.User, error) {
	if err := s.db.Create(&user).Error; err != nil {
		s.logger.Error("failed to create user", zap.Error(err))
		return nil, err
	}
	s.logger.Info("user created successfully", zap.String("username", user.Username))
	return &user, nil
}

// upload avatar
func (s *UserService) UploadAvatar(buffer []byte) (string, error) {
	cld, err := libs.NewCloudinary(s.cfg)
	publicId := fmt.Sprintf("avatar_%d_%d", s.user.ID, time.Now().Unix())
	if err != nil {
		s.logger.Error("failed to create cloudinary client", zap.Error(err))
		return "", err
	}

	uploadResult, err := libs.UploadToCloudinary(cld, buffer, publicId)
	if err != nil {
		s.logger.Error("failed to upload avatar to cloudinary", zap.Error(err))
		return "", err
	}
	return uploadResult.SecureURL, nil
}

// register user
func (s *UserService) RegisterUser(username, email, password string, avatarFile *multipart.FileHeader) (*models.User, error) {
	var avatarURL string

	if avatarFile != nil {
		f, err := avatarFile.Open()
		if err != nil {
			s.logger.Error("failed to open avatar file", zap.Error(err))
			return nil, err
		}
		defer f.Close()

		buffer := make([]byte, avatarFile.Size)
		_, err = f.Read(buffer)
		if err != nil {
			s.logger.Error("failed to read avatar file", zap.Error(err))
			return nil, err
		}
		cld, err := libs.NewCloudinary(s.cfg)
		if err != nil {
			s.logger.Error("failed to create cloudinary client", zap.Error(err))
			return nil, err
		}
		publicId := fmt.Sprintf("avatar_%d_%d", s.user.ID, time.Now().Unix())
		uploadResult, err := libs.UploadToCloudinary(cld, buffer, publicId)
		if err != nil {
			s.logger.Error("failed to upload avatar to cloudinary", zap.Error(err))
			return nil, err
		}
		avatarURL = uploadResult.SecureURL
	}

	user := models.User{
		Username: username,
		Email:    email,
		Password: password,
		Avatar:   avatarURL,
	}

	if err := s.db.Create(&user).Error; err != nil {
		s.logger.Error("failed to create user", zap.Error(err))
		return nil, err
	}
	s.logger.Info("user created successfully", zap.String("username", username))
	return &user, nil
}
