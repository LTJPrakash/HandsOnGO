package main

// Define a struct named 'bill' to represent a bill with items and a tip.
type bill struct {
	name  string            // Name of the bill or customer.
	items map[string]string // A map to hold item names and their values.
	tip   float64           // Tip amount.
}

// Function to create and initialize a new 'bill' instance.
func newBill(name string) bill {
	// Create a new 'bill' with the given name, some initial items, and a tip of 0.
	b := bill{
		name:  name,
		items: map[string]string{"k1": "val1", "k2": "val2"},
		tip:   0,
	}

	return b // Return the newly created 'bill' instance.
}
