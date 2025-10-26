package main

import (
	"fmt"
	"net/http"
)

func handleMain(w http.ResponseWriter, r *http.Request) {
	msg := "helllo mate"
	w.Write([]byte(msg))
}

func main() {
	srv := http.NewServeMux()
	srv.HandleFunc("/home", handleMain)
	fmt.Println("server listening on port :8080")
	http.ListenAndServe(":8080", srv)
}
