package user_services

import (
	"flower-backend/models"

	"go.uber.org/zap"
)

// DeleteUserByID
func (s *UserService) DeleteUserByID(id uint) error {
	if err := s.db.Delete(&models.User{}, id).Error; err != nil {
		s.logger.Error("failed to delete user", zap.Error(err))
		return err
	}
	s.logger.Info("user deleted successfully", zap.Uint("id", id))
	return nil
}

// DeleteUserByEmail
func (s *UserService) DeleteUserByEmail(email string) error {
	if err := s.db.Delete(&models.User{}, "email = ?", email).Error; err != nil {
		s.logger.Error("failed to delete user", zap.Error(err))
		return err
	}
	s.logger.Info("user deleted successfully", zap.String("email", email))
	return nil
}

// DeleteUserByUsername
func (s *UserService) DeleteUserByUsername(username string) error {
	if err := s.db.Delete(&models.User{}, "username = ?", username).Error; err != nil {
		s.logger.Error("failed to delete user", zap.Error(err))
		return err
	}
	s.logger.Info("user deleted successfully", zap.String("username", username))
	return nil
}
