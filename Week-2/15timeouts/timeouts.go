package main

import (
	"fmt"
	"time"
)

// fetchData simulates a long-running operation
func fetchData(result chan string) {
	// Simulate a delay in fetching data
	time.Sleep(3 * time.Second)
	result <- "Data fetched successfully"
}

func main() {
	fmt.Println("---Timeouts---")

	result := make(chan string)

	// Start the fetch operation in a goroutine
	go fetchData(result)

	// Use select to implement a timeout
	select {
	case res := <-result:
		fmt.Println(res) // If fetchData completes, print the result
	case <-time.After(3 * time.Second):
		fmt.Println("Timeout! Fetch operation took too long.")
	}
}
