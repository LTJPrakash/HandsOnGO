package main

import (
	"fmt"
	"time"
)

// f is a function that prints a message along with an index three times.
func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

// firstProcess simulates a process that takes 2 seconds to complete.
func firstProcess() {
	fmt.Println("First process start")
	time.Sleep(2 * time.Second) // Simulating a delay of 2 seconds
	fmt.Println("First process end")
}

// secondProcess simulates a process that would typically take 1 second to complete.
func secondProcess() {
	fmt.Println("Second process start")
	// time.Sleep(1 * time.Second) // Uncomment this line to simulate a delay of 1 second
	fmt.Println("Second process end")
}

func main() {
	fmt.Println("---Goroutines---")

	// Starting firstProcess in a new goroutine
	go firstProcess()

	// Starting secondProcess in another goroutine
	go secondProcess()

	// Main function sleeps for 900 milliseconds to allow goroutines to start
	time.Sleep(900 * time.Millisecond)

	// Calling f function directly in the main thread
	f("direct")

	// Starting f function in a new goroutine
	go f("goroutine")

	// Starting an anonymous function as a goroutine
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// Allowing time for all goroutines to complete
	time.Sleep(3 * time.Second)
	fmt.Println("Execution completed")
}
