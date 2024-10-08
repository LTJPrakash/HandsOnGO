package main

import (
	"net/http"

	"github.com/gin-gonic/gin" // Import the Gin web framework package
	"github.com/gin-gonic/gin/binding"
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
	// This reads c.Request.Body and stores the result into the context.
	if errA := c.ShouldBindBodyWith(&objA, binding.JSON); errA == nil {
		c.String(http.StatusOK, `the body should be formA`)
		// At this time, it reuses body stored in the context.
	} else if errB := c.ShouldBindBodyWith(&objB, binding.JSON); errB == nil {
		c.String(http.StatusOK, `the body should be formB JSON`)
		// And it can accepts other formats
	} else {
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
