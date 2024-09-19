package main

import "github.com/gin-gonic/gin"

// Define a struct to bind the form data for subscription
// The `form:"subscribe"` tag maps the form field "subscribe" to the struct
type SubscriptionForm struct {
	Subscribe bool `form:"subscribe"` // This will bind a boolean value from the form
}

func main() {
	// Create a new Gin router with default middleware (logging and recovery)
	r := gin.Default()

	// Define a POST route at "/subscribe" to handle subscription form submissions
	r.POST("/subscribe", func(c *gin.Context) {
		var form SubscriptionForm

		// Bind the incoming form data to the SubscriptionForm struct
		// If there's an error (e.g., form field is missing or invalid), return a 400 error
		if err := c.ShouldBind(&form); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		// Prepare a message based on the value of the "subscribe" field
		message := "Not Subscribed!" // Default message
		if form.Subscribe {
			message = "Subscribed!" // Change message if the user opted to subscribe
		}

		// Respond with the message as JSON
		c.JSON(200, gin.H{"message": message})
	})

	// Start the server on port 5000
	r.Run(":5000")
}
