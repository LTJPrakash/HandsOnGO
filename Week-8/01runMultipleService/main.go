package main

import (
	"log"      // To log errors or important information
	"net/http" // Provides HTTP server and client implementations
	"time"     // Used for handling timeouts

	"github.com/gin-gonic/gin"   // Gin framework for creating HTTP web servers
	"golang.org/x/sync/errgroup" // Errgroup is used to manage multiple goroutines
)

// Declare an errgroup to handle concurrent operations
var (
	g errgroup.Group
)

// router01 creates the first HTTP router using Gin
func router01() http.Handler {
	e := gin.New()        // Creates a new Gin engine (router)
	e.Use(gin.Recovery()) // Recovery middleware to handle panics and prevent server crashes

	e.GET("/", func(c *gin.Context) { // Define a GET route for the root URL "/"
		c.JSON( // Respond with JSON format
			http.StatusOK, // HTTP status 200 OK
			gin.H{ // gin.H is a shorthand for map[string]interface{}
				"code":    http.StatusOK,
				"message": "Welcome server 01", // Custom message for server 01
			},
		)
	})

	return e // Return the router as an http.Handler interface
}

// router02 creates the second HTTP router using Gin
func router02() http.Handler {
	e := gin.New()        // Another Gin engine for the second server
	e.Use(gin.Recovery()) // Same Recovery middleware to handle panics

	e.GET("/", func(c *gin.Context) { // Define a GET route for server 02
		c.JSON( // Similar response as router01 but for server 02
			http.StatusOK,
			gin.H{
				"code":    http.StatusOK,
				"message": "Welcome server 02", // Custom message for server 02
			},
		)
	})

	return e // Return the router as an http.Handler interface
}

func main() {
	// Define the first server configuration
	server01 := &http.Server{
		Addr:         ":5000",          // Listen on port 5000
		Handler:      router01(),       // Use the first router as the request handler
		ReadTimeout:  5 * time.Second,  // Read timeout to prevent slow client attacks
		WriteTimeout: 10 * time.Second, // Write timeout for the response
	}

	// Define the second server configuration
	server02 := &http.Server{
		Addr:         ":5050",          // Listen on port 5050
		Handler:      router02(),       // Use the second router as the request handler
		ReadTimeout:  5 * time.Second,  // Same read timeout for server 02
		WriteTimeout: 10 * time.Second, // Same write timeout for server 02
	}

	// Start both servers concurrently using errgroup
	g.Go(func() error { // Goroutine for server01
		return server01.ListenAndServe() // Start the server, listen on port 5000
	})

	g.Go(func() error { // Goroutine for server02
		return server02.ListenAndServe() // Start the server, listen on port 5050
	})

	// Wait for both servers to start; if any server fails, log the error
	if err := g.Wait(); err != nil { // Wait() blocks until both goroutines return
		log.Fatal(err) // Log and exit if there's any error
	}
}
