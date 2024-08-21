package main

import "fmt"

func main() {
	fmt.Println("---Recover---")

	fmt.Println("Starting main function")

	// Defer function to handle any panic that occurs in main
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main:", r)
		}
	}()

	// Call a function that triggers a panic
	performTask()

	fmt.Println("End of main function")
}

func performTask() {
	fmt.Println("Starting performTask function")

	// Defer function to handle any panic that occurs in performTask
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in performTask:", r)
			// Optionally re-panic if you want to propagate the panic further
			// panic(r)
		}
	}()

	// Call a function that will trigger a panic
	causePanic()

	fmt.Println("End of performTask function")
}

func causePanic() {
	fmt.Println("Starting causePanic function")

	// Trigger a panic
	panic("Something went wrong in causePanic!")

	// This line will not be executed because of the panic
	fmt.Println("This line will not be executed in causePanic")
}
