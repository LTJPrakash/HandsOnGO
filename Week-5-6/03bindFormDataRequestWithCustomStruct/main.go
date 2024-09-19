package main

import (
	"github.com/gin-gonic/gin" // Import the Gin web framework
)

// Define a User struct to represent incoming user data
// `binding:"required"` ensures that the fields must be provided when a form is submitted
type User struct {
	Name string `form:"name" binding:"required"` // Bind form data to the "name" field (required)
	Age  int    `form:"age" binding:"required"`  // Bind form data to the "age" field (required)
}

func main() {
	// Create a new Gin router with default middleware (logging and recovery)
	r := gin.Default()

	// Define a POST route at "/user" for handling form submissions
	r.POST("/user", func(c *gin.Context) {
		var user User

		// Bind incoming form data to the user struct
		// If the binding fails (e.g., missing required fields), an error will be returned
		if err := c.ShouldBind(&user); err != nil {
			// If there is an error, return a JSON response with the error message
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		// If binding is successful, return the user data in JSON format
		c.JSON(200, gin.H{"user": user})
	})

	// Start the server on port 5000
	r.Run(":5000")
}
