package main

import (
	"fmt"
	"log"
	"net/http"
)

func addTwoNumbers(a, b int) int {
	return a + b
}

func handleMain(w http.ResponseWriter, r *http.Request) {
	a, b := 1, 2
	msg := fmt.Sprintf("helllo mate, %d + %d = %d\n", a, b, addTwoNumbers(a, b))
	_, err := w.Write([]byte(msg))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func main() {
	srv := http.NewServeMux()
	srv.HandleFunc("/home", handleMain)
	fmt.Println("server listening on port :8080")
	err := http.ListenAndServe(":8080", srv)
	if err != nil {
		log.Fatal(err)
		return
	}
}
