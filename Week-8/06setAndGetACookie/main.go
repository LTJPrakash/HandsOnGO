package main

import (
	"fmt"

	"github.com/gin-gonic/gin" // Import the Gin web framework package
)

func main() {
	// Create a new Gin router with default middleware (logging and recovery)
	r := gin.Default()

	// Define a GET route for "/cookie"
	r.GET("/cookie", func(c *gin.Context) {
		// Try to retrieve the cookie "gin_cookie" from the request
		cookie, err := c.Cookie("gin_cookie")

		if err != nil {
			// If the cookie doesn't exist, set a default value
			cookie = "NotSet"
			// Set a new cookie with name "gin_cookie", value "test", expiration of 3600 seconds (1 hour), and other parameters
			c.SetCookie("gin_cookie", "test", 3600, "/", "localhost", false, true)
		}

		// Print the cookie value to the console for debugging
		fmt.Printf("Cookie value: %s \n", cookie)
	})

	// Start the server on port 5000
	r.Run(":5000") // Listen and serve on 0.0.0.0:5000
}
