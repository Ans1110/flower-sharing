package config

import (
	"flower-backend/utils"
	"strings"
	"time"
)

type Config struct {
	Port                 string
	APIBaseURL           string
	DBHost               string
	DBPort               string
	DBUser               string
	DBPassword           string
	DBName               string
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
	CloudinaryFolder     string
	WhiteListAdminEmails []string
	AllowOrigins         []string
	RequestTimeout       time.Duration
	ReadTimeout          time.Duration
	WriteTimeout         time.Duration
	IdleTimeout          time.Duration
	// Database connection pool settings
	DBMaxOpenConns    int
	DBMaxIdleConns    int
	DBConnMaxLifetime time.Duration
	DBConnMaxIdleTime time.Duration
	// OAuth configuration
	GoogleClientID     string
	GoogleClientSecret string
	GoogleRedirectURL  string
	GithubClientID     string
	GithubClientSecret string
	GithubRedirectURL  string
	FrontendURL        string
}

func LoadConfig() *Config {
	port := utils.GetEnv("PORT", "8080")
	goEnv := utils.GetEnv("GO_ENV", "development")
	apiBaseURL := utils.GetEnv("API_BASE_URL", "http://localhost:8080")
	dbHost := utils.GetEnv("DB_HOST", "localhost")
	dbPort := utils.GetEnv("DB_PORT", "3306")
	dbUser := utils.GetEnv("DB_USER", "root")
	dbPassword := utils.GetEnv("DB_PASSWORD", "root")
	dbName := utils.GetEnv("DB_NAME", "flower_sharing")

	jwtSecret := utils.MustGetEnv("JWT_SECRET")
	jwtRefreshSecret := utils.MustGetEnv("JWT_REFRESH_SECRET")
	jwtExpiry := utils.ParseDuration(utils.GetEnv("JWT_EXPIRY", "1h"))
	jwtRefreshExpiry := utils.ParseDuration(utils.GetEnv("JWT_REFRESH_EXPIRY", "720h"))

	defaultResLimit := 20
	defaultResOffset := 0

	cloudinaryCloudName := utils.MustGetEnv("CLOUDINARY_CLOUD_NAME")
	cloudinaryAPIKey := utils.MustGetEnv("CLOUDINARY_API_KEY")
	cloudinaryAPISecret := utils.MustGetEnv("CLOUDINARY_API_SECRET")
	cloudinaryFolder := utils.MustGetEnv("CLOUDINARY_FOLDER")

	whiteListAdminEmails := strings.Split(utils.MustGetEnv("WHITE_LIST_ADMIN_EMAILS"), ",")

	allowOrigins := strings.Split(utils.MustGetEnv("ALLOW_ORIGINS"), ",")

	// Timeout configurations
	requestTimeout := utils.ParseDuration(utils.GetEnv("REQUEST_TIMEOUT", "30s"))
	readTimeout := utils.ParseDuration(utils.GetEnv("READ_TIMEOUT", "15s"))
	writeTimeout := utils.ParseDuration(utils.GetEnv("WRITE_TIMEOUT", "15s"))
	idleTimeout := utils.ParseDuration(utils.GetEnv("IDLE_TIMEOUT", "60s"))

	// Database connection pool configurations
	dbMaxOpenConns := utils.ParseInt(utils.GetEnv("DB_MAX_OPEN_CONNS", "100"))
	dbMaxIdleConns := utils.ParseInt(utils.GetEnv("DB_MAX_IDLE_CONNS", "10"))
	dbConnMaxLifetime := utils.ParseDuration(utils.GetEnv("DB_CONN_MAX_LIFETIME", "1h"))
	dbConnMaxIdleTime := utils.ParseDuration(utils.GetEnv("DB_CONN_MAX_IDLE_TIME", "10m"))

	// OAuth configurations
	googleClientID := utils.GetEnv("GOOGLE_CLIENT_ID", "")
	googleClientSecret := utils.GetEnv("GOOGLE_CLIENT_SECRET", "")
	googleRedirectURL := utils.GetEnv("GOOGLE_REDIRECT_URL", "http://localhost:8080/api/v1/auth/google/callback")
	githubClientID := utils.GetEnv("GITHUB_CLIENT_ID", "")
	githubClientSecret := utils.GetEnv("GITHUB_CLIENT_SECRET", "")
	githubRedirectURL := utils.GetEnv("GITHUB_REDIRECT_URL", "http://localhost:8080/api/v1/auth/github/callback")
	frontendURL := utils.GetEnv("FRONTEND_URL", "http://localhost:3000")

	return &Config{
		Port:                 port,
		APIBaseURL:           apiBaseURL,
		DBHost:               dbHost,
		DBPort:               dbPort,
		DBUser:               dbUser,
		DBPassword:           dbPassword,
		DBName:               dbName,
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
		CloudinaryFolder:     cloudinaryFolder,
		WhiteListAdminEmails: whiteListAdminEmails,
		AllowOrigins:         allowOrigins,
		RequestTimeout:       requestTimeout,
		ReadTimeout:          readTimeout,
		WriteTimeout:         writeTimeout,
		IdleTimeout:          idleTimeout,
		DBMaxOpenConns:       dbMaxOpenConns,
		DBMaxIdleConns:       dbMaxIdleConns,
		DBConnMaxLifetime:    dbConnMaxLifetime,
		DBConnMaxIdleTime:    dbConnMaxIdleTime,
		GoogleClientID:       googleClientID,
		GoogleClientSecret:   googleClientSecret,
		GoogleRedirectURL:    googleRedirectURL,
		GithubClientID:       githubClientID,
		GithubClientSecret:   githubClientSecret,
		GithubRedirectURL:    githubRedirectURL,
		FrontendURL:          frontendURL,
	}
}
