package main

import (
	"fmt"
)

// Define a struct named 'customer' to represent customer details.
type customer struct {
	name  string // Customer's name.
	phone string // Customer's phone number.
}

// Define a struct named 'order' which embeds the 'customer' struct and includes additional fields.
type order struct {
	id     string  // Order ID.
	amount float32 // Order amount.
	// createdAt time.Time // Time when the order was created.
	customer // Embedding the 'customer' struct.
}

func main() {
	fmt.Println("---Struct Embedding---")

	// Create a new 'order' instance, initializing the embedded 'customer' struct as well.
	newOrder := order{
		id:     "1",
		amount: 50,
		customer: customer{
			name:  "John",
			phone: "9654125895",
		},
		// createdAt: time.Now(), // Set the creation time of the order to the current time.
	}

	// Print the details of the new 'order'.
	fmt.Println("Initial Order:", newOrder)

	// Modify the embedded 'customer' field's 'name' in the 'order' instance.
	newOrder.customer.name = "Denny"

	// Print the updated details of the 'order'.
	fmt.Println("Updated Order:", newOrder)
}
