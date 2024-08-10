package main

import "fmt"

// MapKeys returns a slice of keys from the given map.
func MapKeys[K comparable, V any](m map[K]V) []K {
	// Create a slice to store the keys with a length equal to the number of keys in the map.
	r := make([]K, 0, len(m))
	// Iterate over the map to get each key.
	for k := range m {
		// Append each key to the slice.
		r = append(r, k)
	}
	return r // Return the slice of keys.
}

// Min is a generic function that returns the minimum of two values of any ordered type.
func Min[T int | float64 | string](a, b T) T {
	// Compare the two values and return the smaller one.
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println("---Generics---")

	// Demonstrate the Min function with different types.
	fmt.Println(Min(3, 5))           // Works with int
	fmt.Println(Min(3.2, 1.5))       // Works with float64
	fmt.Println(Min("apple", "zoo")) // Works with string

	// Create a map with integer keys and string values.
	var m = map[int]string{1: "2", 2: "4", 4: "8"}

	// Print the keys of the map by calling MapKeys.
	fmt.Println("Keys:", MapKeys(m))
}
