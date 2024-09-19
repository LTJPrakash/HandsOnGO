package main

import "github.com/gin-gonic/gin"

func main() {
	// Disable log's color
	// Uncomment the next line to disable color output in the logs
	// gin.DisableConsoleColor()

	// Force log's color
	// This forces the logs to always use color, even if the console doesn’t support it
	gin.ForceConsoleColor()

	// Create a new Gin router with default middleware (logging and recovery)
	r := gin.Default()

	// Define a simple GET route at "/ping"
	r.GET("/ping", func(c *gin.Context) {
		// Send a JSON response with a message
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Start the server on port 5000
	r.Run(":5000")
}
