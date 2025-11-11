package post_services

import (
	"go.uber.org/zap"
)

// DeletePostByID
func (s *PostService) DeletePostByID(id uint) error {
	if err := s.repo.DeleteByID(id); err != nil {
		s.logger.Error("failed to delete post", zap.Error(err))
		return err
	}
	s.logger.Info("post deleted successfully", zap.Uint("id", id))
	return nil
}
