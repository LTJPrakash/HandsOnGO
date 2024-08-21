package main

import (
	"fmt"
	"sort"
)

// Person represents an individual with a name and age.
type Person struct {
	Name string
	Age  int
}

func main() {
	fmt.Println("---Sorting By Functions---")

	// Create a slice of Person with some sample data
	people := []Person{
		{"Alice", 30},
		{"Bob", 25},
		{"Charlie", 35},
		{"Alice", 25}, // Same name but different age
	}

	// Sort by Age (ascending)
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})

	fmt.Println("Sorted by age:")
	for _, person := range people {
		fmt.Printf("%s, %d years old\n", person.Name, person.Age)
	}

	// Sort by Name (alphabetical) and Age (ascending) for people with the same name
	sort.Slice(people, func(i, j int) bool {
		if people[i].Name == people[j].Name {
			return people[i].Age < people[j].Age // Sort by age if names are the same
		}
		return people[i].Name < people[j].Name // Otherwise, sort by name
	})

	fmt.Println("\nSorted by name and age:")
	for _, person := range people {
		fmt.Printf("%s, %d years old\n", person.Name, person.Age)
	}
}
