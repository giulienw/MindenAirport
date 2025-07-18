// Package initializers provides initialization functions for the MindenAirport application.
// This includes loading environment variables and other startup configuration.
package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

// LoadEnvs loads environment variables from a .env file in the current directory.
// This function is typically called during application startup to load configuration
// such as database connection strings, JWT secrets, and other environment-specific settings.
//
// The function gracefully handles the case where no .env file exists (useful for production
// environments where environment variables are set directly by the system).
//
// Expected .env file format:
//
//	CONNECTIONSTRING=user/password@host:port/service_name
//	JWT_SECRET=your-secret-key
//	PORT=8080
//
// If the .env file cannot be loaded, a warning is logged but the application continues
// to run, allowing it to use system environment variables instead.
func LoadEnvs() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found or could not be loaded:", err)
	}
}
