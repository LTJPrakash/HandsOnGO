package main

import "fmt"

func main() {
	fmt.Println("---Non Blocking Channel Operations---")
	// Create a channel for string messages
	messages := make(chan string)

	// Create a channel for boolean signals
	signals := make(chan bool)

	// Non-blocking receive: Try to receive from messages channel
	select {
	case msg := <-messages:
		fmt.Println("Received message:", msg)
	default:
		fmt.Println("No message received")
	}

	// Prepare a message to send
	msg := "hi"

	// Non-blocking send: Try to send a message into messages channel
	select {
	case messages <- msg:
		fmt.Println("Sent message:", msg)
	default:
		fmt.Println("No message sent")
	}

	// Non-blocking multi-way select: Check both channels for activity
	select {
	case msg := <-messages:
		fmt.Println("Received message:", msg)
	case sig := <-signals:
		fmt.Println("Received signal:", sig)
	default:
		fmt.Println("No activity")
	}
}
