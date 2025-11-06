package user_services

import (
	"flower-backend/models"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

// GetUserByID
func (s *UserService) GetUserByID(id uint) (*models.User, error) {
	if err := s.db.Where("id = ?", id).First(&s.user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			s.logger.Error("user not found", zap.Uint("id", id))
			return nil, gorm.ErrRecordNotFound
		}
		s.logger.Error("failed to get user by id", zap.Error(err))
		return nil, err
	}
	s.logger.Info("user fetched successfully", zap.Uint("id", id))
	return &s.user, nil
}

// GetUserByEmail
func (s *UserService) GetUserByEmail(email string) (*models.User, error) {
	if err := s.db.Where("email = ?", email).First(&s.user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			s.logger.Error("user not found", zap.String("email", email))
			return nil, gorm.ErrRecordNotFound
		}
		s.logger.Error("failed to get user by email", zap.Error(err))
		return nil, err
	}
	s.logger.Info("user fetched successfully", zap.String("email", email))
	return &s.user, nil
}

// GetUserByUsername
func (s *UserService) GetUserByUsername(username string) (*models.User, error) {
	if err := s.db.Where("username = ?", username).First(&s.user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			s.logger.Error("user not found", zap.String("username", username))
			return nil, gorm.ErrRecordNotFound
		}
		s.logger.Error("failed to get user by username", zap.Error(err))
		return nil, err
	}
	s.logger.Info("user fetched successfully", zap.String("username", username))
	return &s.user, nil
}

// GetUserByIDWithSelect
func (s *UserService) GetUserByIDWithSelect(id uint, selectFields []string) (*models.User, error) {
	if err := s.db.Select(selectFields).Where("id = ?", id).First(&s.user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			s.logger.Error("user not found", zap.Uint("id", id))
			return nil, gorm.ErrRecordNotFound
		}
		s.logger.Error("failed to get user by id", zap.Error(err))
		return nil, err
	}
	s.logger.Info("user fetched successfully", zap.Uint("id", id))
	return &s.user, nil
}

// GetUserByEmailWithSelect
func (s *UserService) GetUserByEmailWithSelect(email string, selectFields []string) (*models.User, error) {
	if err := s.db.Select(selectFields).Where("email = ?", email).First(&s.user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			s.logger.Error("user not found", zap.String("email", email))
			return nil, gorm.ErrRecordNotFound
		}
		s.logger.Error("failed to get user by email", zap.Error(err))
		return nil, err
	}
	s.logger.Info("user fetched successfully", zap.String("email", email))
	return &s.user, nil
}

// GetUserByUsernameWithSelect
func (s *UserService) GetUserByUsernameWithSelect(username string, selectFields []string) (*models.User, error) {
	if err := s.db.Select(selectFields).Where("username = ?", username).First(&s.user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			s.logger.Error("user not found", zap.String("username", username))
			return nil, gorm.ErrRecordNotFound
		}
		s.logger.Error("failed to get user by username", zap.Error(err))
		return nil, err
	}
	s.logger.Info("user fetched successfully", zap.String("username", username))
	return &s.user, nil
}

// GetUserAll
func (s *UserService) GetUserAll() ([]models.User, error) {
	var users []models.User
	if err := s.db.Find(&users).Error; err != nil {
		s.logger.Error("failed to get all users", zap.Error(err))
		return nil, err
	}
	s.logger.Info("all users fetched successfully")
	return users, nil
}
