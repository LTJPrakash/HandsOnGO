package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create a new Gin router with default middleware (logging and recovery)
	router := gin.Default()

	// Custom HTTP server configuration
	s := &http.Server{
		Addr:           ":5000",          // Server will listen on port 5000
		Handler:        router,           // Use the Gin router as the HTTP handler
		ReadTimeout:    10 * time.Second, // Maximum duration for reading the request
		WriteTimeout:   10 * time.Second, // Maximum duration for writing the response
		MaxHeaderBytes: 1 << 20,          // Maximum size for request headers (1 MB)
	}

	// Start the server with the custom configuration
	s.ListenAndServe()
}
