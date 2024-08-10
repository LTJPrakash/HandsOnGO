package main

import "fmt"

// Demonstrates the usage of pointers in Go.
func main() {
	fmt.Println("---Pointers---")

	// Declare a pointer variable without initializing it.
	var ptr *int
	fmt.Println("Ptr:", ptr) // The value of an uninitialized pointer is nil.

	// Initialize an integer variable.
	myNum := 17

	// Create a pointer that holds the address of 'myNum'.
	var ptr1 = &myNum
	fmt.Println("Address of myNum:", &myNum)               // Print the address of 'myNum'.
	fmt.Println("Value of ptr1 (address of myNum):", ptr1) // Print the address stored in 'ptr1'.
	fmt.Println("Value pointed to by ptr1:", *ptr1)        // Print the value at the address stored in 'ptr1' (which is 'myNum').

	// Modify the value of 'myNum' through the pointer 'ptr1'.
	*ptr1 = *ptr1 * 2
	fmt.Println("New value of myNum:", myNum) // Print the updated value of 'myNum'.
}
