package main

import (
	"fmt"
	"sort"
)

// Person represents a person with a name and age.
type Person struct {
	Name string
	Age  int
}

// ByAge is a type that implements sort.Interface for sorting a slice of Person by Age.
type ByAge []Person

// Len returns the number of elements in the ByAge slice.
func (a ByAge) Len() int { return len(a) }

// Less reports whether the element with index i should sort before the element with index j.
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

// Swap exchanges the elements with indexes i and j.
func (a ByAge) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

// main function demonstrates sorting a slice of Person structs by Age.
func main() {
	fmt.Println("---Sorting---")

	// Create a slice of Person
	people := []Person{
		{"Alice", 30},
		{"Bob", 25},
		{"Charlie", 35},
	}

	// Sort people by age using ByAge
	sort.Sort(ByAge(people))

	// Print the sorted slice
	fmt.Println("Sorted by age:")
	for _, person := range people {
		fmt.Printf("%s, %d years old\n", person.Name, person.Age)
	}
}
