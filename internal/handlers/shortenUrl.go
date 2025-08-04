package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"urlshortener/internal/storage"
	"urlshortener/internal/utils"
)

type ShortenRequest struct {
	URL string `json:"url"`
}

type ShortenResponse struct {
	ShortCode string `json:"short_code"`
}

func ShortenHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Read the body from the request and confirm err
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var req ShortenRequest
	if err := json.Unmarshal(body, &req); err != nil{
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.URL == "" {
		http.Error(w, "URL is required", http.StatusBadRequest)
		return
	}

	shortCode := utils.GenerateShortCode(req.URL)

	storage.SaveURL(shortCode, req.URL)

	response := ShortenResponse{ShortCode: shortCode}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}