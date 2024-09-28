package main

import (
	"fmt"

	"github.com/gin-gonic/gin" // Import the Gin web framework package
)

func main() {
	// Initialize the Gin router with default middleware (logging and recovery)
	router := gin.Default()

	// Define a POST route for handling data
	router.POST("/post", func(c *gin.Context) {
		// Extract query parameters as a map
		ids := c.QueryMap("ids")

		// Extract form parameters as a map
		names := c.PostFormMap("names")

		// Print the extracted parameters to the console
		fmt.Printf("ids: %v; names: %v\n", ids, names)

		// Respond back with the received data (for demonstration)
		c.JSON(200, gin.H{
			"ids":   ids,
			"names": names,
		})
	})

	// Start the server on port 5000
	router.Run(":5000") // Listen and serve on 0.0.0.0:5000
}
