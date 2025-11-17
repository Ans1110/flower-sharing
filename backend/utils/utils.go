package utils

import (
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
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
	re := regexp.MustCompile(`/^[a-z0-9]{2,}$/i`)
	return re.MatchString(username)
}

func ValidatePassword(password string) bool {
	// Check minimum length
	if len(password) < 8 {
		return false
	}

	// Check for at least one lowercase letter
	hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)
	// Check for at least one uppercase letter
	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
	// Check for at least one digit
	hasDigit := regexp.MustCompile(`[0-9]`).MatchString(password)
	// Check for at least one special character
	hasSpecial := regexp.MustCompile(`[#?!@$%^&*-]`).MatchString(password)

	return hasLower && hasUpper && hasDigit && hasSpecial
}

// formatUnixTime formats Unix timestamp as string
func FormatUnixTime(timestamp int64) string {
	return time.Unix(timestamp, 0).Format(time.RFC3339)
}

// formatOrigin
func FormatOrigins() []string {
	allowOriginsRaw := MustGetEnv("ALLOW_ORIGINS")
	allowOriginsRaw = strings.Trim(allowOriginsRaw, "[]\"")
	allowOrigins := strings.Split(allowOriginsRaw, ",")

	return allowOrigins
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		logger, _ := zap.NewProduction()
		logger.Fatal("failed to hash password", zap.Error(err))
		return "", err
	}
	return string(hashedPassword), nil
}

// ParseUint parses a string to a uint
func ParseUint(s string, logger *zap.SugaredLogger) (uint, error) {
	num, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		logger.Error("failed to parse uint", zap.String("value", s), zap.Error(err))
		return 0, err
	}
	return uint(num), nil
}

// ParseInt parses a string to an int with a default value
func ParseInt(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		logger, _ := zap.NewProduction()
		logger.Fatal("failed to parse int", zap.String("value", s), zap.Error(err))
		return 0
	}
	return num
}
