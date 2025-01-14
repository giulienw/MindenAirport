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
func GetAirlines(db database.Database) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		// Your handler code goes in here - e.g.
		c.IndentedJSON(http.StatusOK, db.GetAirlines())
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
func GetAirlineByID(db database.Database) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		id := c.Param("id")
		c.IndentedJSON(http.StatusOK, db.GetAirlineByID(id))
	}

	return gin.HandlerFunc(fn)
}

func AirlineRoutes(router *gin.RouterGroup, db database.Database) {
	router.GET("/", GetAirlines(db))
	router.GET("/:id", GetAirlineByID(db))
}
