package user_repository

import (
	"flower-backend/models"
	"time"

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
	if token.ExpiresAt.IsZero() {
		token.ExpiresAt = time.Now().Add(r.cfg.JWTRefreshExpiry)
	}
	return r.db.Create(token).Error
}

func (r *userRepository) DeleteExpiredTokens(now time.Time) (int64, error) {
	result := r.db.Where("expires_at < ?", now).Delete(&models.Token{})
	if result.Error != nil {
		r.logger.Error("failed to delete expired tokens", zap.Error(result.Error))
		return 0, result.Error
	}
	return result.RowsAffected, nil
}
