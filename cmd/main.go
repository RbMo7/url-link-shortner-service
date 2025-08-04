package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"urlshortener/internal/handlers"
	"urlshortener/internal/storage"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../.env")

	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	port := os.Getenv("PORT")
	baseURL := os.Getenv("BASE_URL")
	dbHost := os.Getenv("DB_HOST")


	// Get relative path for assets mapping
	// cwd, err := os.Getwd()
	// if err != nil {
	// 	log.Fatal("Failed to get current working directory:", err)
	// }
	// log.Println("Current working directory is:", cwd)

	fmt.Printf("Server will run on port: %s\n", port)
	fmt.Printf("Base URL is: %s\n", baseURL)
	fmt.Printf("Database host is: %s\n", dbHost)

	err = storage.Init()

	if err != nil {
		log.Fatalf("Failed to initialize storage: %v", err)
	}

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "Welcome to the URL Shortener!")
	// 	fmt.Printf("Received request: %s %s\n", r.Method, r.URL.Path)
	// })

	http.HandleFunc("/shorten", handlers.ShortenHandler)
	http.HandleFunc("/", handlers.RedirectURLHandler)

	http.ListenAndServe(":"+port, nil)
}
