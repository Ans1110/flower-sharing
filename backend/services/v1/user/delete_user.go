package user_services

import (
	"go.uber.org/zap"
)

// DeleteUserByID
func (s *userService) DeleteUserByID(id uint) error {
	if err := s.repo.DeleteByID(id); err != nil {
		s.logger.Error("failed to delete user", zap.Error(err))
		return err
	}
	s.logger.Info("user deleted successfully", zap.Uint("id", id))
	return nil
}
