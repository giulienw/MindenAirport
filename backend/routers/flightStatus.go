package routers

import (
	"net/http"

	"mindenairport/database"

	"github.com/gin-gonic/gin"
)

func GetFlightStatuses(db database.Database) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		// Your handler code goes in here - e.g.
		c.IndentedJSON(http.StatusOK, db.GetFlightStatuses())
	}

	return gin.HandlerFunc(fn)
}

func GetFlightStatusByID(db database.Database) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		id := c.Param("id")
		c.IndentedJSON(http.StatusOK, db.GetFlightStatusByID(id))
	}

	return gin.HandlerFunc(fn)
}

func FlightStatusRoutes(router *gin.RouterGroup, db database.Database) {
	router.GET("/", GetFlightStatuses(db))
	router.GET("/:id", GetFlightStatusByID(db))
}
