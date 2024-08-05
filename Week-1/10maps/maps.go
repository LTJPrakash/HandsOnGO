package main

import (
	"fmt"
	"maps" // Import maps package for utility functions
)

func main() {
	fmt.Println("---Maps---")
	// Example 1: Creating and initializing a map
	m := make(map[string]int)
	m["k1"] = 7
	m["k2"] = 13
	fmt.Println("map:", m)

	// Getting a value
	v1 := m["k1"]
	fmt.Println("v1:", v1)

	// Key doesn't exist
	v3 := m["k3"]
	fmt.Println("v3:", v3)

	// Getting value and presence check
	v4, prs := m["k3"]
	fmt.Println("v4:", v4, "present:", prs)

	// Length of the map
	fmt.Println("len:", len(m))

	// Deleting an entry
	delete(m, "k2")
	fmt.Println("map:", m)

	// Declare and initialize a map
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	// Using maps package for utility functions
	n2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	}

	// Additional example with iteration
	langs := make(map[string]string)
	langs["JS"] = "JavaScript"
	langs["C"] = "C"
	langs["Cpp"] = "C++"
	langs["PY"] = "Python"

	fmt.Println("Langs : ", langs)
	fmt.Println("Langs : ", langs["PY"])

	delete(langs, "Cpp")
	fmt.Println("Langs : ", langs)

	for _, value := range langs {
		fmt.Println("For =", value)
	}
}
