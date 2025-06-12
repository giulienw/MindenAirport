package routers

import (
	"net/http"
	"strings"

	"mindenairport/database"
	"mindenairport/models"
	"mindenairport/utils"

	"github.com/gin-gonic/gin"
)

// Register handles user registration
func Register(db database.Database) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req models.RegisterRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data", "details": err.Error()})
			return
		}

		// Normalize email to lowercase
		req.Email = strings.ToLower(strings.TrimSpace(req.Email))

		// Check if email already exists
		exists, err := db.CheckEmailExists(req.Email)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
			return
		}

		if exists {
			c.JSON(http.StatusConflict, gin.H{"error": "Email already registered"})
			return
		}

		// Create user
		user, err := db.CreateUser(req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
			return
		}

		// Generate JWT token
		token, expiresAt, err := utils.GenerateJWT(user.ID, user.Email)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
			return
		}

		// Return success response
		response := models.AuthResponse{
			Token:     token,
			ExpiresAt: expiresAt,
			User:      user.ToUserResponse(),
		}

		c.JSON(http.StatusCreated, gin.H{
			"message": "User registered successfully",
			"data":    response,
		})
	}
}

// Login handles user authentication
func Login(db database.Database) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req models.LoginRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data", "details": err.Error()})
			return
		}

		// Normalize email to lowercase
		req.Email = strings.ToLower(strings.TrimSpace(req.Email))

		// Get user by email
		user, err := db.GetUserByEmail(req.Email)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
			return
		}

		if user == nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
			return
		}

		// Check if user is active
		if !user.Active {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Account is deactivated"})
			return
		}

		// Verify password
		if !utils.CheckPasswordHash(req.Password, user.Password) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
			return
		}

		// Update last login time
		if err := db.UpdateUserLastLogin(user.ID); err != nil {
			// Log error but don't fail the login
			// log.Printf("Failed to update last login for user %s: %v", user.ID, err)
		}

		// Generate JWT token
		token, expiresAt, err := utils.GenerateJWT(user.ID, user.Email)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
			return
		}

		// Return success response
		response := models.AuthResponse{
			Token:     token,
			ExpiresAt: expiresAt,
			User:      user.ToUserResponse(),
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Login successful",
			"data":    response,
		})
	}
}

// GetProfile returns the authenticated user's profile
func GetProfile(db database.Database) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get user ID from context (set by auth middleware)
		userID, exists := c.Get("userID")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		// Get user from database
		user, err := db.GetUserByID(userID.(string))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
			return
		}

		if user == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": user.ToUserResponse(),
		})
	}
}

// Logout handles user logout (mainly for client-side token removal)
func Logout() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Logout successful",
		})
	}
}

// RefreshToken generates a new token for authenticated users
func RefreshToken(db database.Database) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get user ID from context (set by auth middleware)
		userID, exists := c.Get("userID")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		email, exists := c.Get("email")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		// Generate new JWT token
		token, expiresAt, err := utils.GenerateJWT(userID.(string), email.(string))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": gin.H{
				"token":     token,
				"expiresAt": expiresAt,
			},
		})
	}
}

// GetDashboard returns dashboard data for the authenticated user
func GetDashboard(db database.Database) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get user ID from context (set by auth middleware)
		userID, exists := c.Get("userID")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		// Get user profile
		user, err := db.GetUserByID(userID.(string))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user profile"})
			return
		}

		if user == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}

		// Get user's tickets
		tickets, err := db.GetTicketsByUserID(userID.(string))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get tickets"})
			return
		}

		// Initialize tickets slice if nil
		if tickets == nil {
			tickets = []models.Ticket{}
		}

		// Calculate statistics
		totalTickets := len(tickets)
		activeTickets := 0
		totalSpent := 0.0

		for _, ticket := range tickets {
			if ticket.Status == "CONFIRMED" || ticket.Status == "CHECKED_IN" {
				activeTickets++
			}
			totalSpent += ticket.Price
		}

		// Prepare dashboard response
		dashboardData := gin.H{
			"user": user.ToUserResponse(),
			"statistics": gin.H{
				"totalTickets":  totalTickets,
				"activeTickets": activeTickets,
				"totalSpent":    totalSpent,
			},
			"recentTickets": tickets, // All tickets, you might want to limit this to recent ones
		}

		c.JSON(http.StatusOK, gin.H{
			"data":    dashboardData,
			"message": "Dashboard data retrieved successfully",
		})
	}
}

// AuthRoutes sets up authentication routes
func AuthRoutes(router *gin.RouterGroup, db database.Database) {
	// Public authentication routes
	router.POST("/register", Register(db))
	router.POST("/login", Login(db))
	router.POST("/logout", Logout())
}
