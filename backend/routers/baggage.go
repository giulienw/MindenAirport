package routers

import (
	"fmt"
	"net/http"

	"mindenairport/database"
	"mindenairport/models"

	"github.com/gin-gonic/gin"
)

// GetBaggageByID retrieves a specific baggage by ID
func GetBaggageByID(db database.Database) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		baggage, err := db.GetBaggageByID(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve baggage"})
			return
		}

		if baggage == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Baggage not found"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data":    baggage,
			"message": "Baggage retrieved successfully",
		})
	}
}

// GetMyBaggage retrieves all baggage for the authenticated user
func GetMyBaggage(db database.Database) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get user ID from context (set by auth middleware)
		userID, exists := c.Get("userID")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		// Get baggage for the user
		baggageList, err := db.GetBaggageByUserID(userID.(string))
		if err != nil {
			fmt.Println("Error retrieving baggage:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve baggage"})
			return
		}

		// Return empty array if no baggage found
		if baggageList == nil {
			baggageList = []models.Baggage{}
		}

		c.JSON(http.StatusOK, gin.H{
			"data":    baggageList,
			"count":   len(baggageList),
			"message": "Baggage retrieved successfully",
		})
	}
}

// GetBaggageByFlight retrieves all baggage for a specific flight
func GetBaggageByFlight(db database.Database) gin.HandlerFunc {
	return func(c *gin.Context) {
		flightID := c.Param("flightId")

		baggageList, err := db.GetBaggageByFlightID(flightID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve baggage for flight"})
			return
		}

		// Return empty array if no baggage found
		if baggageList == nil {
			baggageList = []models.Baggage{}
		}

		c.JSON(http.StatusOK, gin.H{
			"data":    baggageList,
			"count":   len(baggageList),
			"message": "Flight baggage retrieved successfully",
		})
	}
}

// GetBaggageByTrackingNumber retrieves baggage by tracking number
func GetBaggageByTrackingNumber(db database.Database) gin.HandlerFunc {
	return func(c *gin.Context) {
		trackingNumber := c.Query("tracking")
		if trackingNumber == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Tracking number is required"})
			return
		}

		baggage, err := db.GetBaggageByTrackingNumber(trackingNumber)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve baggage"})
			return
		}

		if baggage == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Baggage not found with the provided tracking number"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data":    baggage,
			"message": "Baggage found successfully",
		})
	}
}

// CreateBaggage creates a new baggage entry
func CreateBaggage(db database.Database) gin.HandlerFunc {
	return func(c *gin.Context) {
		var baggage models.Baggage

		if err := c.ShouldBindJSON(&baggage); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data", "details": err.Error()})
			return
		}

		// Get user ID from context (set by auth middleware)
		userID, exists := c.Get("userID")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		// Set the user ID for the baggage
		baggage.AirportUserID = userID.(string)

		// Validate required fields
		if baggage.FlightID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Flight ID is required"})
			return
		}

		if baggage.Weight <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Weight must be greater than 0"})
			return
		}

		// Create the baggage
		createdBaggage, err := db.CreateBaggage(baggage)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create baggage"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"data":    createdBaggage,
			"message": "Baggage created successfully",
		})
	}
}

// UpdateBaggage updates an existing baggage entry
func UpdateBaggage(db database.Database) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var baggage models.Baggage

		if err := c.ShouldBindJSON(&baggage); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data", "details": err.Error()})
			return
		}

		// Get user ID from context (set by auth middleware)
		userID, exists := c.Get("userID")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		// Check if the baggage exists and belongs to the user
		existingBaggage, err := db.GetBaggageByID(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve baggage"})
			return
		}

		if existingBaggage == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Baggage not found"})
			return
		}

		if existingBaggage.AirportUserID != userID.(string) {
			c.JSON(http.StatusForbidden, gin.H{"error": "You can only update your own baggage"})
			return
		}

		// Validate required fields
		if baggage.Weight <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Weight must be greater than 0"})
			return
		}

		// Update the baggage
		updatedBaggage, err := db.UpdateBaggage(id, baggage)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update baggage"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data":    updatedBaggage,
			"message": "Baggage updated successfully",
		})
	}
}

// DeleteBaggage deletes a baggage entry
func DeleteBaggage(db database.Database) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		// Get user ID from context (set by auth middleware)
		userID, exists := c.Get("userID")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		// Check if the baggage exists and belongs to the user
		existingBaggage, err := db.GetBaggageByID(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve baggage"})
			return
		}

		if existingBaggage == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Baggage not found"})
			return
		}

		if existingBaggage.AirportUserID != userID.(string) {
			c.JSON(http.StatusForbidden, gin.H{"error": "You can only delete your own baggage"})
			return
		}

		// Delete the baggage
		err = db.DeleteBaggage(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete baggage"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Baggage deleted successfully",
		})
	}
}

// BaggageRoutes sets up baggage routes
func BaggageRoutes(router *gin.RouterGroup, db database.Database) {
	// Protected routes (require authentication)
	router.GET("/my", GetMyBaggage(db)) // Get authenticated user's baggage
	router.POST("/", CreateBaggage(db)) // Create new baggage
	//router.GET("/track", GetBaggageByTrackingNumber(db))    // Track baggage by tracking number
	router.GET("/flight/:flightId", GetBaggageByFlight(db)) // Get baggage for specific flight
	router.GET("/:id", GetBaggageByID(db))                  // Get specific baggage by ID
	router.PUT("/:id", UpdateBaggage(db))                   // Update baggage
	router.DELETE("/:id", DeleteBaggage(db))                // Delete baggage
}
