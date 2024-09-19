package main

import (
	"log"
	"time" // Import the time package to handle date and time formats

	"github.com/gin-gonic/gin" // Import the Gin web framework
)

// Define a struct to hold the query parameters
// The `time_format` tag specifies the expected format for the birthday field
// The `time_utc` tag indicates that the birthday should be treated in UTC
type Person struct {
	Name     string    `form:"name"`                                           // Bind query parameter "name" to Name field
	Address  string    `form:"address"`                                        // Bind query parameter "address" to Address field
	Birthday time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"` // Bind birthday in "YYYY-MM-DD" format
}

func main() {
	// Create a new Gin router with default middleware (logging and recovery)
	route := gin.Default()

	// Define a GET route at "/testing" which handles query parameters
	route.GET("/testing", startPage)

	// Start the server on port 5000
	route.Run(":5000")
}

// Handler function for the "/testing" endpoint
func startPage(c *gin.Context) {
	var person Person

	// Bind the query parameters from the request to the Person struct
	// If the binding is successful, it will return nil (no error)
	if c.ShouldBind(&person) == nil {
		// Log the bound values to the console for debugging purposes
		log.Println("Name:", person.Name)
		log.Println("Address:", person.Address)
		log.Println("Birthday:", person.Birthday)
	}

	// Return a success message to the client
	c.String(200, "Success")
}
