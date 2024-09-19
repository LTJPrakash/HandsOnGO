package main

import (
	"net/http" // Importing the standard net/http package for HTTP constants like StatusOK

	"github.com/gin-gonic/gin" // Importing the Gin web framework
)

func main() {
	// Create a new Gin router with default middleware (logging and recovery)
	r := gin.Default()

	// Define a GET route at "/someJSON" that returns JSON data in ASCII format
	r.GET("/someJSON", func(c *gin.Context) {
		// Create a map (key-value pair) to store the JSON data to be sent
		data := map[string]interface{}{
			"lang": "GO语言", // This is "Go Language" in Chinese characters
			"tag":  "<br>", // A string that contains an HTML tag
		}

		// Alternatively, you can use gin.H for creating JSON (both are valid)
		// data := gin.H{"message": "Hello, 世界", "tag": "br"}

		// Respond with JSON in ASCII format, which converts non-ASCII characters (like Chinese)
		// into their Unicode representation.
		// For example: {"lang":"GO\u8bed\u8a00","tag":"\u003cbr\u003e"}
		c.AsciiJSON(http.StatusOK, data)
	})

	// Start the server on port 5000
	r.Run(":5000")
}
