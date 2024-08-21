package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("---Atomic Counters---")

	var counter int64 // Shared counter variable, using int64 for atomic operations

	var wg sync.WaitGroup
	numGoroutines := 10           // Number of goroutines to start
	incrementPerGoroutine := 1000 // Number of increments per goroutine

	wg.Add(numGoroutines) // Set the number of goroutines to wait for

	// Start the goroutines
	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done() // Decrement the WaitGroup counter when this goroutine finishes
			for j := 0; j < incrementPerGoroutine; j++ {
				atomic.AddInt64(&counter, 1) // Atomically increment the counter
			}
		}()
	}

	wg.Wait() // Wait for all goroutines to complete

	// Load and print the final value of the counter atomically
	fmt.Printf("Final counter value: %d\n", atomic.LoadInt64(&counter))
}
