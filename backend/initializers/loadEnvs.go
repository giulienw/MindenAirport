package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

// LoadEnvs loads the .env file
func LoadEnvs() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found or could not be loaded:", err)
	}
}
