package user_services

import (
	"errors"
	"flower-backend/models"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

// GetUserByID
func (s *userService) GetUserByID(id uint) (*models.User, error) {
	user, err := s.repo.GetByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			s.logger.Error("user not found", zap.Uint("id", id))
			return nil, gorm.ErrRecordNotFound
		}
		s.logger.Error("failed to get user by id", zap.Error(err))
		return nil, err
	}
	s.logger.Info("user fetched successfully", zap.Uint("id", id))
	return user, nil
}

// GetUserByEmail
func (s *userService) GetUserByEmail(email string) (*models.User, error) {
	user, err := s.repo.GetByEmail(email)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			s.logger.Error("user not found", zap.String("email", email))
			return nil, gorm.ErrRecordNotFound
		}
		s.logger.Error("failed to get user by email", zap.Error(err))
		return nil, err
	}
	s.logger.Info("user fetched successfully", zap.String("email", email))
	return user, nil
}

// GetUserByUsername
func (s *userService) GetUserByUsername(username string) (*models.User, error) {
	user, err := s.repo.GetByUsername(username)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			s.logger.Error("user not found", zap.String("username", username))
			return nil, gorm.ErrRecordNotFound
		}
		s.logger.Error("failed to get user by username", zap.Error(err))
		return nil, err
	}
	s.logger.Info("user fetched successfully", zap.String("username", username))
	return user, nil
}

// GetUserByIDWithSelect
func (s *userService) GetUserByIDWithSelect(id uint, selectFields []string) (*models.User, error) {
	user, err := s.repo.GetByIDWithSelect(id, selectFields)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			s.logger.Error("user not found", zap.Uint("id", id))
			return nil, gorm.ErrRecordNotFound
		}
		s.logger.Error("failed to get user by id", zap.Error(err))
		return nil, err
	}
	s.logger.Info("user fetched successfully", zap.Uint("id", id))
	return user, nil
}

// GetUserAll
func (s *userService) GetUserAll() ([]models.User, error) {
	users, err := s.repo.GetAll()
	if err != nil {
		s.logger.Error("failed to get all users", zap.Error(err))
		return nil, err
	}
	s.logger.Info("all users fetched successfully")
	return users, nil
}

// checking user ownership
func (s *userService) CheckUserOwnership(id uint, userID uint) (bool, error) {
	user, err := s.repo.GetByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			s.logger.Error("user not found", zap.Uint("id", id))
			return false, gorm.ErrRecordNotFound
		}
		s.logger.Error("failed to check user ownership", zap.Error(err))
		return false, err
	}
	if user.Role == "admin" {
		s.logger.Info("user owned by admin", zap.Uint("id", id))
		return true, nil
	}
	if user.ID != userID {
		s.logger.Error("user not owned by user", zap.Uint("id", id))
		return false, errors.New("user not owned by user")
	}
	s.logger.Info("user ownership checked successfully", zap.Uint("id", id))
	return true, nil
}
