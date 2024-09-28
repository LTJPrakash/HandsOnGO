package main

import (
	"net/http"

	"github.com/gin-gonic/gin" // Import the Gin web framework package
)

func main() {
	// Initialize the Gin router with default middleware (logging and recovery)
	router := gin.Default()

	// Route to match /user/:name (e.g., /user/john)
	router.GET("/user/:name", func(c *gin.Context) {
		// Extract the path parameter "name"
		name := c.Param("name")
		// Respond with a greeting message
		c.String(http.StatusOK, "Hello %s", name)
	})

	// Route to match /user/:name/*action (e.g., /user/john/send)
	// This will also match /user/john/ (and redirect if no action is provided)
	router.GET("/user/:name/*action", func(c *gin.Context) {
		// Extract the path parameters "name" and "action"
		name := c.Param("name")
		action := c.Param("action")
		// Construct a message based on the parameters
		message := name + " is " + action
		// Respond with the constructed message
		c.String(http.StatusOK, message)
	})

	// Start the server on port 5000
	router.Run(":5000") // Listen and serve on 0.0.0.0:5000
}
