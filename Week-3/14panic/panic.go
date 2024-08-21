package main

import "fmt"

func main() {
	fmt.Println("---Panic---")

	fmt.Println("Starting the program")

	// Trigger a panic with a custom message
	panic("This is a panic message")

	// This line will never be executed because of the panic
	fmt.Println("This line will not be executed")
}
