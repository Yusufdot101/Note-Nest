package main

import (
	"log"
	"net/http"

	"github.com/Yusufdot101/note-nest/internal/user"
	"github.com/julienschmidt/httprouter"
)

const PORT = ":8080"

func main() {
	router := httprouter.New()
	user.RegisterRoutes(router)

	log.Printf("server running on %s\n", PORT)
	err := http.ListenAndServe(PORT, router)
	if err != nil {
		log.Fatal(err)
	}
}
