package main

import "fmt"

func main() {
	fmt.Println("---Range---")
	// Example 1: Using range with slices
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	// Range with index and value
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	// Example 2: Using range with maps
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// Range over keys only
	for k := range kvs {
		fmt.Println("key:", k)
	}

	// Example 3: Using range with strings
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
