package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin" // Import the Gin web framework package
)

func main() {
	// Create a new Gin router with default middleware (logging and recovery)
	r := gin.Default()

	// Set a lower memory limit for multipart forms (default is 32 MiB)
	r.MaxMultipartMemory = 8 << 20 // 8 MiB

	// Handle file upload
	r.POST("/upload", func(c *gin.Context) {
		// Get the file from the form input
		file, err := c.FormFile("file")
		if err != nil {
			log.Println("Error getting file:", err)
			c.String(http.StatusBadRequest, "Error getting file")
			return
		}

		// Log the file name
		log.Println("Uploaded File:", file.Filename)

		// Specify the destination to save the uploaded file
		dst := filepath.Join("uploads", file.Filename) // Define the destination path

		// Save the uploaded file to the specified path
		if err := c.SaveUploadedFile(file, dst); err != nil {
			log.Println("Error saving file:", err)
			c.String(http.StatusInternalServerError, "Error saving file")
			return
		}

		// Respond with a success message
		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})

	// Start the server on port 5000
	r.Run(":5000") // Listen and serve on 0.0.0.0:5000
}
