package main

import (
	"github.com/gin-gonic/gin"

	_ "github.com/godror/godror"

	"mindenairport/database"
	"mindenairport/initializers"
	"mindenairport/routers"
)

var db database.Database

func init() {
	initializers.LoadEnvs()
	db = database.CreateConnection()
}

func main() {

	router := gin.Default()
	apiRouter := router.Group("/api")
	routers.AirlineRoutes(apiRouter.Group("/airline"), db)
	routers.AirportRoutes(apiRouter.Group("/airport"), db)
	routers.FlightStatusRoutes(apiRouter.Group("/flightStatus"), db)
	routers.FlightRoutes(apiRouter.Group("/flight"), db)
	routers.TicketRoutes(apiRouter.Group("/ticket"), db)

	router.Run("localhost:8080")
}
