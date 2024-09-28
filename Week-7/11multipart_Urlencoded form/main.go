package main

import (
	"github.com/gin-gonic/gin" // Import the Gin web framework package
)

func main() {
	// Initialize the Gin router with default middleware (logging and recovery)
	router := gin.Default()

	// Define a POST route for handling form submissions at "/form_post"
	router.POST("/form_post", func(c *gin.Context) {
		// Extract the "message" field from the form data
		message := c.PostForm("message")

		// Extract the "nick" field from the form data with a default value of "anonymous"
		nick := c.DefaultPostForm("nick", "anonymous")

		// Respond back with the posted data in JSON format
		c.JSON(200, gin.H{
			"status":  "posted", // Status of the post request
			"message": message,  // The message that was posted
			"nick":    nick,     // The nickname, or "anonymous" if not provided
		})
	})

	// Start the server on port 5000
	router.Run(":5000") // Listen and serve on 0.0.0.0:5000
}
