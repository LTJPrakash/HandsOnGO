package main

import (
	"fmt"
	"time"
)

// RateLimiter controls the rate of operations using a token bucket algorithm.
type RateLimiter struct {
	tokens    int           // Current number of tokens available
	maxTokens int           // Maximum number of tokens
	rate      time.Duration // Rate at which tokens are replenished
	stop      chan struct{} // Channel to signal when to stop token replenishment
}

// NewRateLimiter creates a new RateLimiter with the specified max tokens and rate.
func NewRateLimiter(maxTokens int, rate time.Duration) *RateLimiter {
	rl := &RateLimiter{
		tokens:    maxTokens,
		maxTokens: maxTokens,
		rate:      rate,
		stop:      make(chan struct{}),
	}

	go rl.refillTokens() // Start the token replenishment process
	return rl
}

// refillTokens replenishes tokens at a specified rate.
func (rl *RateLimiter) refillTokens() {
	ticker := time.NewTicker(rl.rate)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			rl.tokens = rl.maxTokens // Refill tokens
		case <-rl.stop:
			return // Exit the goroutine when stop signal is received
		}
	}
}

// Allow checks if a request can be allowed based on token availability.
func (rl *RateLimiter) Allow() bool {
	if rl.tokens > 0 {
		rl.tokens-- // Use one token
		return true
	}
	return false
}

// Stop signals the RateLimiter to stop replenishing tokens.
func (rl *RateLimiter) Stop() {
	close(rl.stop)
}

func main() {
	fmt.Println("---Rate Limiting---")

	rateLimiter := NewRateLimiter(3, 1*time.Second) // Create a RateLimiter with 5 tokens per second

	for i := 0; i < 10; i++ {
		if rateLimiter.Allow() {
			fmt.Println("Request allowed")
		} else {
			fmt.Println("Request denied")
		}
		time.Sleep(200 * time.Millisecond) // Simulate request interval
	}

	rateLimiter.Stop() // Stop the RateLimiter
}
