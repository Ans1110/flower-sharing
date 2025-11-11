package user_repository

import (
	"flower-backend/models"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

func (r *userRepository) GetByID(id uint) (*models.User, error) {
	var user models.User
	if err := r.db.Where("id = ?", id).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, gorm.ErrRecordNotFound
		}
		r.logger.Error("failed to get user by id", zap.Error(err))
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetByEmail(email string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, gorm.ErrRecordNotFound
		}
		r.logger.Error("failed to get user by email", zap.Error(err))
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetByUsername(username string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("username = ?", username).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, gorm.ErrRecordNotFound
		}
		r.logger.Error("failed to get user by username", zap.Error(err))
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetByIDWithSelect(id uint, selectFields []string) (*models.User, error) {
	var user models.User
	if err := r.db.Select(selectFields).Where("id = ?", id).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, gorm.ErrRecordNotFound
		}
		r.logger.Error("failed to get user by id with select", zap.Error(err))
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetByEmailWithSelect(email string, selectFields []string) (*models.User, error) {
	var user models.User
	if err := r.db.Select(selectFields).Where("email = ?", email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, gorm.ErrRecordNotFound
		}
		r.logger.Error("failed to get user by email with select", zap.Error(err))
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetByUsernameWithSelect(username string, selectFields []string) (*models.User, error) {
	var user models.User
	if err := r.db.Select(selectFields).Where("username = ?", username).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, gorm.ErrRecordNotFound
		}
		r.logger.Error("failed to get user by username with select", zap.Error(err))
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetAll() ([]models.User, error) {
	var users []models.User
	if err := r.db.Find(&users).Error; err != nil {
		r.logger.Error("failed to get all users", zap.Error(err))
		return nil, err
	}
	return users, nil
}

