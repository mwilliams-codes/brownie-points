package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Users!")
}

func main() {
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/user", userHandler)

	fmt.Println("Starting server on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
