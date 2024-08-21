package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("---Defer---")

	fmt.Println("Starting main function")

	// Open a file for writing (create if it doesn't exist)
	file, err := os.Create("example.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}

	// Ensure the file is closed when the function exits
	defer func() {
		fmt.Println("Closing file")
		if err := file.Close(); err != nil {
			fmt.Println("Error closing file:", err)
		}
	}()

	// Write a string to the file
	_, err = file.WriteString("Hello, World!")
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("Finished writing to file")
}
