package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()

	// Customize the log format using LoggerWithFormatter middleware
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// Define a custom log format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,                       // Client IP address
			param.TimeStamp.Format(time.RFC1123), // Timestamp in RFC1123 format
			param.Method,                         // HTTP method (GET, POST, etc.)
			param.Path,                           // Request path
			param.Request.Proto,                  // HTTP protocol version
			param.StatusCode,                     // HTTP status code
			param.Latency,                        // Request latency
			param.Request.UserAgent(),            // User agent of the client
			param.ErrorMessage,                   // Error message if any
		)
	}))

	// Recovery middleware to handle panics and return 500 Internal Server Error
	router.Use(gin.Recovery())

	// Define a simple GET route
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	// Start the server on port 5000
	router.Run(":5000")
}

// func main() {
// 	f, _ := os.Create("gin.log")
// 	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
// 	router := gin.Default()
// 	router.GET("/ping", func(c *gin.Context) {
// 		c.String(200, "pong")
// 	})
// 	router.Run(":5000")
// }
