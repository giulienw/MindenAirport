// Package middleware provides HTTP middleware functions for the MindenAirport API.
// This includes authentication, CORS, logging, and other cross-cutting concerns.
package middleware

import (
	"net/http"
	"strings"

	"mindenairport/utils"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware validates JWT tokens for protected routes.
// This middleware checks for a valid Bearer token in the Authorization header
// and validates it using the JWT utility functions.
//
// Expected header format: "Authorization: Bearer <jwt-token>"
//
// If validation succeeds, the middleware sets the following in the Gin context:
//   - "userID": The authenticated user's ID
//   - "email": The authenticated user's email
//
// Returns HTTP 401 Unauthorized if:
//   - No Authorization header is present
//   - Header doesn't start with "Bearer "
//   - Token is invalid, expired, or malformed
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}

		// Check for Bearer token format
		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header format"})
			c.Abort()
			return
		}

		// Extract token
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// Validate token
		claims, err := utils.ValidateJWT(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		// Set user information in context
		c.Set("userID", claims.UserID)
		c.Set("email", claims.Email)

		c.Next()
	}
}

// OptionalAuthMiddleware validates JWT tokens but doesn't require them
func OptionalAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.Next()
			return
		}

		if strings.HasPrefix(authHeader, "Bearer ") {
			tokenString := strings.TrimPrefix(authHeader, "Bearer ")
			claims, err := utils.ValidateJWT(tokenString)
			if err == nil {
				c.Set("userID", claims.UserID)
				c.Set("email", claims.Email)
			}
		}

		c.Next()
	}
}
