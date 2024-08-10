package main

import (
	"errors"
	"fmt"
)

// divide performs integer division and returns an error if division by zero is attempted.
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero") // Return an error for division by zero.
	}
	return a / b, nil // Return the result and a nil error if division is successful.
}

// CustomError is a struct that implements the error interface.
type CustomError struct {
	Code    int    // Error code.
	Message string // Error message.
}

// Error() method for CustomError to satisfy the error interface.
func (e *CustomError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.Code, e.Message) // Format the error message.
}

// riskyOperation simulates a function that returns a custom error.
func riskyOperation() error {
	return &CustomError{Code: 404, Message: "Resource not found"} // Return a custom error.
}

func main() {
	fmt.Println("---Errors And Custom Errors---")

	// Uncomment this section to test the divide function and error handling.
	// result, err := divide(10, 0)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }
	// fmt.Println("Result:", result)

	// Call riskyOperation and handle the returned error.
	erre := riskyOperation()
	if erre != nil {
		fmt.Println("An error occurred:", erre) // Print the error message.
		// Type assertion to access specific fields of CustomError.
		if customErr, ok := erre.(*CustomError); ok {
			fmt.Println("Error Code:", customErr.Code) // Print the custom error code.
		}
	}
}
