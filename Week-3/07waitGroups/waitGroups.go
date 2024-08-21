package main

import (
	"fmt"
	"sync"
	"time"
)

// doWork simulates a worker performing a task.
func doWork(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Mark this goroutine as done when the function completes
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second) // Simulate work with a 1-second sleep
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	fmt.Println("---Wait Groups---")

	var wg sync.WaitGroup // Create a WaitGroup

	numWorkers := 5 // Number of workers to start

	// Add the number of goroutines (workers) to the WaitGroup
	wg.Add(numWorkers)

	// Start the workers as goroutines
	for i := 1; i <= numWorkers; i++ {
		go doWork(i, &wg)
	}

	// Wait for all goroutines in the WaitGroup to finish
	wg.Wait()
	fmt.Println("All workers are done")
}
