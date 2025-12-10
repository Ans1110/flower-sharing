package libs

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"flower-backend/config"
	"flower-backend/log"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
)

var cfg = config.LoadConfig()
var logger = log.InitLog().Sugar()

// GenerateAccessToken generates an access token with user ID (backward compatible)
func GenerateAccessToken(UserId uint) string {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": UserId,
		"exp": time.Now().Add(cfg.JWTExpiry).Unix(),
	})
	accessTokenString, err := accessToken.SignedString([]byte(cfg.JWTSecret))
	if err != nil {
		logger.Error("failed to generate access token", zap.Error(err))
		return ""
	}
	return accessTokenString
}

// GenerateRefreshToken generates a refresh token with user ID (backward compatible)
func GenerateRefreshToken(UserId uint) string {
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": UserId,
		"exp": time.Now().Add(cfg.JWTRefreshExpiry).Unix(),
	})
	refreshTokenString, err := refreshToken.SignedString([]byte(cfg.JWTRefreshSecret))
	if err != nil {
		logger.Error("failed to generate refresh token", zap.Error(err))
		return ""
	}
	return refreshTokenString
}

// GenerateRandomString generates a random string of specified length
func GenerateRandomString(length int) string {
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)[:length]
}

func VerifyAccessToken(tokenString string) (uint, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(cfg.JWTSecret), nil
	})

	if err != nil {
		return 0, err
	}

	if !token.Valid {
		return 0, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("invalid token claims")
	}

	sub, ok := claims["sub"].(float64)
	if !ok {
		return 0, errors.New("subject claim missing or invalid")
	}
	return uint(sub), nil
}

func VerifyRefreshToken(tokenString string) (uint, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(cfg.JWTRefreshSecret), nil
	})
	if err != nil {
		return 0, err
	}

	if !token.Valid {
		return 0, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("invalid token claims")
	}

	sub, ok := claims["sub"].(float64)
	if !ok {
		return 0, errors.New("subject claim missing or invalid")
	}
	return uint(sub), nil
}
