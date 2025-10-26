package main

import (
	"fmt"
	"net/http"
)

func addTwoNumbers(a, b int) int {
	return a + b
}

func handleMain(w http.ResponseWriter, r *http.Request) {
	a, b := 1, 2
	msg := fmt.Sprintf("helllo mate, %d + %d = %d\n", a, b, addTwoNumbers(a, b))
	w.Write([]byte(msg))
}

func main() {
	srv := http.NewServeMux()
	srv.HandleFunc("/home", handleMain)
	fmt.Println("server listening on port :8080")
	http.ListenAndServe(":8080", srv)
}
