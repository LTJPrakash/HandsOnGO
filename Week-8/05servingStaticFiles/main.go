package main

import (
	"net/http"

	"github.com/gin-gonic/gin" // Import the Gin web framework package
)

func main() {
	// Create a new Gin router with default middleware (logging and recovery)
	r := gin.Default()

	// Serve static files from the "./assets" directory under the "/assets" route
	r.Static("/assets", "./assets")

	// Serve files from a custom file system at the "/more_static" route
	r.StaticFS("/more_static", http.Dir("my_file_system"))

	// Serve a single static file (favicon) at the "/favicon.ico" route
	r.StaticFile("/favicon.ico", "./resources/favicon.ico")

	// Start the server on port 5000
	r.Run(":5000") // Listen and serve on 0.0.0.0:5000
}
