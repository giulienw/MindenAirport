package models

import "time"

// LoginRequest represents the request body for login
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

// RegisterRequest represents the request body for registration
type RegisterRequest struct {
	FirstName string    `json:"firstName" binding:"required"`
	LastName  string    `json:"lastName" binding:"required"`
	Email     string    `json:"email" binding:"required,email"`
	Password  string    `json:"password" binding:"required,min=6"`
	Phone     string    `json:"phone,omitempty"`
	Birthdate time.Time `json:"birthdate" binding:"required"`
}

// AuthResponse represents the response for authentication endpoints
type AuthResponse struct {
	Token     string       `json:"token"`
	ExpiresAt time.Time    `json:"expiresAt"`
	User      UserResponse `json:"user"`
}

// UserResponse represents a sanitized user response (without password)
type UserResponse struct {
	ID        string    `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone,omitempty"`
	Birthdate time.Time `json:"birthdate"`
	Active    bool      `json:"active"`
	Role      string    `json:"role"`
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
