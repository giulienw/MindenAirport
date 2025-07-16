package routers

import (
	"net/http"

	"mindenairport/database"
	"mindenairport/models"

	"github.com/gin-gonic/gin"
)

func GetTicketByID(db database.Database) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		id := c.Param("id")

		ticket, err := db.GetTicketByID(id)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve ticket"})
			return
		}
		c.IndentedJSON(http.StatusOK, ticket)
	}

	return gin.HandlerFunc(fn)
}

// GetMyTickets retrieves all tickets for the authenticated user
func GetMyTickets(db database.Database) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get user ID from context (set by auth middleware)
		userID, exists := c.Get("userID")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		// Get tickets for the user
		tickets, err := db.GetTicketsByUserID(userID.(string))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve tickets"})
			return
		}

		// Return empty array if no tickets found
		if tickets == nil {
			tickets = []models.Ticket{}
		}

		c.JSON(http.StatusOK, gin.H{
			"data":    tickets,
			"count":   len(tickets),
			"message": "Tickets retrieved successfully",
		})
	}
}

func TicketRoutes(router *gin.RouterGroup, db database.Database) {
	router.GET("/my", GetMyTickets(db))   // Get authenticated user's tickets
	router.GET("/:id", GetTicketByID(db)) // Get specific ticket by ID
}
