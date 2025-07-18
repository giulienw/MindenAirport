// Package main provides the entry point for the MindenAirport backend API server.
// This application provides a RESTful API for managing airport operations including
// flights, passengers, baggage, tickets, and administrative functions.
//
// The API is built using the Gin web framework and connects to an Oracle database
// using stored procedures for data operations.
//
// Main features:
//   - User authentication and authorization with JWT tokens
//   - Flight management and tracking
//   - Baggage tracking system
//   - Ticket booking and management
//   - Administrative dashboard and user management
//   - Airport and airline information management
//
// The server runs on port 8080 by default and provides endpoints under /api.
package main

import (
	"github.com/gin-gonic/gin"

	_ "github.com/godror/godror" // Oracle database driver

	"mindenairport/database"
	"mindenairport/initializers"
	"mindenairport/middleware"
	"mindenairport/routers"
)

// db is the global database connection instance used throughout the application
var db database.Database

// init initializes the application by loading environment variables
// and establishing the database connection before main() runs
func init() {
	initializers.LoadEnvs()
	db = database.CreateConnection()
}

// main sets up the HTTP server with all routes and middleware,
// then starts listening for requests on port 8080
func main() {
	router := gin.Default()

	// Configure CORS - use custom CORS middleware for proper frontend access

	// Create API router group for versioning and organization
	apiRouter := router.Group("/api")

	// Apply CORS middleware to all API routes
	apiRouter.Use(middleware.CORSMiddleware())

	// Health check endpoint for monitoring and load balancers
	apiRouter.HEAD("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "healthy", "message": "MindenAirport API is running"})
	})

	// ======= PUBLIC ROUTES (no authentication required) =======

	// Authentication routes - registration, login, password reset
	routers.AuthRoutes(apiRouter.Group("/auth"), db)

	// Public information routes accessible to all users
	routers.AirlineRoutes(apiRouter.Group("/airline"), db)
	routers.AirportRoutes(apiRouter.Group("/airport"), db)
	routers.FlightStatusRoutes(apiRouter.Group("/flightStatus"), db)
	routers.FlightRoutes(apiRouter.Group("/flight"), db)

	// Public baggage tracking - allows tracking without authentication
	publicBaggage := apiRouter.Group("/baggage")
	publicBaggage.GET("/track", routers.GetBaggageByTrackingNumber(db))

	// ======= PROTECTED ROUTES (authentication required) =======

	// Protected routes that require valid JWT token
	protected := apiRouter.Group("/")
	protected.Use(middleware.AuthMiddleware())
	routers.TicketRoutes(protected.Group("/ticket"), db)
	routers.BaggageRoutes(protected.Group("/baggage"), db)

	// ======= ADMIN ROUTES (authentication + admin role required) =======

	// Admin routes for administrative functions
	adminProtected := apiRouter.Group("/admin")
	adminProtected.Use(middleware.AuthMiddleware())
	routers.AdminRoutes(adminProtected, db)

	// ======= PROTECTED AUTH ROUTES =======

	// Protected authentication routes for logged-in users
	authProtected := apiRouter.Group("/auth")
	authProtected.Use(middleware.AuthMiddleware())
	authProtected.GET("/profile", routers.GetProfile(db))
	authProtected.GET("/dashboard", routers.GetDashboard(db))
	authProtected.POST("/refresh", routers.RefreshToken(db))

	// Start the HTTP server on port 8080
	router.Run(":8080")
}
