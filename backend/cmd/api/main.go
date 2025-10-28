package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Yusufdot101/note-nest/internal/user"
	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
)

const PORT = ":8080"

func main() {
	dsn := os.Getenv("DSN")
	router := httprouter.New()
	DB, err := openDB(dsn)
	if err != nil {
		log.Fatal(err)
	}
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
