package routers

import (
	"net/http"

	"mindenairport/database"

	"github.com/gin-gonic/gin"
)

// GetAirlines godoc
// @Summary Get all airlines
// @Description Get all airlines
// @Produce json
// @Success 200 {array} models.Airline
// @Router /airline [get]
func GetAirports(db database.Database) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		// Your handler code goes in here - e.g.
		c.IndentedJSON(http.StatusOK, db.GetAirports())
	}

	return gin.HandlerFunc(fn)
}

// GetAirlineByID godoc
// @Summary Get airline by ID
// @Description Get airline by ID
// @Produce json
// @Param id path string true "Airline ID"
// @Success 200 {object} models.Airline
// @Router /airline/{id} [get]
func GetAirportByID(db database.Database) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		id := c.Param("id")
		c.IndentedJSON(http.StatusOK, db.GetAirportByID(id))
	}

	return gin.HandlerFunc(fn)
}

func AirportRoutes(router *gin.RouterGroup, db database.Database) {
	router.GET("/", GetAirports(db))
	router.GET("/:id", GetAirportByID(db))
}
