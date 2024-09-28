package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

// Person struct to hold the name and address fields
type Person struct {
	Name    string `form:"name"`    // Bind query parameter "name"
	Address string `form:"address"` // Bind query parameter "address"
}

func main() {
	// Initialize the Gin router with default middleware (logging and recovery)
	route := gin.Default()

	// Define a route that accepts any HTTP method on "/testing"
	route.Any("/testing", startPage)

	// Start the server on port 5000
	route.Run(":5000") // Listen and serve on 0.0.0.0:5000
}

// startPage handles the incoming request and binds query parameters to the Person struct
func startPage(c *gin.Context) {
	var person Person

	// Bind query parameters to the person struct
	if c.ShouldBindQuery(&person) == nil {
		// Log the extracted values
		log.Println("====== Only Bind By Query String ======")
		log.Println("Name:", person.Name)
		log.Println("Address:", person.Address)
	}

	// Respond with a success message
	c.String(200, "Success")
}
