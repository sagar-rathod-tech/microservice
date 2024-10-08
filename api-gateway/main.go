package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sagar-rathod-tech/api-gateway/middleware"
	"github.com/sagar-rathod-tech/api-gateway/routes"
)

func main() {
	// Create a new Gin router
	r := gin.Default()

	// Apply middleware globally
	r.Use(middleware.LoggingMiddleware())        // Logging for all requests
	r.Use(middleware.AuthenticationMiddleware()) // Authentication for all requests

	// Register routes (forwards requests to respective services)
	routes.RegisterRoutes(r)

	// Start the API Gateway on port 8080
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
