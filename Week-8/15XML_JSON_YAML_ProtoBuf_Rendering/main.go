package main

import (
	"net/http"

	"testdata/protoexample" // Import your protobuf generated file

	"github.com/gin-gonic/gin" // Import the Gin web framework package
)

func main() {
	// Create a default Gin router with logging and recovery middleware
	r := gin.Default()

	// Endpoint returning JSON response
	r.GET("/someJSON", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})

	// Endpoint returning JSON response using a struct
	r.GET("/moreJSON", func(c *gin.Context) {
		var msg struct {
			Name    string `json:"user"` // "user" will be the key in the JSON response
			Message string
			Number  int
		}
		msg.Name = "Lena"
		msg.Message = "hey"
		msg.Number = 123
		c.JSON(http.StatusOK, msg) // Will output: {"user": "Lena", "Message": "hey", "Number": 123}
	})

	// Endpoint returning XML response
	r.GET("/someXML", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})

	// Endpoint returning YAML response
	r.GET("/someYAML", func(c *gin.Context) {
		c.YAML(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})

	// Endpoint returning Protocol Buffers response
	r.GET("/someProtoBuf", func(c *gin.Context) {
		reps := []int64{1, 2} // Example data for protobuf
		label := "test"
		data := &protoexample.Test{ // Assuming protoexample.Test is your protobuf message type
			Label: &label,
			Reps:  reps,
		}

		// Serializing the data to protobuf format
		c.ProtoBuf(http.StatusOK, data) // Returns serialized protobuf data
	})

	// Start the server on port 5000
	r.Run(":5000") // Listen and serve on 0.0.0.0:5000
}
