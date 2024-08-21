package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("---String Formatting---")

	name := "Alice"
	age := 30
	height := 5.8
	isStudent := false

	// Formatting integers and strings
	fmt.Printf("Name: %s\n", name) // Format a string
	fmt.Printf("Age: %d\n", age)   // Format an integer

	// Formatting floating-point numbers
	fmt.Printf("Height: %.1f meters\n", height) // Format a float with one decimal place
	fmt.Printf("Pi: %f\n", math.Pi)             // Format a float with default precision

	// Boolean formatting
	fmt.Printf("Is student: %t\n", isStudent) // Format a boolean

	// Formatting with default representation
	fmt.Printf("Default format: %v\n", height) // Default format for a value

	// Formatting with type representation
	fmt.Printf("Type of height: %T\n", height) // Print the type of the variable

	// Formatting complex types
	person := struct {
		Name string
		Age  int
	}{
		Name: "Bob",
		Age:  25,
	}
	fmt.Printf("Person: %+v\n", person) // Print a struct with field names

	// Using Sprintf to create formatted strings
	result := fmt.Sprintf("Hello, %s! You are %d years old.", name, age)
	fmt.Println(result) // Output formatted string
}
