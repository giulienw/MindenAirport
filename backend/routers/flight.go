// Package routers provides HTTP route handlers for flight-related operations
// in the MindenAirport API system.
package routers

import (
	"net/http"

	"mindenairport/database"

	"github.com/gin-gonic/gin"
)

// GetFlights handles requests to retrieve all flights in the system.
// This endpoint provides flight information for public viewing including
// schedules, status, and basic flight details.
//
// Response includes:
//   - Flight ID and basic information
//   - Origin and destination airports
//   - Scheduled departure and arrival times
//   - Current flight status
//   - Gate and baggage claim information
//
// Returns:
//   - 200: List of all flights
//   - 500: Internal server error
func GetFlights(db database.Database) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		// Retrieve all flights from database
		c.IndentedJSON(http.StatusOK, db.GetFlights())
	}

	return gin.HandlerFunc(fn)
}

// GetFlightByID handles requests to retrieve a specific flight by its ID.
// This endpoint provides detailed information about a single flight.
//
// URL Parameters:
//   - id: The unique flight identifier
//
// Returns:
//   - 200: Flight details if found
//   - 500: Internal server error or flight not found
func GetFlightByID(db database.Database) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		id := c.Param("id")
		flight, err := db.GetFlightByID(id)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve flight"})
			return
		}

		c.IndentedJSON(http.StatusOK, flight)
	}

	return gin.HandlerFunc(fn)
}

// FlightRoutes sets up all flight-related routes on the provided router group.
// This includes both public flight information endpoints.
//
// Routes:
//   - GET /flight/ - Get all flights
//   - GET /flight/:id - Get specific flight by ID
func FlightRoutes(router *gin.RouterGroup, db database.Database) {
	router.GET("/", GetFlights(db))
	router.GET("/:id", GetFlightByID(db))
}
