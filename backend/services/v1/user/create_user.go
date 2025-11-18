package user_services

import (
	"flower-backend/libs"
	"flower-backend/models"
	"flower-backend/utils"
	"fmt"
	"mime/multipart"
	"time"

	"go.uber.org/zap"
)

// CreateUser
func (s *userService) CreateUser(user models.User) (*models.User, error) {
	if err := s.repo.Create(&user); err != nil {
		s.logger.Error("failed to create user", zap.Error(err))
		return nil, err
	}
	s.logger.Info("user created successfully", zap.String("username", user.Username))
	return &user, nil
}

// upload avatar
func (s *userService) UploadAvatar(buffer []byte, userID uint) (string, error) {
	cld, err := libs.NewCloudinary(s.cfg)
	if err != nil {
		s.logger.Error("failed to create cloudinary client", zap.Error(err))
		return "", err
	}
	publicId := fmt.Sprintf("avatar_%d_%d", userID, time.Now().Unix())

	uploadResult, err := libs.UploadToCloudinary(cld, buffer, publicId)
	if err != nil {
		s.logger.Error("failed to upload avatar to cloudinary", zap.Error(err))
		return "", err
	}
	return uploadResult.SecureURL, nil
}

// register user
func (s *userService) RegisterUser(username, email, password string, avatarFile *multipart.FileHeader) (*models.User, error) {
	username = utils.SanitizeUsername(username)
	email = utils.SanitizeEmail(email)

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
		publicId := fmt.Sprintf("avatar_%d", time.Now().Unix())
		uploadResult, err := libs.UploadToCloudinary(cld, buffer, publicId)
		if err != nil {
			s.logger.Error("failed to upload avatar to cloudinary", zap.Error(err))
			return nil, err
		}
		avatarURL = uploadResult.SecureURL
	}

	// Determine user role based on whitelist
	role := "user"
	for _, adminEmail := range s.cfg.WhiteListAdminEmails {
		if email == adminEmail {
			role = "admin"
			break
		}
	}

	user := models.User{
		Username: username,
		Email:    email,
		Password: password,
		Avatar:   avatarURL,
		Role:     role,
	}

	if err := s.repo.Create(&user); err != nil {
		s.logger.Error("failed to create user", zap.Error(err))
		return nil, err
	}
	s.logger.Info("user created successfully", zap.String("username", username), zap.String("role", role))
	return &user, nil
}

// create token
func (s *userService) CreateToken(token *models.Token) error {
	if err := s.repo.CreateToken(token); err != nil {
		s.logger.Error("failed to create token", zap.Error(err))
		return err
	}
	return nil
}
