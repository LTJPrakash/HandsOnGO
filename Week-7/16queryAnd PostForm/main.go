package main

import (
	"fmt"

	"github.com/gin-gonic/gin" // Import the Gin web framework package
)

func main() {
	// Initialize the Gin router with default middleware (logging and recovery)
	router := gin.Default()

	// Define a POST route for handling incoming data
	router.POST("/post", func(c *gin.Context) {
		// Extract query parameters
		id := c.Query("id")                 // Get the 'id' query parameter
		page := c.DefaultQuery("page", "0") // Get the 'page' query parameter, default to "0" if not provided

		// Extract form parameters
		name := c.PostForm("name")       // Get the 'name' form parameter
		message := c.PostForm("message") // Get the 'message' form parameter

		// Log the received data to the console
		fmt.Printf("id: %s; page: %s; name: %s; message: %s\n", id, page, name, message)

		// Respond back with the received data for demonstration
		c.JSON(200, gin.H{
			"id":      id,
			"page":    page,
			"name":    name,
			"message": message,
		})
	})

	// Start the server on port 5000
	router.Run(":5000") // Listen and serve on 0.0.0.0:5000
}
