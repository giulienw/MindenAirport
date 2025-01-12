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
	routers.AirlineRoutes(router.Group("/airline"), db)
	routers.AirportRoutes(router.Group("/airport"), db)
	router.Run("localhost:8080")
}
