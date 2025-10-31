package main

import (
	"context"
	"database/sql"
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
		dsn = os.Getenv("LOCAL_DB_URL")
	} else {
		dsn = os.Getenv("DOCKER_DB_URL")
	}

	if dsn == "" {
		log.Panic("DB_URL is not set")
	}
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
