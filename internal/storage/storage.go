package storage

import (
	"encoding/json"
	"errors"
	"os"
	"sync"
)

const filePath = "../../assets/urlStorage.json"

var (
	mu       sync.RWMutex
	urlStore map[string]string
)

func Init() error {
	mu.Lock()
	defer mu.Unlock()

	urlStore = make(map[string]string)

	file, err := os.Open(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil // no file yet, skip error
		}
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&urlStore); err != nil {
		return err
	}
	return nil
}

// Save a shortCode → longURL pair
func SaveURL(code, longURL string) error {
	mu.Lock()
	defer mu.Unlock()

	urlStore[code] = longURL
	return writeToFile()
}

// Retrieve longURL for a shortCode
func GetURL(code string) (string, error) {
	mu.RLock()
	defer mu.RUnlock()

	longURL, exists := urlStore[code]
	if !exists {
		return "", errors.New("URL not found")
	}
	return longURL, nil
}

// Write the entire urlStore map to file
func writeToFile() error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(urlStore)
}
