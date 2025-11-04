package app

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/julienschmidt/httprouter"
)

const PORT = ":8080"

type config struct {
	Port           string
	Handler        http.Handler
	TrustedOrigins []string
	DB             struct {
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

func NewApplication() *Application {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
		os.Getenv("SSL_MODE"),
	)

	cfg := &config{
		Port:           PORT,
		TrustedOrigins: strings.Split(os.Getenv("TRUSTED_ORIGINS"), ","),
		DB: struct {
			DSN                   string
			MaxOpenConnections    int
			MaxIdleConnections    int
			ConnectionMaxIdleTime string
		}{
			DSN:                   dsn,
			MaxOpenConnections:    mustInt(os.Getenv("MAX_OPEN_CONNECTIONS")),
			MaxIdleConnections:    mustInt(os.Getenv("MAX_IDLE_CONNECTIONS")),
			ConnectionMaxIdleTime: os.Getenv("CONNECTION_MAX_IDLE_TIME"),
		},
	}

	DB, err := openDB(cfg)
	if err != nil {
		log.Fatal(err)
	}

	router := httprouter.New()
	handler := ConfigureRouter(router, DB)
	cfg.Handler = handler

	app := &Application{
		Config: *cfg,
		DB:     DB,
	}

	return app
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

func mustInt(value string) int {
	i, err := strconv.Atoi(value)
	if err != nil {
		panic(err)
	}
	return i
}
