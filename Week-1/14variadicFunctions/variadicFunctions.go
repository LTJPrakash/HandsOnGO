package main

import "fmt"

// Variadic function that takes an arbitrary number of int arguments
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {
	fmt.Println("---Variadic Functions---")
	// Calling the variadic function with different numbers of arguments
	sum(1, 2)
	sum(1, 2, 3)

	// Calling with a slice
	nums := []int{1, 2, 3, 4}
	sum(nums...)

	// fmt.Println can take any number of arguments
	fmt.Println("Hello", "world", 123, true)
}
