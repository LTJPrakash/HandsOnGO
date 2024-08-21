package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("---String Functions---")

	// Example string with extra spaces and punctuation
	str := "  Hello, Go World!  "

	// Length of the string
	fmt.Println("Length:", len(str))

	// Case Conversion
	fmt.Println("ToLower:", strings.ToLower(str)) // Convert to lowercase
	fmt.Println("ToUpper:", strings.ToUpper(str)) // Convert to uppercase
	fmt.Println("Title:", strings.Title(str))     // Convert to title case (each word starts with uppercase)

	// Trimming
	fmt.Println("TrimSpace:", strings.TrimSpace(str))        // Remove leading and trailing spaces
	fmt.Println("Trim:", strings.Trim(str, " H!"))           // Remove specific characters from both ends
	fmt.Println("TrimLeft:", strings.TrimLeft(str, " H"))    // Remove specific characters from the left
	fmt.Println("TrimRight:", strings.TrimRight(str, " d!")) // Remove specific characters from the right

	// Splitting and Joining
	words := strings.Split(strings.TrimSpace(str), " ") // Split string into words
	fmt.Println("Split:", words)                        // Print the slice of words
	fmt.Println("Join:", strings.Join(words, "-"))      // Join words with a hyphen

	// Searching
	fmt.Println("Contains 'Go':", strings.Contains(str, "Go"))               // Check if substring exists
	fmt.Println("HasPrefix '  Hello':", strings.HasPrefix(str, "  Hello"))   // Check prefix
	fmt.Println("HasSuffix 'World!  ':", strings.HasSuffix(str, "World!  ")) // Check suffix
	fmt.Println("Index of 'Go':", strings.Index(str, "Go"))                  // Find index of substring

	// Replacing
	fmt.Println("Replace 'World' with 'Golang':", strings.Replace(str, "World", "Golang", 1)) // Replace first occurrence

	// Repeating
	fmt.Println("Repeat 'Go ' 3 times:", strings.Repeat("Go ", 3)) // Repeat a substring
}
