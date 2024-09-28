package main

import (
	"github.com/gin-gonic/gin" // Import the Gin web framework package
)

// LoginForm represents the structure for login form data
type LoginForm struct {
	User     string `form:"user" binding:"required"`     // User field with required binding
	Password string `form:"password" binding:"required"` // Password field with required binding
}

func main() {
	// Create a new Gin router with default middleware (logging and recovery)
	router := gin.Default()

	// Define the login route with POST method
	router.POST("/login", func(c *gin.Context) {
		var form LoginForm

		// Bind the incoming form data to the LoginForm struct
		if err := c.ShouldBind(&form); err == nil {
			// Check if the credentials are correct
			if form.User == "user" && form.Password == "password" {
				// Successful login response
				c.JSON(200, gin.H{"status": "you are logged in"})
			} else {
				// Unauthorized response
				c.JSON(401, gin.H{"status": "unauthorized"})
			}
		} else {
			// Bad request response if binding fails
			c.JSON(400, gin.H{"error": err.Error()})
		}
	})

	// Start the server on port 5000
	router.Run(":5000") // Listen and serve on 0.0.0.0:5000
}
