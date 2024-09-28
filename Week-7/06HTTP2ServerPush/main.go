package main

import (
	"html/template"
	"log"

	"github.com/gin-gonic/gin" // Import the Gin framework
)

// Define an HTML template
var html = template.Must(template.New("https").Parse(`
<html>
<head>
  <title>Https Test</title>
  <script src="/assets/app.js"></script>
</head>
<body>
  <h1 style="color:red;">Welcome, Ginner!</h1>
</body>
</html>
`))

func main() {
	r := gin.Default() // Initialize the Gin router with default middleware (logging and recovery)

	// Serve static files from the "./assets" directory, such as JavaScript files
	r.Static("/assets", "./assets")

	// Set the HTML template that we defined above
	r.SetHTMLTemplate(html)

	// Define a GET route for the root "/"
	r.GET("/", func(c *gin.Context) {
		// Check if HTTP/2 Push is supported
		if pusher := c.Writer.Pusher(); pusher != nil {
			// Use the pusher.Push() method to push the "app.js" file to the client
			if err := pusher.Push("/assets/app.js", nil); err != nil {
				log.Printf("Failed to push: %v", err)
			}
		}
		// Serve the HTML template with a status message
		c.HTML(200, "https", gin.H{
			"status": "success",
		})
	})

	// Run the server using HTTPS on port 5000
	// Use self-signed certificate files: server.pem (certificate) and server.key (private key)
	r.RunTLS(":5000", "./testdata/server.pem", "./testdata/server.key")
}
