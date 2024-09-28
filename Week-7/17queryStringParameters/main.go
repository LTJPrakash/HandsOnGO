package main

import (
	"net/http"

	"github.com/gin-gonic/gin" // Import the Gin web framework package
)

func main() {
	// Initialize the Gin router with default middleware (logging and recovery)
	router := gin.Default()

	// Define a GET route for handling the /welcome endpoint
	router.GET("/welcome", func(c *gin.Context) {
		// Extract query parameters
		firstname := c.DefaultQuery("firstname", "Guest") // Get 'firstname' query parameter, default to "Guest"
		lastname := c.Query("lastname")                   // Get 'lastname' query parameter

		// Respond with a greeting message
		c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})

	// Start the server on port 5000
	router.Run(":5000") // Listen and serve on 0.0.0.0:5000
}
