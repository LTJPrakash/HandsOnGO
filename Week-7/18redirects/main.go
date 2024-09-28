package main

import (
	"github.com/gin-gonic/gin" // Import the Gin web framework package
)

func main() {
	// Initialize the Gin router with default middleware (logging and recovery)
	r := gin.Default()

	// Define a GET route for /test
	r.GET("/test", func(c *gin.Context) {
		// Modify the request URL path to /test2
		c.Request.URL.Path = "/test2"
		// Handle the modified request context
		r.HandleContext(c)
	})

	// Define a GET route for /test2
	r.GET("/test2", func(c *gin.Context) {
		// Respond with a JSON object
		c.JSON(200, gin.H{"hello": "world"})
	})

	// Start the server on port 5000
	r.Run(":5000") // Listen and serve on 0.0.0.0:5000
}
