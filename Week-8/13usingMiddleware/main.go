package main

import (
	"net/http"

	"github.com/gin-gonic/gin" // Import the Gin web framework package
)

// AuthRequired is a middleware function that checks if the user is authorized
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		// For demonstration, we'll just accept all requests.
		// In a real application, you'd check for authorization tokens, etc.
		c.Next() // Proceed to the next handler
	}
}

// MyBenchLogger is a middleware function that logs benchmark data
func MyBenchLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		c.Next() // Process the request

		// Log the request duration
		duration := c.Request.Context().Value("duration")
		if duration != nil {
			c.Writer.Header().Set("X-Request-Duration", duration.(string))
		}
	}
}

// Handler for the benchmark endpoint
func benchEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Benchmark endpoint reached"})
}

// Handler for the login endpoint
func loginEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Login endpoint"})
}

// Handler for the submit endpoint
func submitEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Submit endpoint"})
}

// Handler for the read endpoint
func readEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Read endpoint"})
}

// Handler for the analytics endpoint
func analyticsEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Analytics endpoint"})
}

func main() {
	// Create a router without any middleware by default
	r := gin.New()

	// Global middleware
	r.Use(gin.Logger())  // Logger middleware
	r.Use(gin.Recovery()) // Recovery middleware

	// Per route middleware for benchmark logging
	r.GET("/benchmark", MyBenchLogger(), benchEndpoint)

	// Authorization group with AuthRequired middleware
	authorized := r.Group("/")
	authorized.Use(AuthRequired())
	{
		authorized.POST("/login", loginEndpoint)
		authorized.POST("/submit", submitEndpoint)
		authorized.POST("/read", readEndpoint)

		// Nested group for testing
		testing := authorized.Group("testing")
		testing.GET("/analytics", analyticsEndpoint)
	}

	// Listen and serve on 0.0.0.0:5000
	r.Run(":5000") // Start the server on port 5000
}
