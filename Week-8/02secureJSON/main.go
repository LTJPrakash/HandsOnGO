package main

import (
	"net/http"

	"github.com/gin-gonic/gin" // Import the Gin web framework package
)

func main() {
	// Create a new Gin router with default middleware (logging and recovery)
	r := gin.Default()

	// Define a GET route that returns secure JSON
	r.GET("/someJSON", func(c *gin.Context) {
		// Create a sample list of names to return as JSON
		names := []string{"lena", "austin", "foo"}

		// Use SecureJSON to prevent JSON Hijacking attacks
		// This adds a prefix to the JSON response to make it safer when served to browsers
		// The response will be prefixed with "while(1);" which breaks malicious script execution
		c.SecureJSON(http.StatusOK, names)
	})

	// Start the server on port 5000
	// It will listen on 0.0.0.0:5000, meaning it is accessible on localhost:5000
	r.Run(":5000") // Listen and serve
}
