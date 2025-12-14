package tasks

import (
	user_repository "flower-backend/repositories/v1/user"
	"time"

	"go.uber.org/zap"
)

// StartTokenCleanup launches a background ticker to prune expired refresh tokens.
// It runs every hour and logs the count of deleted tokens.
func StartTokenCleanup(repo user_repository.UserRepository, logger *zap.Logger) {
	go func() {
		ticker := time.NewTicker(1 * time.Hour)
		defer ticker.Stop()

		for range ticker.C {
			now := time.Now()
			if deleted, err := repo.DeleteExpiredTokens(now); err != nil {
				logger.Error("token cleanup failed", zap.Error(err))
			} else if deleted > 0 {
				logger.Info("token cleanup removed expired tokens", zap.Int64("count", deleted))
			}
		}
	}()
}
