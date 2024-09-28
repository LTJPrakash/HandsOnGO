package main

import (
	"github.com/gin-gonic/gin" // Import the Gin web framework package
)

// Define the endpoint functions (placeholders for now)
func loginEndpoint(c *gin.Context) {
	c.JSON(200, gin.H{"message": "login successful"})
}

func submitEndpoint(c *gin.Context) {
	c.JSON(200, gin.H{"message": "submission successful"})
}

func readEndpoint(c *gin.Context) {
	c.JSON(200, gin.H{"message": "reading successful"})
}

func main() {
	// Create a new Gin router with default middleware (logging and recovery)
	r := gin.Default()

	// Simple group: v1
	v1 := r.Group("/v1") // Group for version 1 of the API
	{
		v1.POST("/login", loginEndpoint)   // Route for login
		v1.POST("/submit", submitEndpoint) // Route for submitting data
		v1.POST("/read", readEndpoint)     // Route for reading data
	}

	// Simple group: v2
	v2 := r.Group("/v2") // Group for version 2 of the API
	{
		v2.POST("/login", loginEndpoint)   // Route for login
		v2.POST("/submit", submitEndpoint) // Route for submitting data
		v2.POST("/read", readEndpoint)     // Route for reading data
	}

	// Start the server on port 5000
	r.Run(":5000") // The server will listen and serve on 0.0.0.0:5000
}
