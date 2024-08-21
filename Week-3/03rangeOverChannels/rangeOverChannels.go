package main

import "fmt"

func main() {
	fmt.Println("---Range Over Channels---")

	// Create a channel to send integer values
	ch := make(chan int)

	// Start a goroutine to send data into the channel
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i // Send the value of i to the channel
			fmt.Println("Sent:", i)
		}
		close(ch) // Close the channel to signal completion
		fmt.Println("Channel closed in goroutine.")
	}()

	// Range over the channel to receive values until the channel is closed
	fmt.Println("Receiving from channel:")
	for val := range ch {
		fmt.Println("Received:", val)
	}

	fmt.Println("Channel closed and iteration finished.")
}
