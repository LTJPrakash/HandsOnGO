package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("---Mutexes---")

	var mu sync.Mutex // Mutex to control access to the shared resource
	var counter int   // Shared counter variable

	var wg sync.WaitGroup
	numGoroutines := 10           // Number of goroutines to start
	incrementPerGoroutine := 1000 // Number of increments per goroutine

	wg.Add(numGoroutines) // Set the number of goroutines to wait for

	// Start the goroutines
	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done() // Decrement the WaitGroup counter when this goroutine finishes
			for j := 0; j < incrementPerGoroutine; j++ {
				mu.Lock()   // Acquire the mutex lock before accessing the critical section
				counter++   // Critical section: increment the counter
				mu.Unlock() // Release the mutex lock after accessing the critical section
			}
		}()
	}

	wg.Wait() // Wait for all goroutines to complete

	// Print the final value of the counter
	fmt.Printf("Final counter value: %d\n", counter)
}
