package main

import (
	"fmt"
	"math"
)

// Declare a constant of type string
const s string = "constant"

func main() {
	fmt.Println("---Constant---")
	// Print the value of the string constant
	fmt.Println(s)

	// Declare a numeric constant
	const n = 500000000
	fmt.Println(n)

	// Perform arithmetic with constants
	// The expression d is a constant with arbitrary precision
	const d = 3e20 / n
	fmt.Println(d)

	// Numeric constants have no type until explicitly converted
	// Convert the constant d to int64
	fmt.Println(int64(d))

	// Use a numeric constant in a function that requires a specific type
	// math.Sin expects a float64, so n is implicitly converted to float64
	fmt.Println(math.Sin(float64(n)))
}
