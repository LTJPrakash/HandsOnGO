package main

import "fmt"

// sumTillN calculates the sum of all integers from 1 to n using recursion.
func sumTillN(n int) int {
	// Base case: If n is 1, return 1 (the sum of numbers from 1 to 1 is 1).
	if n == 1 {
		return 1
	}

	// Recursive case: Add the current number n to the result of sumTillN(n-1).
	return n + sumTillN(n-1)
}

func main() {
	fmt.Println("---Recursion---")

	var n int
	fmt.Printf("Enter a number: ")
	fmt.Scan(&n) // Read an integer input from the user.

	// Calculate and print the sum of all integers from 1 to n.
	fmt.Println("Sum of numbers from 1 to", n, "is:", sumTillN(n))
}
