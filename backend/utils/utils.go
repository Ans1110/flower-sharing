package utils

import (
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func init() {
	// Load .env file from the project root
	// Try multiple possible locations
	wd, _ := os.Getwd()
	envPaths := []string{
		".env",
		filepath.Join(wd, ".env"),
		filepath.Join(wd, "..", ".env"),
		filepath.Join(wd, "backend", ".env"),
	}

	for _, path := range envPaths {
		if err := godotenv.Load(path); err == nil {
			break
		}
	}
}

func GetEnv(key, defaultValue string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return defaultValue
}

func MustGetEnv(key string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}

	logger, _ := zap.NewProduction()
	logger.Fatal("environment variable is not set", zap.String("key", key))
	return ""
}

func ParseDuration(s string) time.Duration {
	if len(s) > 1 && s[len(s)-1] == 'd' {
		days, err := strconv.Atoi(s[:len(s)-1])
		if err == nil {
			return time.Duration(days) * 24 * time.Hour
		}
	}
	dur, err := time.ParseDuration(s)
	if err != nil {
		// Use zap logger directly to avoid import cycle
		logger, _ := zap.NewProduction()
		logger.Fatal("invalid duration", zap.String("duration", s), zap.Error(err))
		return 0
	}
	return dur
}

func ValidateEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}

func ValidateUsername(username string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(username)
}
