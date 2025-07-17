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
		"http://localhost:3000",     // Frontend in Docker
		"http://127.0.0.1:3000",     // Frontend alternative
		"http://localhost:5173",     // Frontend dev server
		"http://127.0.0.1:5173",     // Frontend dev server alternative
		"http://localhost:8080",     // Backend self
		"http://127.0.0.1:8080",     // Backend self alternative
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
