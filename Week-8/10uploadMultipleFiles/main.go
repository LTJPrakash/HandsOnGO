package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin" // Import the Gin web framework package
)

func main() {
	// Create a new Gin router with default middleware (logging and recovery)
	r := gin.Default()

	// Set a lower memory limit for multipart forms (default is 32 MiB)
	r.MaxMultipartMemory = 8 << 20 // 8 MiB

	// Create a directory for uploads if it doesn't exist
	if err := os.MkdirAll("uploads", os.ModePerm); err != nil {
		log.Fatalf("Failed to create uploads directory: %v", err)
	}

	// Handle file uploads
	r.POST("/upload", func(c *gin.Context) {
		// Get the multipart form
		form, err := c.MultipartForm() // Correctly call the method and capture both return values
		if err != nil {
			log.Println("Error retrieving multipart form:", err)
			c.String(http.StatusBadRequest, "Error retrieving form data")
			return
		}

		files := form.File["upload[]"] // Expecting an array of files

		// Check if any files were uploaded
		if len(files) == 0 {
			c.String(http.StatusBadRequest, "No files uploaded")
			return
		}

		// Iterate over the files and save each one
		for _, file := range files {
			log.Println("Uploading file:", file.Filename)

			// Define the destination path for each uploaded file
			dst := filepath.Join("uploads", file.Filename)

			// Save the uploaded file to the specified path
			if err := c.SaveUploadedFile(file, dst); err != nil {
				log.Println("Error saving file:", err)
				c.String(http.StatusInternalServerError, fmt.Sprintf("Error saving file: %s", file.Filename))
				return
			}
		}

		// Respond with the number of files uploaded
		c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
	})

	// Start the server on port 5000
	// By default, it listens on localhost:5000, which can be accessed by your browser or API clients
	r.Run(":5000") // Listen and serve on 0.0.0.0:5000
}
