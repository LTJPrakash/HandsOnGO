package main

import "github.com/gin-gonic/gin"

// Define a struct to bind and validate URI parameters
type Person struct {
	ID   string `uri:"id" binding:"required,uuid"` // Bind URI parameter "id" and validate it as a required UUID
	Name string `uri:"name" binding:"required"`    // Bind URI parameter "name" and validate it as required
}

func main() {
	// Create a new Gin router with default middleware
	route := gin.Default()

	// Define a route that captures two URI parameters: "name" and "id"
	route.GET("/:name/:id", func(c *gin.Context) {
		var person Person

		// Bind the URI parameters to the Person struct
		if err := c.ShouldBindUri(&person); err != nil {
			// If there's an error (e.g., missing or invalid parameters), return a 400 error
			c.JSON(400, gin.H{"msg": err.Error()})
			return
		}

		// Respond with a JSON object containing the validated name and UUID
		c.JSON(200, gin.H{"name": person.Name, "uuid": person.ID})
	})

	// Start the server on port 5000
	route.Run(":5000")
}
