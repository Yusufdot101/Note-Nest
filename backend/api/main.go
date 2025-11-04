package main

import (
	"log"

	"github.com/Yusufdot101/note-nest/internal/app"
	"github.com/Yusufdot101/note-nest/internal/utilities"
	_ "github.com/lib/pq"
)

func init() {
	utilities.LoadEnv(".env")
}

func main() {
	a, err := app.NewApplication()
	if err != nil {
		log.Fatal(err)
	}
	err = a.Serve()
	if err != nil {
		log.Fatal(err)
	}
}
