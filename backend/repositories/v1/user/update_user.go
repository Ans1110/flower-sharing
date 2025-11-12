package user_repository

import (
	"flower-backend/models"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

func (r *userRepository) Update(user *models.User) error {
	if err := r.db.Save(user).Error; err != nil {
		r.logger.Error("failed to update user", zap.Error(err))
		return err
	}
	return nil
}

func (r *userRepository) UpdateByIDWithSelect(id uint, updates map[string]any, selectFields []string) (*models.User, error) {
	var user models.User

	if err := r.db.First(&user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, gorm.ErrRecordNotFound
		}
		r.logger.Error("failed to find user", zap.Error(err))
		return nil, err
	}

	filtered := make(map[string]any)
	for _, field := range selectFields {
		if val, ok := updates[field]; ok {
			filtered[field] = val
		}
	}

	if len(filtered) == 0 {
		return &user, nil
	}

	if err := r.db.First(&user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, gorm.ErrRecordNotFound
		}
		r.logger.Error("failed to find user", zap.Error(err))
		return nil, err
	}

	r.logger.Info("user updated successfully", zap.Uint("id", id))
	return &user, nil
}
