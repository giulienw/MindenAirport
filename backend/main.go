package main

import (
	"github.com/gin-gonic/gin"

	_ "github.com/godror/godror"

	"mindenairport/database"
	"mindenairport/initializers"
	"mindenairport/middleware"
	"mindenairport/routers"

	"time"

	"github.com/gin-contrib/cors"
)

var db database.Database

func init() {
	initializers.LoadEnvs()
	db = database.CreateConnection()
}

func main() {
	router := gin.Default()

	// Configure CORS - use simple CORS middleware for better compatibility
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	apiRouter := router.Group("/api")

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

	router.Run("localhost:8080")
}
