// Package main defines the entry point for the executable program.
package main

// Import the "fmt" package for formatted I/O operations like printing to the console.
import "fmt"

// The main function is the entry point of the Go program.
func main() {
	// Println outputs the specified string to the console followed by a newline.
	fmt.Println("---Hello World---")
}

// To run the program, save the code in a file named hello-world.go and use the following command:
// $ go run hello-world.go
// This command compiles and runs the program, producing the output:
// hello world

// Sometimes we’ll want to build our programs into binaries. We can do this using the go build command:
// $ go build hello-world.go
// This generates an executable binary in the same directory:
// $ ls
// hello-world    hello-world.go

// We can then execute the built binary directly from the terminal:
// $ ./hello-world
// The output will be:
// hello world

// Now that we can run and build basic Go programs, let’s learn more about the language.
