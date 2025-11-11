package user_repository

import (
	"flower-backend/models"

	"go.uber.org/zap"
)

func (r *userRepository) DeleteByID(id uint) error {
	if err := r.db.Delete(&models.User{}, id).Error; err != nil {
		r.logger.Error("failed to delete user by id", zap.Error(err))
		return err
	}
	return nil
}

func (r *userRepository) DeleteByEmail(email string) error {
	if err := r.db.Delete(&models.User{}, "email = ?", email).Error; err != nil {
		r.logger.Error("failed to delete user by email", zap.Error(err))
		return err
	}
	return nil
}

func (r *userRepository) DeleteByUsername(username string) error {
	if err := r.db.Delete(&models.User{}, "username = ?", username).Error; err != nil {
		r.logger.Error("failed to delete user by username", zap.Error(err))
		return err
	}
	return nil
}

