package main

import (
	"bufio"
	"fmt"
	"os"
)

// Global variable declaration
var users int = 450

// Global constant declaration
const LoginPass string = "abc123"

func main() {
	fmt.Println("---Variables---")

	// Variable declarations and initializations
	// Explicit variable declaration with initialization
	var a = "initial"
	fmt.Println(a) // Output: initial

	// Multiple variable declaration with initialization
	var b, c int = 1, 2
	fmt.Println(b, c) // Output: 1 2

	// Type inference for variable initialization
	var d = true
	fmt.Println(d)

	// Zero value for uninitialized variables
	var e int
	fmt.Println(e)

	// Short variable declaration and initialization
	f := "apple"
	fmt.Println(f)

	// Local variable declarations with type and initialization
	var username string = "Dixit"
	fmt.Println(username)
	fmt.Printf("Type : %T\n", username)

	var status bool = true
	fmt.Println(status)
	fmt.Printf("Type : %T\n", status)

	var val int = 17
	fmt.Println(val)
	fmt.Printf("Type : %T\n", val)

	var floatVal float32 = 17.932819
	fmt.Println(floatVal)
	fmt.Printf("Type : %T\n", floatVal)

	// Default zero value
	var defaultVal int
	fmt.Println(defaultVal)
	fmt.Printf("Type : %T\n", defaultVal)

	// Short variable declaration for local scope
	totalUsers := 5000
	fmt.Println(totalUsers)

	// Accessing global variable
	fmt.Println(users)

	// Accessing global constant
	fmt.Println(LoginPass)

	// User input handling
	welcomMsg := "Welcome to First Lab"
	fmt.Println(welcomMsg)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Give feedback point: ")

	// Reading input from user
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for ", input)
	fmt.Printf("Type of feedback: %T\n", input)
}
