package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("---Timers---")

	// Create a timer that fires after 5 seconds
	timer := time.NewTimer(5 * time.Second)
	fmt.Println("Timer started, waiting for 5 seconds...")

	// Start a goroutine to stop the timer after 2 seconds
	go func() {
		time.Sleep(2 * time.Second)
		if stop := timer.Stop(); stop {
			fmt.Println("Timer stopped before expiration")
		}
	}()

	// Wait for the timer to expire or be stopped
	select {
	case <-timer.C:
		fmt.Println("Timer expired")
	case <-time.After(3 * time.Second):
		fmt.Println("Timer was stopped, waiting to reset")
	}

	// Reset the timer to fire after 4 seconds
	if !timer.Stop() {
		<-timer.C // Drain the channel if the timer has already fired
	}
	timer.Reset(4 * time.Second)
	fmt.Println("Timer reset, waiting for 4 seconds...")

	<-timer.C
	fmt.Println("Timer expired after reset")

	// Using time.After for a one-time delayed operation
	fmt.Println("Waiting for 2 seconds using time.After...")
	<-time.After(2 * time.Second)
	fmt.Println("2 seconds have passed (time.After)")
}
