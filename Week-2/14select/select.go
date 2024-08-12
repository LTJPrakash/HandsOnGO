package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("---Select---")

	ch1 := make(chan string)
	ch2 := make(chan string)

	// Goroutine that sends messages to ch1 every second
	go func() {
		for {
			time.Sleep(1 * time.Second)
			ch1 <- "message from ch1"
		}
	}()

	// Goroutine that sends messages to ch2 every two seconds
	go func() {
		for {
			time.Sleep(2 * time.Second)
			ch2 <- "message from ch2"
		}
	}()

	// Receive messages from either channel using select
	for i := 0; i < 5; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println("Received:", msg1)
		case msg2 := <-ch2:
			fmt.Println("Received:", msg2)
		}
	}
}
