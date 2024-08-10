package main

import "fmt"

func main() {
	fmt.Println("---Structs---")

	// Create a new 'bill' using the 'newBill' function.
	myBill := newBill("DA's bill")

	// Print the 'bill' instance. It will show the name, items, and tip.
	fmt.Println(myBill)
}
