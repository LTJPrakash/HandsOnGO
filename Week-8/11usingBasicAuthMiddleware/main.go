package main

import (
	"net/http"

	"github.com/gin-gonic/gin" // Import the Gin web framework package
)

// Simulate some private data
var secrets = gin.H{
	"foo":    gin.H{"email": "foo@bar.com", "phone": "123433"},
	"austin": gin.H{"email": "austin@example.com", "phone": "666"},
	"lena":   gin.H{"email": "lena@guapa.com", "phone": "523443"},
}

func main() {
	// Create a new Gin router with default middleware (logging and recovery)
	r := gin.Default()

	// Group using gin.BasicAuth() middleware
	// gin.Accounts is a shortcut for map[string]string, where key is the username and value is the password
	authorized := r.Group("/admin", gin.BasicAuth(gin.Accounts{
		"foo":    "bar",
		"austin": "1234",
		"lena":   "hello2",
		"manu":   "4321",
	}))

	// /admin/secrets endpoint
	// Hit "http://localhost:5000/admin/secrets" to access this endpoint
	authorized.GET("/secrets", func(c *gin.Context) {
		// Get user, it was set by the BasicAuth middleware
		user, exists := c.Get(gin.AuthUserKey)
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		// Cast user to string and retrieve the secret for the authenticated user
		if secret, ok := secrets[user.(string)]; ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": secret})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO SECRET :("})
		}
	})

	// Start the server on port 5000
	// By default, it listens on localhost:5000, which can be accessed by your browser or API clients
	r.Run(":5000") // Listen and serve on 0.0.0.0:5000
}
