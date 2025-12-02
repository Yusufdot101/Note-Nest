/*
Package app provides configuration and setting up of the api server
*/
package app

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
)

// PORT is the api addr
const PORT = ":8080" // 8080 is the default addr

type config struct {
	Port    string
	Handler http.Handler
	DB      struct {
		DSN                   string
		MaxOpenConnections    int
		MaxIdleConnections    int
		ConnectionMaxIdleTime string
	}
	Limiter struct {
		Enabled bool
		Burst   int
		Rate    float64
	}
	SMTP struct {
		Host     string
		Port     int
		Sender   string
		Username string
		Password string
	}
}

// Application is the api application sturct that has the config and the Database
type Application struct {
	Config config
	DB     *sql.DB
}

// NewApplication returns a new Application strcut with configured options read from .env files
func NewApplication() (*Application, error) {
	requiredEnvVars := []string{
		"DB_USER", "DB_PASS", "DB_HOST", "DB_PORT", "DB_NAME", "SSL_MODE",
		"TRUSTED_ORIGINS", "MAX_OPEN_CONNECTIONS", "MAX_IDLE_CONNECTIONS",
		"CONNECTION_MAX_IDLE_TIME", "RATE_LIMIT_BURST", "RATE_LIMIT_RATE",
		"SMTP_HOST", "SMTP_PORT", "SMTP_SENDER",
		"RESET_PASSWORD_TOKEN_EXPIRATION_TIME", "ACCESS_TOKEN_EXPIRATION_TIME", "FRONTEND_BASE_URL",
		"REFRESH_TOKEN_EXPIRATION_TIME",
		"GOOGLE_CLIENT_ID", "GOOGLE_CLIENT_SECRET", "GOOGLE_REDIRECT_URL",
	}

	emailProvider := os.Getenv("EMAIL_PROVIDER")
	if emailProvider == "mailtrap" {
		requiredEnvVars = append(requiredEnvVars, "SMTP_USERNAME", "SMTP_PASSWORD")
	}

	for _, envVar := range requiredEnvVars {
		if os.Getenv(envVar) == "" {
			return nil, fmt.Errorf("required environment variable %s is not set", envVar)
		}
	}

	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
		os.Getenv("SSL_MODE"),
	)

	maxOpen, err := strconv.Atoi(os.Getenv("MAX_OPEN_CONNECTIONS"))
	if err != nil {
		return nil, err
	}

	maxIdle, err := strconv.Atoi(os.Getenv("MAX_IDLE_CONNECTIONS"))
	if err != nil {
		return nil, err
	}

	rateLimiterBurst, err := strconv.Atoi(os.Getenv("RATE_LIMIT_BURST"))
	if err != nil {
		return nil, err
	}

	rateLimiterRate, err := strconv.ParseFloat(os.Getenv("RATE_LIMIT_RATE"), 64)
	if err != nil {
		return nil, err
	}

	rateLimiterEnabled := os.Getenv("RATE_LIMIT_ENABLED") != "false" // default true

	host := os.Getenv("SMTP_HOST")
	port, err := strconv.Atoi(os.Getenv("SMTP_PORT"))
	if err != nil {
		return nil, err
	}
	sender := os.Getenv("SMTP_SENDER")
	username := os.Getenv("SMTP_USERNAME")
	password := os.Getenv("SMTP_PASSWORD")

	cfg := &config{
		Port: PORT,
		DB: struct {
			DSN                   string
			MaxOpenConnections    int
			MaxIdleConnections    int
			ConnectionMaxIdleTime string
		}{
			DSN:                   dsn,
			MaxOpenConnections:    maxOpen,
			MaxIdleConnections:    maxIdle,
			ConnectionMaxIdleTime: os.Getenv("CONNECTION_MAX_IDLE_TIME"),
		},
		Limiter: struct {
			Enabled bool
			Burst   int
			Rate    float64
		}{
			Enabled: rateLimiterEnabled,
			Burst:   rateLimiterBurst,
			Rate:    rateLimiterRate,
		},
		SMTP: struct {
			Host     string
			Port     int
			Sender   string
			Username string
			Password string
		}{
			Host:     host,
			Port:     port,
			Sender:   sender,
			Username: username,
			Password: password,
		},
	}

	DB, err := openDB(cfg)
	if err != nil {
		return nil, err
	}

	app := &Application{
		DB: DB,
	}
	router := httprouter.New()
	handler := configureRouter(router, cfg, app.DB)
	cfg.Handler = handler
	app.Config = *cfg

	return app, nil
}

func openDB(cfg *config) (*sql.DB, error) {
	DB, err := sql.Open("postgres", cfg.DB.DSN)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// test the connection
	err = DB.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	// set limits on the database pool
	DB.SetMaxOpenConns(cfg.DB.MaxOpenConnections)
	DB.SetMaxIdleConns(cfg.DB.MaxIdleConnections)

	duration, err := time.ParseDuration(cfg.DB.ConnectionMaxIdleTime)
	if err != nil {
		return nil, err
	}

	DB.SetConnMaxIdleTime(duration)
	return DB, nil
}
