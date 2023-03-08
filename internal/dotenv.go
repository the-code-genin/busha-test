package internal

import (
	"os"

	"github.com/joho/godotenv"
)

// Attempts to load .env variables into memory if .env file exists
func LoadEnvVariables(ctx *AppContext) error {
	// Set default config container
	ctx.SetConfig(&Config{map[string]interface{}{}})

	// Skip operation if .env file does not exist
	if _, err := os.Stat(".env"); os.IsNotExist(err) {
		return nil
	}

	// Load env variables into memory
	return godotenv.Load(".env")
}
