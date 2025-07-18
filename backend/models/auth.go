// Package models contains data structures for authentication and user management.
// These models handle user registration, login, and authentication responses.
package models

import "time"

// LoginRequest represents the request body for user login authentication.
// Used when users attempt to sign in to their accounts.
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`    // User's email address (must be valid email format)
	Password string `json:"password" binding:"required,min=6"` // User's password (minimum 6 characters)
}

// RegisterRequest represents the request body for new user registration.
// Contains all required information to create a new user account.
type RegisterRequest struct {
	FirstName string    `json:"firstName" binding:"required"`      // User's first name
	LastName  string    `json:"lastName" binding:"required"`       // User's last name
	Email     string    `json:"email" binding:"required,email"`    // User's email (must be unique and valid)
	Password  string    `json:"password" binding:"required,min=6"` // Password (minimum 6 characters)
	Phone     string    `json:"phone,omitempty"`                   // Optional phone number
	Birthdate time.Time `json:"birthdate" binding:"required"`      // User's date of birth
}

// AuthResponse represents the response for successful authentication operations.
// Returned after successful login or registration with JWT token and user info.
type AuthResponse struct {
	Token     string       `json:"token"`     // JWT token for authenticated requests
	ExpiresAt time.Time    `json:"expiresAt"` // Token expiration timestamp
	User      UserResponse `json:"user"`      // Sanitized user information
}

// UserResponse represents a sanitized user response without sensitive data.
// Used in API responses to avoid exposing password hashes or other sensitive info.
type UserResponse struct {
	ID          string    `json:"id"`        // Unique user identifier
	FirstName   string    `json:"firstName"` // User's first name
	LastName    string    `json:"lastName"`
	Email       string    `json:"email"`
	Phone       string    `json:"phone,omitempty"`
	Birthdate   time.Time `json:"birthdate"`
	Active      bool      `json:"active"`
	Role        string    `json:"role"`
	TicketCount int       `json:"ticketCount,omitempty"` // Optional field for user ticket count
}

// ToUserResponse converts AirportUser to UserResponse (removes password)
func (u *AirportUser) ToUserResponse() UserResponse {
	return UserResponse{
		ID:        u.ID,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		Phone:     u.Phone,
		Birthdate: u.Birthdate,
		Active:    u.Active,
		Role:      u.Role,
	}
}
