package user_services

import (
	"flower-backend/models"

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
