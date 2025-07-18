// Package middleware provides HTTP middleware functions for cross-cutting concerns.
package middleware

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// CORSMiddleware returns a configured CORS middleware that handles Cross-Origin Resource Sharing.
// This allows the frontend application to make requests to the API from different origins.
//
// Configuration includes:
//   - Allowed origins: localhost variations for development
//   - Allowed methods: All standard HTTP methods
//   - Allowed headers: Common headers including Authorization
//   - Credentials: Enabled for authenticated requests
//   - Max age: 12 hours for preflight cache
//
// The middleware is configured for development with specific localhost origins.
// In production, this should be updated to include only the actual frontend domain.
func CORSMiddleware() gin.HandlerFunc {
	// Get allowed origins from environment variable if set
	allowedOrigins := []string{
		"http://localhost:3000", // Frontend in Docker
		"http://127.0.0.1:3000", // Frontend alternative
		"http://localhost:5173", // Frontend dev server
		"http://127.0.0.1:5173", // Frontend dev server alternative
		"http://localhost:8080", // Backend self
		"http://127.0.0.1:8080", // Backend self alternative
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
