package main

import (
	"fmt"
	"time"
)

// Task represents a unit of work to be processed by a worker.
type Task struct {
	ID int
}

// Worker represents a worker that processes tasks.
func worker(id int, tasks <-chan Task, results chan<- string) {
	for task := range tasks {
		// Simulate task processing
		time.Sleep(time.Second) // Simulate a 1-second task processing time
		results <- fmt.Sprintf("Worker %d processed task %d", id, task.ID)
	}
}

func main() {
	fmt.Println("---Worker Pools---")

	const numWorkers = 3 // Number of workers in the pool
	const numTasks = 10  // Total number of tasks to process

	// Create channels for tasks and results
	tasks := make(chan Task, numTasks)     // Buffered channel to hold tasks
	results := make(chan string, numTasks) // Buffered channel to hold results

	// Start the worker pool
	for i := 1; i <= numWorkers; i++ {
		go worker(i, tasks, results) // Start each worker as a goroutine
	}

	// Send tasks to the task channel
	for i := 1; i <= numTasks; i++ {
		tasks <- Task{ID: i} // Create a task with an ID and send it to the channel
	}
	close(tasks) // Close the task channel to signal that no more tasks will be sent

	// Collect results from the workers
	for i := 1; i <= numTasks; i++ {
		fmt.Println(<-results) // Print each result as it's received
	}
}
