package main

import (
	"crypto/tls"
	"fmt"
	"github.com/quic-go/quic-go/http3"
	"log"
	"net/http"
)

// Handler for the root endpoint
func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HTTP/3 server in Go!")
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
	server := &http3.Server{
		Addr: ":8445", // Using port 8445
		TLSConfig: &tls.Config{
			MinVersion:   tls.VersionTLS13,
			Certificates: []tls.Certificate{loadTLSCertificate()},
		},
	}

	log.Println("Starting HTTP/3 server on https://localhost:8445")

	// Start the HTTP/3 server
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

// loadTLSCertificate loads the TLS certificate and key
func loadTLSCertificate() tls.Certificate {
	cert, err := tls.LoadX509KeyPair("server.crt", "server.key")
	if err != nil {
		log.Fatalf("Failed to load TLS certificate: %v", err)
	}
	return cert
}
