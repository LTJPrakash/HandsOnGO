package main

import "fmt"

func main() {
	fmt.Println("---For Loop---")

	// Basic for loop with a single condition
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// Classic initial/condition/after loop
	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	// Fixed version using a slice for range
	nums := []int{0, 1, 2}
	for i := range nums {
		fmt.Println("range", i)
	}

	// Infinite loop with break
	for {
		fmt.Println("loop")
		break
	}

	// Loop with continue to skip even numbers
	for n := 0; n < 6; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

	// Example with goto, continue, and break
	num := 1
	for num < 11 {
		if num == 3 {
			goto label
		}
		if num == 5 {
			num++
			continue
		}
		if num == 7 {
			break
		}
		fmt.Println(num)
		num += 1
	}
label:
	fmt.Println("from label")

	fmt.Println("End")
}
