// Package utils provides utility functions for authentication, password hashing,
// and JWT token management used throughout the MindenAirport application.
package utils

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// jwtSecret holds the secret key used for signing JWT tokens
var jwtSecret = []byte(getJWTSecret())

// Claims represents the JWT claims structure containing user information
// and standard JWT registered claims for token validation.
type Claims struct {
	UserID               string `json:"user_id"` // Unique identifier for the authenticated user
	Email                string `json:"email"`   // User's email address
	jwt.RegisteredClaims        // Standard JWT claims (exp, iat, etc.)
}

// getJWTSecret returns the JWT secret from environment variables.
// Falls back to a default secret for development if JWT_SECRET is not set.
//
// Security Note: In production, always set a strong, unique JWT_SECRET environment variable.
func getJWTSecret() string {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		// Default secret for development - should be changed in production
		return "your-super-secret-jwt-key-change-this-in-production"
	}
	return secret
}

// HashPassword securely hashes a password using bcrypt with default cost.
// Uses bcrypt's adaptive hashing function which is resistant to timing attacks
// and becomes more secure over time as hardware improves.
//
// Parameters:
//   - password: Plain text password to hash
//
// Returns:
//   - string: Bcrypt hash of the password
//   - error: Any error that occurred during hashing
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// CheckPasswordHash compares a plain text password with its bcrypt hash.
// Uses constant-time comparison to prevent timing attacks.
//
// Parameters:
//   - password: Plain text password to verify
//   - hash: Bcrypt hash to compare against
//
// Returns:
//   - bool: true if password matches hash, false otherwise
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// GenerateJWT creates a signed JWT token for user authentication.
// The token includes user ID and email in the claims and expires in 24 hours.
//
// Parameters:
//   - userID: Unique identifier for the user
//   - email: User's email address
//
// Returns:
//   - string: Signed JWT token
//   - time.Time: Token expiration time
//   - error: Any error that occurred during token generation
func GenerateJWT(userID, email string) (string, time.Time, error) {
	expirationTime := time.Now().Add(24 * time.Hour) // Token expires in 24 hours

	claims := &Claims{
		UserID: userID,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   userID,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)

	return tokenString, expirationTime, err
}

// ValidateJWT validates and parses a JWT token, returning the claims if valid.
func ValidateJWT(tokenString string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
