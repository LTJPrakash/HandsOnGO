package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// Logger is a custom middleware function that logs request latency and status.
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Record the start time
		t := time.Now()

		// Set a custom context variable
		c.Set("example", "12345")

		// Proceed with the request
		c.Next()

		// Calculate the latency of the request
		latency := time.Since(t)
		log.Print(latency) // Log the latency

		// Access and log the response status code
		status := c.Writer.Status()
		log.Println(status) // Log the status code
	}
}

// AuthMiddleware is a custom middleware function for authentication (commented out).
// func AuthMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		token := c.GetHeader("Authorization")
// 		if token != "valid-token" {
// 			c.AbortWithStatus(401) // Abort and send a 401 status if the token is invalid
// 			return
// 		}
// 		c.Next() // Continue if the token is valid
// 	}
// }

func main() {
	r := gin.New() // Create a new Gin router without default middleware

	// Use the custom Logger middleware
	r.Use(Logger())
	// Use the custom AuthMiddleware (commented out)
	// r.Use(AuthMiddleware())

	// Define a simple GET route
	r.GET("/test", func(c *gin.Context) {
		// Retrieve and print the custom context variable
		example := c.MustGet("example").(string)
		log.Println(example) // Output: "12345"
	})

	// Start the server on port 5000
	r.Run(":5000")
}
