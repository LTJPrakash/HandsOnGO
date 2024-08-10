package main

import "fmt"

// Demonstrates handling of strings and runes in Go.
func main() {
	fmt.Println("---Strings and Runes---")

	// Define a string containing both ASCII and non-ASCII characters.
	s := "Hello, 世界"
	fmt.Println("Original string:", s) // Print the original string.

	// Print the byte value of the character at index 1.
	// Note: This will print the ASCII value for 'e', which is 101.
	fmt.Println("Byte value at index 1:", s[1])

	// Convert the string to a slice of runes to handle non-ASCII characters correctly.
	runes := []rune(s)

	// Iterate over the runes and print each one with its index.
	for i, r := range runes {
		// Print the index and the character represented by the rune.
		fmt.Printf("Rune %d: %c\n", i, r)
	}
}
