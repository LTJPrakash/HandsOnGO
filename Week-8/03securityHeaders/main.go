package main

import (
	"net/http"

	"github.com/gin-gonic/gin" // Import the Gin web framework package
)

func main() {
	// Create a new Gin router with default middleware (logging and recovery)
	r := gin.Default()

	// Specify the expected host (domain) that is allowed to access the server
	expectedHost := "localhost:5000"

	// Middleware to set security headers and validate the Host header
	r.Use(func(c *gin.Context) {
		// Check if the request's Host header matches the expected host
		if c.Request.Host != expectedHost {
			// If the Host header is invalid, abort the request with a 400 Bad Request status
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid host header"})
			return
		}

		// Add security headers to the response
		c.Header("X-Frame-Options", "DENY")                                                                                                                                    // Prevent the page from being loaded in an iframe (mitigates clickjacking)
		c.Header("Content-Security-Policy", "default-src 'self'; connect-src *; font-src *; script-src-elem * 'unsafe-inline'; img-src * data:; style-src * 'unsafe-inline';") // Control where resources can be loaded from
		c.Header("X-XSS-Protection", "1; mode=block")                                                                                                                          // Enable browser XSS protection
		c.Header("Strict-Transport-Security", "max-age=31536000; includeSubDomains; preload")                                                                                  // Force HTTPS for one year
		c.Header("Referrer-Policy", "strict-origin")                                                                                                                           // Limit referrer data shared to external sites
		c.Header("X-Content-Type-Options", "nosniff")                                                                                                                          // Prevent MIME-type sniffing
		c.Header("Permissions-Policy", "geolocation=(),midi=(),sync-xhr=(),microphone=(),camera=(),magnetometer=(),gyroscope=(),fullscreen=(self),payment=()")                 // Restrict browser features

		// Continue with the request processing
		c.Next()
	})

	// Define a simple GET route for the "/ping" path
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong", // Respond with "pong" when the /ping route is accessed
		})
	})

	// Start the server on port 5000, accessible via localhost:5000
	r.Run(":5000") // Listen and serve on 0.0.0.0:5000
}
