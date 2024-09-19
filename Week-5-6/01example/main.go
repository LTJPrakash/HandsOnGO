package main

import (
	"github.com/gin-gonic/gin" // Import the Gin web framework package
)

func main() {
	// Create a new Gin router with default middleware (logging and recovery)
	r := gin.Default()

	// Define a GET route at "/ping"
	// When a GET request is made to this endpoint, the server responds with a JSON object
	r.GET("/ping", func(c *gin.Context) {
		// Respond with a JSON object containing the message "pong"
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Start the server on port 5000
	// By default, it listens on localhost:5000, which can be accessed by your browser or API clients
	r.Run(":5000") // Listen and serve on 0.0.0.0:5000
}
