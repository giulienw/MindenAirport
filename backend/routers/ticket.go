package routers

import (
	"net/http"

	"mindenairport/database"

	"github.com/gin-gonic/gin"
)

func GetTicketByID(db database.Database) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		id := c.Param("id")
		c.IndentedJSON(http.StatusOK, db.GetTicketByID(id))
	}

	return gin.HandlerFunc(fn)
}

func TicketRoutes(router *gin.RouterGroup, db database.Database) {
	router.GET("/:id", GetFlightStatusByID(db))
}
