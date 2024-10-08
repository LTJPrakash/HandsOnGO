package main

import (
	"net/http"

	"github.com/gin-gonic/gin" // Import the Gin web framework package
)

func main() {
	// Create a new Gin router with default middleware (logging and recovery)
	r := gin.Default()

	// Define a GET route that fetches and serves external data
	r.GET("/someDataFromReader", func(c *gin.Context) {
		// Fetch data (in this case, an image) from an external source (GitHub URL)
		response, err := http.Get("https://images.unsplash.com/profile-1446404465118-3a53b909cc82?ixlib=rb-0.3.5&q=80&fm=jpg&crop=faces&cs=tinysrgb&fit=crop&h=32&w=32&s=a2f8c40e39b8dfee1534eb32acfa6bc7")

		// Handle error cases: if there's an error or the response status isn't OK (200), return Service Unavailable
		if err != nil || response.StatusCode != http.StatusOK {
			c.Status(http.StatusServiceUnavailable)
			return
		}

		// Get the response body (the image data), content length, and content type (e.g., image/png)
		reader := response.Body
		contentLength := response.ContentLength
		contentType := response.Header.Get("Content-Type")

		// Set extra headers, including a Content-Disposition header to suggest downloading the file
		extraHeaders := map[string]string{
			"Content-Disposition": `attachment; filename="gopher.png"`, // Suggests the file name for the download
		}

		// Use DataFromReader to send the fetched data (image) to the client with appropriate headers
		c.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)
	})

	// Start the server on port 5000
	r.Run(":5000") // Listen and serve on 0.0.0.0:5000
}
