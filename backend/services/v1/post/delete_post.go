package post_services

import (
	"go.uber.org/zap"
)

// DeletePostByID
func (s *postService) DeletePostByID(postID, userID uint) error {
	if err := s.repo.DeleteByID(postID, userID); err != nil {
		s.logger.Error("failed to delete post", zap.Error(err))
		return err
	}
	s.logger.Info("post deleted successfully", zap.Uint("postID", postID), zap.Uint("userID", userID))
	return nil
}
