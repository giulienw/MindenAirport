package routers

import (
	"net/http"

	"mindenairport/database"

	"github.com/gin-gonic/gin"
)

func GetFlights(db database.Database) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		// Your handler code goes in here - e.g.
		c.IndentedJSON(http.StatusOK, db.GetFlights())
	}

	return gin.HandlerFunc(fn)
}

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

func FlightRoutes(router *gin.RouterGroup, db database.Database) {
	router.GET("/", GetFlights(db))
	router.GET("/:id", GetFlightByID(db))
}
