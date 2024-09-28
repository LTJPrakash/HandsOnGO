package main

import (
	"net/http"

	"github.com/gin-gonic/gin" // Import the Gin web framework package
)

func main() {
	// Initialize the Gin router with default middleware (logging and recovery)
	r := gin.Default()

	// Define a route for JSONP requests
	r.GET("/JSONP", func(c *gin.Context) {
		data := map[string]interface{}{
			"foo": "bar",
		}

		// Retrieve the callback parameter from the query string
		// callback := c.Query("callback")

		// Respond with JSONP format using the callback parameter
		c.JSONP(http.StatusOK, data)
	})

	// Start the server on port 5000
	r.Run(":5000") // Listen and serve on 0.0.0.0:5000
}
