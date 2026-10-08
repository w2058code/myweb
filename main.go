package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// Vercel routes traffic to the $PORT environment variable (defaults to 80)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Fallback for local development
	}

	// Define your endpoints
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello from a Go app running inside Docker on Vercel!")
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "OK")
	})

	log.Printf("Server starting on port %s...", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
