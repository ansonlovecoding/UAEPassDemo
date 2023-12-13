package utils

import (
	"math/rand"
	"time"
)

// generate random string
func generateRandomString(length int) string {
	time.Now().Nanosecond()
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	result := make([]byte, length)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}

	return string(result)
}
