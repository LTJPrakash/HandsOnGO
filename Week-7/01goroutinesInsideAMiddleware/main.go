package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin" // Import the Gin web framework package
)

func main() {
	// Create a new Gin router with default middleware (logging and recovery)
	r := gin.Default()

	// Route handler for asynchronous tasks
	r.GET("/long_async", func(c *gin.Context) {
		// Create a copy of the context to safely use in a Goroutine
		cCp := c.Copy()

		// Start a new Goroutine for processing a long-running task
		go func() {
			// Simulate a long task (5 seconds delay)
			time.Sleep(5 * time.Second)

			// Log the request path using the copied context
			// IMPORTANT: Use the copied context (cCp) to avoid race conditions
			log.Println("Done! in path " + cCp.Request.URL.Path)
		}()
	})

	// Route handler for synchronous tasks
	r.GET("/long_sync", func(c *gin.Context) {
		// Simulate a long task (5 seconds delay)
		time.Sleep(5 * time.Second)

		// Log the request path using the original context
		// Since there is no Goroutine, the original context is safe to use
		log.Println("Done! in path " + c.Request.URL.Path)
	})

	// Start the server and listen on port 5000
	// By default, it listens on localhost:5000, accessible through browser or API clients
	r.Run(":5000")
}
