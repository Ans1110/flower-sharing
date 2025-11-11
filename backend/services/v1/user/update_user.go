package user_services

import (
	"flower-backend/models"

	"go.uber.org/zap"
)

// UpdateUser
func (s *UserService) UpdateUser(user models.User) (*models.User, error) {
	if err := s.repo.Update(&user); err != nil {
		s.logger.Error("failed to update user", zap.Error(err))
		return nil, err
	}
	s.logger.Info("user updated successfully", zap.String("username", user.Username))
	return &user, nil
}

// UpdateUserByID
func (s *UserService) UpdateUserByID(id uint, user models.User) (*models.User, error) {
	if err := s.repo.UpdateByID(id, &user); err != nil {
		s.logger.Error("failed to update user", zap.Error(err))
		return nil, err
	}
	s.logger.Info("user updated successfully", zap.String("username", user.Username))
	return &user, nil
}

// UpdateUserByEmail
func (s *UserService) UpdateUserByEmail(email string, user models.User) (*models.User, error) {
	if err := s.repo.Update(&user); err != nil {
		s.logger.Error("failed to update user", zap.Error(err))
		return nil, err
	}
	s.logger.Info("user updated successfully", zap.String("username", user.Username))
	return &user, nil
}

// UpdateUserByUsername
func (s *UserService) UpdateUserByUsername(username string, user models.User) (*models.User, error) {
	if err := s.repo.Update(&user); err != nil {
		s.logger.Error("failed to update user", zap.Error(err))
		return nil, err
	}
	s.logger.Info("user updated successfully", zap.String("username", user.Username))
	return &user, nil
}

// UpdateUserByIDWithSelect
func (s *UserService) UpdateUserByIDWithSelect(id uint, user models.User, selectFields []string) (*models.User, error) {
	if err := s.repo.UpdateByIDWithSelect(id, &user, selectFields); err != nil {
		s.logger.Error("failed to update user", zap.Error(err))
		return nil, err
	}
	s.logger.Info("user updated successfully", zap.Uint("id", id))
	return &user, nil
}
