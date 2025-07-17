package main

import (
	"github.com/gin-gonic/gin"

	_ "github.com/godror/godror"

	"mindenairport/database"
	"mindenairport/initializers"
	"mindenairport/middleware"
	"mindenairport/routers"
)

var db database.Database

func init() {
	initializers.LoadEnvs()
	db = database.CreateConnection()
}

func main() {
	router := gin.Default()

	// Configure CORS - use custom CORS middleware for proper frontend access
	

	apiRouter := router.Group("/api")

	apiRouter.Use(middleware.CORSMiddleware())

	// Health check endpoint
	apiRouter.HEAD("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "healthy", "message": "MindenAirport API is running"})
	})

	// Authentication routes (public)
	routers.AuthRoutes(apiRouter.Group("/auth"), db)

	// Public routes
	routers.AirlineRoutes(apiRouter.Group("/airline"), db)
	routers.AirportRoutes(apiRouter.Group("/airport"), db)
	routers.FlightStatusRoutes(apiRouter.Group("/flightStatus"), db)
	routers.FlightRoutes(apiRouter.Group("/flight"), db)

	// Public baggage tracking
	publicBaggage := apiRouter.Group("/baggage")
	publicBaggage.GET("/track", routers.GetBaggageByTrackingNumber(db))

	// Protected routes (require authentication)
	protected := apiRouter.Group("/")
	protected.Use(middleware.AuthMiddleware())
	routers.TicketRoutes(protected.Group("/ticket"), db)
	routers.BaggageRoutes(protected.Group("/baggage"), db)

	// Admin routes (require authentication and admin role)
	adminProtected := apiRouter.Group("/admin")
	adminProtected.Use(middleware.AuthMiddleware())
	routers.AdminRoutes(adminProtected, db)

	// Protected auth routes
	authProtected := apiRouter.Group("/auth")
	authProtected.Use(middleware.AuthMiddleware())
	authProtected.GET("/profile", routers.GetProfile(db))
	authProtected.GET("/dashboard", routers.GetDashboard(db))
	authProtected.POST("/refresh", routers.RefreshToken(db))

	router.Run(":8080")
}
