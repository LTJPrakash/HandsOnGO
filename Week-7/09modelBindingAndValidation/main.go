package main

import (
	"net/http"

	"github.com/gin-gonic/gin" // Import the Gin web framework package
)

// Login struct for binding user credentials
type Login struct {
	User     string `form:"user" json:"user" xml:"user" binding:"required"`             // User field with tags for different formats
	Password string `form:"password" json:"password" xml:"password" binding:"required"` // Password field with tags for different formats
}

func main() {
	// Create a new Gin router with default middleware (logging and recovery)
	router := gin.Default()

	// Example for binding JSON
	router.POST("/loginJSON", func(c *gin.Context) {
		var json Login
		// Bind the incoming JSON to the Login struct
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Validate user credentials
		if json.User != "manu" || json.Password != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}

		// Successful login response
		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})

	// Example for binding XML
	router.POST("/loginXML", func(c *gin.Context) {
		var xml Login
		// Bind the incoming XML to the Login struct
		if err := c.ShouldBindXML(&xml); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Validate user credentials
		if xml.User != "manu" || xml.Password != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}

		// Successful login response
		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})

	// Example for binding a HTML form
	router.POST("/loginForm", func(c *gin.Context) {
		var form Login
		// Infer binder to use based on content-type header
		if err := c.ShouldBind(&form); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Validate user credentials
		if form.User != "manu" || form.Password != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}

		// Successful login response
		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})

	// Start the server on port 5000
	router.Run(":5000") // Listen and serve on 0.0.0.0:5000
}
