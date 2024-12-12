package main

import (
	"fmt"
	"log"
	"net/http"
)

// Handler for the root endpoint
func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HTTP/1.1 server in Go!")
}

// Handler for the /hello endpoint
func helloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}
	fmt.Fprintf(w, "Hello, %s!", name)
}

// Handler for the /health endpoint
func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "OK")
}

func main() {
	// Register HTTP handlers
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/health", healthHandler)

	// Define the server properties
	server := &http.Server{
		Addr:    ":8082", // Listen on port 8080
		Handler: nil,     // Use default mux
	}

	log.Println("Starting HTTP/1.1 server on http://localhost:8080")

	// Start the server
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
