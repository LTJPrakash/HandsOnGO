package main

import "fmt"

func main() {
	fmt.Println("---Closing Channels---")

	// Create a channel to communicate integers
	ch := make(chan int)

	// Start a goroutine to send data into the channel
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i // Send the value of i into the channel
		}
		close(ch) // Close the channel after sending all data
	}()

	// Receiving from the channel until it's closed
	for val := range ch {
		fmt.Println(val)
	}

	fmt.Println("Channel closed and all data received.")
}
