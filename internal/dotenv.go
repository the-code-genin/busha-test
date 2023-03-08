package internal

import (
	"os"

	"github.com/joho/godotenv"
)

// Attempts to load .env variables into memory if .env file exists
func LoadEnvVariables() error {
	// Skip operation if .env file does not exist
	if _, err := os.Stat(".env"); os.IsNotExist(err) {
		return nil
	}
	return godotenv.Load(".env")
}
