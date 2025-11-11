package user_services

import (
	"flower-backend/models"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

// GetUserByID
func (s *UserService) GetUserByID(id uint) (*models.User, error) {
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
func (s *UserService) GetUserByEmail(email string) (*models.User, error) {
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
func (s *UserService) GetUserByUsername(username string) (*models.User, error) {
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
func (s *UserService) GetUserByIDWithSelect(id uint, selectFields []string) (*models.User, error) {
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

// GetUserByEmailWithSelect
func (s *UserService) GetUserByEmailWithSelect(email string, selectFields []string) (*models.User, error) {
	user, err := s.repo.GetByEmailWithSelect(email, selectFields)
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

// GetUserByUsernameWithSelect
func (s *UserService) GetUserByUsernameWithSelect(username string, selectFields []string) (*models.User, error) {
	user, err := s.repo.GetByUsernameWithSelect(username, selectFields)
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

// GetUserAll
func (s *UserService) GetUserAll() ([]models.User, error) {
	users, err := s.repo.GetAll()
	if err != nil {
		s.logger.Error("failed to get all users", zap.Error(err))
		return nil, err
	}
	s.logger.Info("all users fetched successfully")
	return users, nil
}
