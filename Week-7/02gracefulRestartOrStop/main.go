//go:build go1.8
// +build go1.8

package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create a new Gin router with default middleware
	router := gin.Default()

	// Define a simple route to simulate a long-running task
	router.GET("/", func(c *gin.Context) {
		// Simulate a long task by sleeping for 5 seconds
		time.Sleep(5 * time.Second)
		c.String(http.StatusOK, "Welcome Gin Server")
	})

	// Create an HTTP server using Gin router, listening on port 5000
	srv := &http.Server{
		Addr:    ":5000",          // Server listens on localhost:5000
		Handler: router.Handler(), // Use the Gin router as the handler
	}

	// Start the server in a separate Goroutine so that it doesn’t block
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err) // Log an error if the server fails to start
		}
	}()

	// Channel to listen for system interrupt or termination signals
	quit := make(chan os.Signal, 1)
	// Capture SIGINT (Ctrl+C) and SIGTERM signals
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Block until we receive an interrupt signal
	<-quit
	log.Println("Shutdown Server ...")

	// Create a context with a timeout of 5 seconds for graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Attempt to gracefully shutdown the server
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err) // Log any errors during shutdown
	}

	// Wait for the context's timeout or completion of active requests
	select {
	case <-ctx.Done():
		log.Println("Timeout of 5 seconds.")
	}

	// Log that the server is exiting
	log.Println("Server exiting")
}
