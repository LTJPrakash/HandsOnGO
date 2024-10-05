package main

import (
	"log"

	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin" // Import the Gin web framework package
)

func main() {
	// Create a new Gin router with default middleware (logging and recovery)
	r := gin.Default()

	// Ping handler
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	// Start the server with automatic TLS for the specified domains
	log.Fatal(autotls.Run(r, "example1.com", "example2.com"))

	// Alternative setup using autocert.Manager (commented out)
	// m := autocert.Manager{
	// 	Prompt:     autocert.AcceptTOS,
	// 	HostPolicy: autocert.HostWhitelist("example1.com", "example2.com"),
	// 	Cache:      autocert.DirCache("/var/www/.cache"),
	// }

	// log.Fatal(autotls.RunWithManager(r, &m))
}
