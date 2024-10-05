package main

import (
	"net/http"

	"github.com/gin-gonic/gin" // Import the Gin web framework package
)

// Define the structure for formA
type formA struct {
	Foo string `json:"foo" xml:"foo" binding:"required"` // The Foo field is required
}

// Define the structure for formB
type formB struct {
	Bar string `json:"bar" xml:"bar" binding:"required"` // The Bar field is required
}

// SomeHandler processes incoming requests
func SomeHandler(c *gin.Context) {
	objA := formA{}
	objB := formB{}

	// Attempt to bind the request body to objA
	if errA := c.ShouldBind(&objA); errA == nil {
		c.String(http.StatusOK, `the body should be formA`)
		// If objA binding is successful, respond with a message
	} else if errB := c.ShouldBind(&objB); errB == nil {
		// If objA binding fails, attempt to bind to objB
		c.String(http.StatusOK, `the body should be formB`)
	} else {
		// Handle the error if both bindings fail
		c.String(http.StatusBadRequest, "Invalid request body")
	}
}

func main() {
	// Create a new Gin router with default middleware (logging and recovery)
	r := gin.Default()

	// Register the SomeHandler for the desired route
	r.POST("/someendpoint", SomeHandler)

	// Start the server on port 5000
	r.Run(":5000") // Listen and serve on 0.0.0.0:5000
}
