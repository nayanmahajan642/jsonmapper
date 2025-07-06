package main

import (
	"time"

	"github.com/gin-contrib/cors" // Middleware to handle CORS (Cross-Origin Resource Sharing)
	"github.com/gin-gonic/gin"
	"github.com/nayanmahajan642/jsonmapper/controller" // Controller layer for request handling
)

func main() {
	// Initialize a new Gin router with default middleware: logger and recovery (crash-free)
	r := gin.Default()

	// Set up CORS to allow requests from the frontend running on localhost:3000
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},  // Frontend origin
		AllowMethods:     []string{"GET", "POST", "OPTIONS"}, // Allowed HTTP methods
		AllowHeaders:     []string{"Origin", "Content-Type"}, // Allowed headers in requests
		ExposeHeaders:    []string{"Content-Length"},         // Headers exposed to frontend
		AllowCredentials: true,                               // Allow cookies or auth headers
		MaxAge:           12 * time.Hour,                     // Cache CORS config for 12 hours
	}))

	// Create a route group for all API routes with prefix "/api"
	api := r.Group("/api")
	{
		// Handle POST requests to "/api/map-json" with the MapHandler function
		api.POST("/map-json", controller.MapHandler)
	}

	// Start the HTTP server on port 8080
	r.Run(":8080")
}
