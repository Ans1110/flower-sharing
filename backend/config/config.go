package config

import (
	"flower-backend/utils"
	"strings"
	"time"
)

type Config struct {
	Port                 string
	WhiteListedOrigins   []string
	DatabaseURL          string
	GO_ENV               string
	JWTSecret            string
	JWTRefreshSecret     string
	JWTExpiry            time.Duration
	JWTRefreshExpiry     time.Duration
	DefaultResLimit      int
	DefaultResOffset     int
	CloudinaryCloudName  string
	CloudinaryAPIKey     string
	CloudinaryAPISecret  string
	WhiteListAdminEmails []string
	AllowOrigins         []string
}

func LoadConfig() *Config {
	port := utils.GetEnv("PORT", "8080")
	goEnv := utils.GetEnv("GO_ENV", "development")
	dbURL := utils.MustGetEnv("DB_URL")
	whiteListedOrigins := []string{}
	jwtSecret := utils.MustGetEnv("JWT_SECRET")
	jwtRefreshSecret := utils.MustGetEnv("JWT_REFRESH_SECRET")
	jwtExpiry := utils.ParseDuration(utils.GetEnv("JWT_EXPIRY", "1h"))
	jwtRefreshExpiry := utils.ParseDuration(utils.GetEnv("JWT_REFRESH_EXPIRY", "720h"))
	defaultResLimit := 20
	defaultResOffset := 0
	cloudinaryCloudName := utils.MustGetEnv("CLOUDINARY_CLOUD_NAME")
	cloudinaryAPIKey := utils.MustGetEnv("CLOUDINARY_API_KEY")
	cloudinaryAPISecret := utils.MustGetEnv("CLOUDINARY_API_SECRET")
	whiteListAdminEmails := strings.Split(utils.MustGetEnv("WHITE_LIST_ADMIN_EMAILS"), ",")
	// AllowOrigins
	allowOriginsRaw := utils.MustGetEnv("ALLOW_ORIGINS")
	allowOriginsRaw = strings.Trim(allowOriginsRaw, "[]\"")
	allowOrigins := strings.Split(allowOriginsRaw, ",")
	return &Config{
		Port:                 port,
		WhiteListedOrigins:   whiteListedOrigins,
		DatabaseURL:          dbURL,
		GO_ENV:               goEnv,
		JWTSecret:            jwtSecret,
		JWTRefreshSecret:     jwtRefreshSecret,
		JWTExpiry:            jwtExpiry,
		JWTRefreshExpiry:     jwtRefreshExpiry,
		DefaultResLimit:      defaultResLimit,
		DefaultResOffset:     defaultResOffset,
		CloudinaryCloudName:  cloudinaryCloudName,
		CloudinaryAPIKey:     cloudinaryAPIKey,
		CloudinaryAPISecret:  cloudinaryAPISecret,
		WhiteListAdminEmails: whiteListAdminEmails,
		AllowOrigins:         allowOrigins,
	}
}
