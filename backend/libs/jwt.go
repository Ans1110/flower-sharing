package libs

import (
	"errors"
	"flower-backend/config"
	"flower-backend/log"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.uber.org/zap"
)

var cfg = config.LoadConfig()
var logger = log.InitLog().Sugar()

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
