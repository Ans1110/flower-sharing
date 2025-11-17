package user_repository

import (
	"flower-backend/models"

	"go.uber.org/zap"
)

func (r *userRepository) Create(user *models.User) error {
	if err := r.db.Create(user).Error; err != nil {
		r.logger.Error("failed to create user", zap.Error(err))
		return err
	}
	return nil
}
