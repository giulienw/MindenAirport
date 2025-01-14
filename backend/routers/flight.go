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
		c.IndentedJSON(http.StatusOK, db.GetFlightByID(id))
	}

	return gin.HandlerFunc(fn)
}

func FlightRoutes(router *gin.RouterGroup, db database.Database) {
	router.GET("/", GetFlights(db))
	router.GET("/:id", GetFlightByID(db))
}
