package utils

import (
	"crypto/sha1"
	"encoding/base64"
)

func GenerateShortCode(url string) string {
	hash := sha1.Sum([]byte(url))

	return base64.URLEncoding.EncodeToString(hash[:])[:8]
}