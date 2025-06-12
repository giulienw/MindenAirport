package middleware

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// CORSMiddleware returns a configured CORS middleware
func CORSMiddleware() gin.HandlerFunc {
	// Get allowed origins from environment variable if set
	allowedOrigins := []string{
		"http://127.0.0.1:8080", // Allow all origins in development
	}

	config := cors.Config{
		AllowMethods: []string{
			"GET",
			"POST",
			"PUT",
			"PATCH",
			"DELETE",
			"HEAD",
			"OPTIONS",
		},
		AllowHeaders: []string{
			"Origin", "Authorization", "Content-Type",
		},
		ExposeHeaders: []string{
			"Content-Length",
			"Content-Type",
			"Access-Control-Allow-Origin",
		},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}

	// In development mode, allow all origins
	config.AllowOrigins = allowedOrigins

	return cors.New(config)
}
