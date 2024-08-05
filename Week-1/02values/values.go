package main

import "fmt"

func main() {
	fmt.Println("---Values---")

	// Strings, which can be concatenated using the + operator.
	fmt.Println("go" + "lang")

	// Another example of string concatenation
	firstName := "John"
	lastName := "Doe"
	fullName := firstName + " " + lastName
	fmt.Println("Full Name:", fullName)

	// Integers and floats.
	// Integer addition
	fmt.Println("1+1 =", 1+1)

	// Floating-point division
	fmt.Println("7.0/3.0 =", 7.0/3.0)

	// Additional example: integer division
	fmt.Println("10/4 =", 10/4) // Output: 10/4 = 2 (integer division truncates the decimal part)

	// Additional example: using float for division
	fmt.Println("10.0/4.0 =", 10.0/4.0) // Output: 10.0/4.0 = 2.5

	// Booleans, with boolean operators as you'd expect.
	// Logical AND
	fmt.Println(true && false)

	// Logical OR
	fmt.Println(true || false)

	// Logical NOT
	fmt.Println(!true)

	// Additional boolean example: comparing integers
	a := 5
	b := 10
	fmt.Println("a > b =", a > b)
	fmt.Println("a < b =", a < b)
}
