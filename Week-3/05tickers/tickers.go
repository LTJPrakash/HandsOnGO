package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("---Tickers---")

	// Create a ticker that ticks every 1 second
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop() // Ensure the ticker is stopped when done

	// Channel to signal when to stop the ticker
	done := make(chan bool)

	// Start a goroutine to handle the ticks
	go func() {
		for {
			select {
			case t := <-ticker.C: // Receive the time from the ticker channel
				fmt.Println("Tick at", t)
			case <-done: // Listen for the signal to stop the ticker
				fmt.Println("Ticker stopped")
				return
			}
		}
	}()

	// Let the ticker run for 5 ticks, then stop it
	time.Sleep(5 * time.Second)
	done <- true

	// Additional delay to allow all prints to complete
	time.Sleep(1 * time.Second)
	fmt.Println("Program finished")
}
