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

func (r *userRepository) CreateToken(token *models.Token) error {
	if err := r.db.Where("user_id = ?", token.UserID).Delete(&models.Token{}).Error; err != nil {
		r.logger.Error("failed to delete token", zap.Error(err))
		return err
	}
	return r.db.Create(token).Error
}
