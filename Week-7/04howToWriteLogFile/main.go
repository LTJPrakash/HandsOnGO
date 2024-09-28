package main

import (
	"io"
	"os"

	"github.com/gin-gonic/gin" // Import the Gin web framework package
)

func main() {
	// Disable console color. This is helpful when writing logs to a file, where console color codes are unnecessary.
	gin.DisableConsoleColor()

	// Create a log file to write Gin logs.
	f, _ := os.Create("gin.log") // "gin.log" will store all the logs from the Gin application.

	// Set Gin's default writer to write logs into the file. By default, Gin logs are printed to the console.
	gin.DefaultWriter = io.MultiWriter(f)

	// If you want to write logs both to the file and the console, use the following line:
	// gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	// Create a new Gin router with default middleware (logging and recovery).
	router := gin.Default()

	// Define a simple GET route that responds with "pong" to any request to "/ping".
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong") // Respond with status 200 and the string "pong"
	})

	// Start the server on port 5000
	// By default, it listens on localhost:5000, which can be accessed by your browser or API clients
	router.Run(":5000") // Listen and serve on 0.0.0.0:5000
}
