package main

import (
	"fmt"
)

func task(i int) {
	fmt.Println("task ", i)
}

func main() {
	fmt.Println("---Channels---")

	// Creating a channel of type int
	ch := make(chan int)

	// Start a goroutine that sends a value to the channel
	go func() {
		ch <- 42
	}()

	// Receive the value from the channel in the main goroutine
	value := <-ch
	fmt.Println("Received from channel:", value) // Output: 42

	// Uncomment the following lines to see how goroutines work with channels
	/*
		for i := 1; i <= 10; i++ {
			go task(i)
		}
	*/
}
