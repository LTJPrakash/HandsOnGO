package main

import "fmt"

// getId returns a closure that generates a unique integer each time it is called.
func getId() func() int {
	// Initialize a local variable 'i' to keep track of the unique ID.
	i := 0

	// Return an anonymous function that increments 'i' and returns the updated value.
	return func() int {
		i++      // Increment 'i' each time the closure is called.
		return i // Return the current value of 'i'.
	}
}

func main() {
	fmt.Println("---Closures---")

	// Call getId to obtain a closure that generates unique integers.
	result := getId()

	// Call the closure multiple times and print the results.
	fmt.Println("First call result: ", result())  // Output: 1
	fmt.Println("Second call result: ", result()) // Output: 2
	fmt.Println("Third call result: ", result())  // Output: 3
}
