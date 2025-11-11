package user_repository

import (
	"flower-backend/models"

	"go.uber.org/zap"
)

func (r *userRepository) Update(user *models.User) error {
	if err := r.db.Save(user).Error; err != nil {
		r.logger.Error("failed to update user", zap.Error(err))
		return err
	}
	return nil
}

func (r *userRepository) UpdateByID(id uint, user *models.User) error {
	if err := r.db.Save(user).Error; err != nil {
		r.logger.Error("failed to update user by id", zap.Error(err))
		return err
	}
	return nil
}

func (r *userRepository) UpdateByIDWithSelect(id uint, user *models.User, selectFields []string) error {
	if err := r.db.Select(selectFields).Where("id = ?", id).Save(user).Error; err != nil {
		r.logger.Error("failed to update user by id with select", zap.Error(err))
		return err
	}
	return nil
}

