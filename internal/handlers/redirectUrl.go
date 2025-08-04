package handlers

import (
	"encoding/json"
	"net/http"
	"urlshortener/internal/storage"
)

type RedirectURL struct {
	URL string `json:"url"`
}

func RedirectURLHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	code := r.URL.Path[1:]

	shortUrl, err := storage.GetURL(code)

	if err != nil {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}

	response := RedirectURL{URL: shortUrl}
	w.WriteHeader(http.StatusFound)
	json.NewEncoder(w).Encode(response)
}
