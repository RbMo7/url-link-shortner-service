package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	port := os.Getenv("PORT")
	baseURL := os.Getenv("BASE_URL")
	dbHost := os.Getenv("DB_HOST")

	fmt.Printf("Server will run on port: %s\n", port)
	fmt.Printf("Base URL is: %s\n", baseURL)
	fmt.Printf("Database host is: %s\n", dbHost)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to the URL Shortener!")
		fmt.Printf("Received request: %s %s\n", r.Method, r.URL.Path)
	})

	http.ListenAndServe(":"+port, nil)
}
