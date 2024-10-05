package main

import (
	"net/http"

	"github.com/gin-gonic/gin" // Import the Gin web framework package
)

// Handler for GET requests
func getting(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "GET request received"})
}

// Handler for POST requests
func posting(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "POST request received"})
}

// Handler for PUT requests
func putting(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "PUT request received"})
}

// Handler for DELETE requests
func deleting(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "DELETE request received"})
}

// Handler for PATCH requests
func patching(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "PATCH request received"})
}

// Handler for HEAD requests
func head(c *gin.Context) {
	// You can set headers here if needed, but no body is returned for HEAD requests
	c.Status(http.StatusOK)
}

// Handler for OPTIONS requests
func options(c *gin.Context) {
	// Returning allowed methods for this endpoint
	c.Header("Allow", "GET, POST, PUT, DELETE, PATCH, HEAD, OPTIONS")
	c.Status(http.StatusOK)
}

func main() {
	// Create a new Gin router with default middleware (logging and recovery)
	r := gin.Default()

	// Define routes for different HTTP methods
	r.GET("/someGet", getting)
	r.POST("/somePost", posting)
	r.PUT("/somePut", putting)
	r.DELETE("/someDelete", deleting)
	r.PATCH("/somePatch", patching)
	r.HEAD("/someHead", head)
	r.OPTIONS("/someOptions", options)

	// Start the server on port 5000
	// By default, it listens on localhost:5000, which can be accessed by your browser or API clients
	r.Run(":5000") // Listen and serve on 0.0.0.0:5000
}
