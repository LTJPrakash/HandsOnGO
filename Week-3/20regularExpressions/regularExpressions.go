package main

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Println("---Regular Expressions---")

	// Example string
	text := "Hello, my email is example@example.com and my phone number is (123) 456-7890."

	// Compile a regex pattern to match email addresses
	emailPattern := `[\w\.-]+@[\w\.-]+\.\w+`
	emailRegex := regexp.MustCompile(emailPattern)

	// Find all email addresses in the text
	emails := emailRegex.FindAllString(text, -1)
	fmt.Println("Emails found:", emails)

	// Compile a regex pattern to match phone numbers
	phonePattern := `\(\d{3}\) \d{3}-\d{4}`
	phoneRegex := regexp.MustCompile(phonePattern)

	// Find all phone numbers in the text
	phones := phoneRegex.FindAllString(text, -1)
	fmt.Println("Phone numbers found:", phones)

	// Replace email addresses with "[EMAIL]"
	replacedText := emailRegex.ReplaceAllString(text, "[EMAIL]")
	fmt.Println("Text with emails replaced:", replacedText)

	// Check if the text contains a phone number
	containsPhone := phoneRegex.MatchString(text)
	fmt.Println("Contains phone number:", containsPhone)
}
