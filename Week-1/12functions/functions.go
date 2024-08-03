package main

import "fmt"

func main() {
	fmt.Println("---Functions---")

	res := sum(5, 2)
	fmt.Println("add:", res) // Output: add: 7

}

// Function that takes two ints and returns their sum
func sum(val1 int, val2 int) int {
	return val1 + val2
}
