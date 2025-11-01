package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Yusufdot101/note-nest/internal/custom_errors"
	"github.com/Yusufdot101/note-nest/internal/user"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
)

const PORT = ":8080"

var dsn string

func init() {
	if os.Getenv("APP_ENV") != "docker" {
		if err := godotenv.Load(); err != nil {
			log.Panicf("could not load .env file: %v", err)
		}
	}

	dsn = fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
		os.Getenv("SSL_MODE"),
	)
}

func main() {
	DB, err := openDB(dsn)
	if err != nil {
		log.Fatal(err)
	}

	router := httprouter.New()
	router.NotFound = http.HandlerFunc(custom_errors.NotFoundErrorResponse)
	router.MethodNotAllowed = http.HandlerFunc(custom_errors.MethodNotAllowedErrorResponse)
	user.RegisterRoutes(router, DB)

	log.Printf("server running on %s\n", PORT)
	err = http.ListenAndServe(PORT, router)
	if err != nil {
		log.Fatal(err)
	}
}

func openDB(dsn string) (*sql.DB, error) {
	DB, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	// test the connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = DB.PingContext(ctx)
	if err != nil {
		return nil, err
	}
	return DB, nil
}
