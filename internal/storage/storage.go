package storage

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"sync"
)

const filePath = "../assets/urlStorage.json"

var (
	mu       sync.RWMutex
	urlStore map[string]string
)

func Init() error {
	mu.Lock()
	defer mu.Unlock()

	log.Println("Initializing storage...")

	urlStore = make(map[string]string)

	file, err := os.Open(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		log.Printf("Error opening storage file: %v\n", err)
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&urlStore); err != nil {
		log.Printf("Error decoding storage file: %v\n", err)
		return err
	}

	return nil
}

func SaveURL(code, longURL string) error {
	mu.Lock()
	defer mu.Unlock()

	urlStore[code] = longURL

	return writeToFile()
}

func GetURL(code string) (string, error) {
	mu.RLock()
	defer mu.RUnlock()

	longURL, exists := urlStore[code]

	if !exists {
		return "", errors.New("URL not found")
	}

	return longURL, nil
}

func writeToFile() error {
	file, err := os.Create(filePath)
	if err != nil {
		log.Printf("Failed to create file: %v\n", err)
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(urlStore); err != nil {
		log.Printf("Error encoding data to file: %v\n", err)
		return err
	}

	return nil
}
