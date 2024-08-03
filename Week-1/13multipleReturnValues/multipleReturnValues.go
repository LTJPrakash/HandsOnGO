package main

import "fmt"

// Function that returns two integers
func vals() (int, int) {
	return 3, 7
}

func proAdder(values ...int) (int, string) {
	total := 0
	for _, val := range values {
		total += val
	}
	return total, "From sec args"
}

func main() {
	fmt.Println("---Multiple Return Values---")
	// Multiple assignment to capture both return values
	a, b := vals()
	fmt.Println(a) // Output: 3
	fmt.Println(b) // Output: 7

	// If only one value is needed, use the blank identifier _
	_, c := vals()
	fmt.Println(c) // Output: 7

	moreSum, str := proAdder(5, 1, 2, 0, 012)
	fmt.Println(moreSum, str) // Output: 10 From sec args
}
