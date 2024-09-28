package main

import (
	"github.com/gin-gonic/gin" // Import the Gin web framework package
)

func main() {
	// Initialize the Gin router with default middleware (logging and recovery)
	r := gin.Default()

	// Route to serve JSON with Unicode entities (escaped HTML)
	r.GET("/json", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"html": "<b>Hello, world!</b>", // HTML is escaped
		})
	})

	// Route to serve pure JSON (literal characters)
	r.GET("/purejson", func(c *gin.Context) {
		c.PureJSON(200, gin.H{
			"html": "<b>Hello, world!</b>", // HTML is not escaped
		})
	})

	// Start the server on port 5000
	r.Run(":5000") // Listen and serve on 0.0.0.0:5000
}
