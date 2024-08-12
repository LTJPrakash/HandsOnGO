package main

import (
	"fmt"
	"time"
)

func task(id int, done chan bool) {
	fmt.Printf("task %d starting\n", id)
	time.Sleep(time.Second) // Simulate task by sleeping for 1 second
	fmt.Printf("task %d done\n", id)
	done <- true // Signal that the task is done
}

func main() {
	fmt.Println("---Channel Synchronization---")

	const numOfTask = 3
	done := make(chan bool, numOfTask) // Buffered channel with capacity equal to number of tasks

	// Start each task in a separate goroutine
	for i := 1; i <= numOfTask; i++ {
		go task(i, done)
	}

	// Wait for all workers to complete
	for i := 1; i <= numOfTask; i++ {
		<-done // Receive a signal for each completed task
	}

	fmt.Println("All tasks completed")
}
