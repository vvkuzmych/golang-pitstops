//package main
//
//import (
//	"crypto/tls"
//	"fmt"
//	"log"
//	"net/http"
//)
//
//// Handler for the root endpoint
//func rootHandler(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprintf(w, "Welcome to the HTTP/2 server in Go!")
//}
//
//// Handler for the /hello endpoint
//func helloHandler(w http.ResponseWriter, r *http.Request) {
//	name := r.URL.Query().Get("name")
//	if name == "" {
//		name = "World"
//	}
//	fmt.Fprintf(w, "Hello, %s!", name)
//}
//
//// Handler for the /health endpoint
//func healthHandler(w http.ResponseWriter, r *http.Request) {
//	w.WriteHeader(http.StatusOK)
//	fmt.Fprintf(w, "OK")
//}
//
//func main() {
//	// Register HTTP handlers
//	http.HandleFunc("/", rootHandler)
//	http.HandleFunc("/hello", helloHandler)
//	http.HandleFunc("/health", healthHandler)
//
//	// Create TLS configuration for HTTP/2
//	tlsConfig := &tls.Config{
//		MinVersion: tls.VersionTLS12,
//	}
//
//	// Define the server properties
//	server := &http.Server{
//		Addr:      ":8443", // Listen on port 8443
//		Handler:   nil,     // Use default mux
//		TLSConfig: tlsConfig,
//	}
//
//	log.Println("Starting HTTP/2 server on https://localhost:8443")
//
//	// Start the server with TLS
//	if err := server.ListenAndServeTLS("server.crt", "server.key"); err != nil {
//		log.Fatalf("Server failed to start: %v", err)
//	}
//}

package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
)

// Handler for serving the root HTML page
func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	htmlContent := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>HTTP/2 Example</title>
	</head>
	<body>
		<h1>Welcome to the HTTP/2 Server</h1>
		<p>Visit <a href="/api/hello">/api/hello</a> for a JSON response.</p>
	</body>
	</html>`
	fmt.Fprint(w, htmlContent)
}

// Handler for the /api/hello endpoint
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Ensure we are using HTTP/2
	if r.ProtoMajor != 2 {
		http.Error(w, "Only HTTP/2 is supported", http.StatusHTTPVersionNotSupported)
		return
	}

	// Set response headers and content
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := `{"message": "Hello, HTTP/2!"}`
	_, _ = w.Write([]byte(response))
}

func main() {
	// Register HTTP handlers
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/api/hello", helloHandler)

	// Create TLS configuration for HTTP/2
	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
	}

	// Define the server properties
	server := &http.Server{
		Addr:      ":8443", // Listen on port 8443
		Handler:   nil,     // Use default mux
		TLSConfig: tlsConfig,
	}

	log.Println("Starting HTTP/2 server on https://localhost:8443")

	// Start the server with TLS
	if err := server.ListenAndServeTLS("server.crt", "server.key"); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
