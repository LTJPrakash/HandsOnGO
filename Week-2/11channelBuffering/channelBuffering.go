package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("---Channel Buffering---")
	// Create a buffered channel with a capacity of 3
	ch := make(chan string, 3)

	// Start a goroutine that sends multiple values to the channel
	go func() {
		ch <- "task 1"
		ch <- "task 2"
		ch <- "task 3"
		fmt.Println("Sent 3 tasks to the channel")
		ch <- "task 4" // This will block until the main goroutine receives a value
		fmt.Println("Sent task 4 to the channel")
	}()

	// Simulate some delay before receiving
	time.Sleep(2 * time.Second)

	// Receiving values from the channel
	fmt.Println("Received:", <-ch) // Receives "task 1"
	fmt.Println("Received:", <-ch) // Receives "task 2"
	fmt.Println("Received:", <-ch) // Receives "task 3"
	fmt.Println("Received:", <-ch) // Receives "task 4"
}
