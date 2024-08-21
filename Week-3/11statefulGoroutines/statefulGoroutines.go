package main

import (
	"fmt"
	"sync"
	"time"
)

// Counter maintains a count with thread-safe operations.
type Counter struct {
	mu    sync.Mutex // Mutex to synchronize access to the counter
	value int        // The counter value
}

// Increment safely increments the counter value.
func (c *Counter) Increment() {
	c.mu.Lock()         // Acquire the lock
	defer c.mu.Unlock() // Release the lock after the function returns
	c.value++           // Increment the counter
}

// Reset safely resets the counter value to zero.
func (c *Counter) Reset() {
	c.mu.Lock()         // Acquire the lock
	defer c.mu.Unlock() // Release the lock after the function returns
	c.value = 0         // Reset the counter
}

// Value returns the current counter value safely.
func (c *Counter) Value() int {
	c.mu.Lock()         // Acquire the lock
	defer c.mu.Unlock() // Release the lock after the function returns
	return c.value      // Return the current value
}

// counterWorker simulates work by incrementing the counter multiple times.
func counterWorker(c *Counter, id int, wg *sync.WaitGroup) {
	defer wg.Done() // Decrement the WaitGroup counter when this goroutine finishes
	for i := 0; i < 10; i++ {
		c.Increment() // Safely increment the counter
		fmt.Printf("Worker %d incremented counter\n", id)
		time.Sleep(100 * time.Millisecond) // Simulate work
	}
}

func main() {
	fmt.Println("---Stateful Goroutines---")

	var wg sync.WaitGroup
	counter := &Counter{} // Create a new Counter instance

	wg.Add(2)                         // Set the number of goroutines to wait for
	go counterWorker(counter, 1, &wg) // Start the first worker goroutine
	go counterWorker(counter, 2, &wg) // Start the second worker goroutine

	wg.Wait() // Wait for all worker goroutines to finish

	// Print the final counter value
	fmt.Printf("Final counter value: %d\n", counter.Value())
}
