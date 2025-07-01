package routers

import (
	"net/http"
	"strconv"

	"mindenairport/database"
	"mindenairport/models"

	"github.com/gin-gonic/gin"
)

// checkAdminRole validates that the user has admin role
func checkAdminRole(c *gin.Context, db database.Database) (*models.AirportUser, bool) {
	// Get user ID from context (set by AuthMiddleware)
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return nil, false
	}

	// Get user from database to check role
	user, err := db.GetUserByID(userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return nil, false
	}

	if user == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return nil, false
	}

	// Check if user has admin role
	if user.Role != "ADMIN" && user.Role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Admin access required"})
		return nil, false
	}

	return user, true
}

// GetAdminDashboard returns admin dashboard data
func GetAdminDashboard(db database.Database) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check admin role
		_, authorized := checkAdminRole(c, db)
		if !authorized {
			return
		}

		// Get all flights
		flights := db.GetFlights()

		// Get all airports
		airports := db.GetAirports()

		// Get all airlines
		airlines := db.GetAirlines()

		users := db.GetUserCount()

		// Calculate statistics
		totalFlights := len(flights)
		totalAirports := len(airports)
		totalAirlines := len(airlines)
		totalPassengers := users

		// Count active airlines
		activeAirlines := 0
		for _, airline := range airlines {
			if airline.Active {
				activeAirlines++
			}
		}

		// Prepare dashboard response
		dashboardData := gin.H{
			"stats": gin.H{
				"totalFlights":    totalFlights,
				"totalAirports":   totalAirports,
				"totalAirlines":   totalAirlines,
				"activeAirlines":  activeAirlines,
				"totalPassengers": totalPassengers,
			},
			"recentFlights": flights, // You might want to limit this
			"airports":      airports,
			"airlines":      airlines,
		}

		c.JSON(http.StatusOK, gin.H{
			"data":    dashboardData,
			"message": "Admin dashboard data retrieved successfully",
		})
	}
}

// GetAllUsers returns all users for admin
func GetAllUsers(db database.Database) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check admin role
		_, authorized := checkAdminRole(c, db)
		if !authorized {
			return
		}

		// Get pagination parameters
		page := 1
		limit := 50

		if p := c.Query("page"); p != "" {
			if parsed, err := strconv.Atoi(p); err == nil && parsed > 0 {
				page = parsed
			}
		}

		if l := c.Query("limit"); l != "" {
			if parsed, err := strconv.Atoi(l); err == nil && parsed > 0 && parsed <= 100 {
				limit = parsed
			}
		}

		// Get users (you'll need to implement this in database)
		users, total, err := db.GetAllUsers(page, limit)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve users"})
			return
		}

		// Convert to safe user responses
		var userResponses []models.UserResponse
		for _, user := range users {
			userResponses = append(userResponses, user.ToUserResponse())
		}

		c.JSON(http.StatusOK, gin.H{
			"data": gin.H{
				"users": userResponses,
				"pagination": gin.H{
					"page":  page,
					"limit": limit,
					"total": total,
				},
			},
			"message": "Users retrieved successfully",
		})
	}
}

// GetUserById returns a specific user for admin
func GetUserById(db database.Database) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check admin role
		_, authorized := checkAdminRole(c, db)
		if !authorized {
			return
		}

		userID := c.Param("id")

		user, err := db.GetUserByID(userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
			return
		}

		if user == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data":    user.ToUserResponse(),
			"message": "User retrieved successfully",
		})
	}
}

// UpdateUser allows admin to update user information
func UpdateUser(db database.Database) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check admin role
		_, authorized := checkAdminRole(c, db)
		if !authorized {
			return
		}

		userID := c.Param("id")

		var updateData struct {
			FirstName string `json:"firstName"`
			LastName  string `json:"lastName"`
			Email     string `json:"email"`
			Phone     string `json:"phone"`
			Active    *bool  `json:"active"`
			Role      string `json:"role"`
		}

		if err := c.ShouldBindJSON(&updateData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data", "details": err.Error()})
			return
		}

		// Update user (you'll need to implement this in database)
		err := db.UpdateUserByAdmin(userID, updateData.FirstName, updateData.LastName, updateData.Email, updateData.Phone, updateData.Active, updateData.Role)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "User updated successfully",
		})
	}
}

// DeactivateUser allows admin to deactivate a user
func DeactivateUser(db database.Database) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check admin role
		_, authorized := checkAdminRole(c, db)
		if !authorized {
			return
		}

		userID := c.Param("id")

		err := db.DeactivateUser(userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to deactivate user"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "User deactivated successfully",
		})
	}
}

// GetAllTickets returns all tickets for admin
func GetAllTickets(db database.Database) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check admin role
		_, authorized := checkAdminRole(c, db)
		if !authorized {
			return
		}

		// Get pagination parameters
		page := 1
		limit := 50

		if p := c.Query("page"); p != "" {
			if parsed, err := strconv.Atoi(p); err == nil && parsed > 0 {
				page = parsed
			}
		}

		if l := c.Query("limit"); l != "" {
			if parsed, err := strconv.Atoi(l); err == nil && parsed > 0 && parsed <= 100 {
				limit = parsed
			}
		}

		// Get all tickets (you'll need to implement this in database)
		tickets, total, err := db.GetAllTickets(page, limit)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve tickets"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": gin.H{
				"tickets": tickets,
				"pagination": gin.H{
					"page":  page,
					"limit": limit,
					"total": total,
				},
			},
			"message": "Tickets retrieved successfully",
		})
	}
}

// GetAllBaggage returns all baggage for admin
func GetAllBaggage(db database.Database) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check admin role
		_, authorized := checkAdminRole(c, db)
		if !authorized {
			return
		}

		// Get pagination parameters
		page := 1
		limit := 50

		if p := c.Query("page"); p != "" {
			if parsed, err := strconv.Atoi(p); err == nil && parsed > 0 {
				page = parsed
			}
		}

		if l := c.Query("limit"); l != "" {
			if parsed, err := strconv.Atoi(l); err == nil && parsed > 0 && parsed <= 100 {
				limit = parsed
			}
		}

		// Get all baggage (you'll need to implement this in database)
		baggage, total, err := db.GetAllBaggage(page, limit)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve baggage"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": gin.H{
				"baggage": baggage,
				"pagination": gin.H{
					"page":  page,
					"limit": limit,
					"total": total,
				},
			},
			"message": "Baggage retrieved successfully",
		})
	}
}

// AdminRoutes sets up admin routes
func AdminRoutes(router *gin.RouterGroup, db database.Database) {
	// Admin dashboard
	router.GET("/dashboard", GetAdminDashboard(db))

	// User management
	router.GET("/users", GetAllUsers(db))
	router.GET("/users/:id", GetUserById(db))
	router.PUT("/users/:id", UpdateUser(db))
	router.DELETE("/users/:id", DeactivateUser(db))

	// Ticket management
	router.GET("/tickets", GetAllTickets(db))

	// Baggage management
	router.GET("/baggage", GetAllBaggage(db))
}
