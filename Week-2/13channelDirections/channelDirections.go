package main

import (
	"fmt"
)

// sendValues is a function that only sends values into the channel
func sendValues(ch chan<- int) {
	for i := 1; i <= 5; i++ {
		ch <- i
	}
	close(ch) // Close the channel when done sending
}

// receiveValues is a function that only receives values from the channel
func receiveValues(ch <-chan int) {
	for value := range ch {
		fmt.Println("Received:", value)
	}
}

func main() {
	fmt.Println("---Channel Directions---")

	ch := make(chan int) // Create an unbuffered channel

	// Start a goroutine to send values
	go sendValues(ch)

	// Receive values in the main goroutine
	receiveValues(ch)
}
