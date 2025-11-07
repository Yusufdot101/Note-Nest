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

const PORT = ":8080"

type config struct {
	Port    string
	Handler http.Handler
	DB      struct {
		DSN                   string
		MaxOpenConnections    int
		MaxIdleConnections    int
		ConnectionMaxIdleTime string
	}
}

type Application struct {
	Config config
	DB     *sql.DB
}

func NewApplication() (*Application, error) {
	requiredEnvVars := []string{
		"DB_USER", "DB_PASS", "DB_HOST", "DB_PORT", "DB_NAME", "SSL_MODE",
		"TRUSTED_ORIGINS", "MAX_OPEN_CONNECTIONS", "MAX_IDLE_CONNECTIONS",
		"CONNECTION_MAX_IDLE_TIME",
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
	maxOpen, err := parseInt(os.Getenv("MAX_OPEN_CONNECTIONS"))
	if err != nil {
		return nil, err
	}
	maxIdle, err := parseInt(os.Getenv("MAX_IDLE_CONNECTIONS"))
	if err != nil {
		return nil, err
	}
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
	}

	DB, err := openDB(cfg)
	if err != nil {
		return nil, err
	}

	router := httprouter.New()
	handler := ConfigureRouter(router, DB)
	cfg.Handler = handler

	app := &Application{
		Config: *cfg,
		DB:     DB,
	}

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

func parseInt(value string) (int, error) {
	i, err := strconv.Atoi(value)
	if err != nil {
		return -1, err
	}
	return i, nil
}
