package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

// LoadEnvs loads the .env file
func LoadEnvs() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
