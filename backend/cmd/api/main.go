package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

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
	return DB, nil
}
