package main

import (
	"fmt"
	"math"
)

// Define the geometry interface with two methods: area and perim
type geometry interface {
	area() float64
	perim() float64
}

// Define the rect type with width and height
type rect struct {
	width, height float64
}

// Implement the area method for rect
func (r rect) area() float64 {
	return r.width * r.height
}

// Implement the perim method for rect
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

// Define the circle type with radius
type circle struct {
	radius float64
}

// Implement the area method for circle
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// Implement the perim method for circle
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// measure function takes a geometry interface as a parameter
// and prints the area and perimeter of the shape
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	fmt.Println("---Interfaces---")

	// Create an instance of rect
	r := rect{width: 3, height: 4}

	// Create an instance of circle
	c := circle{radius: 5}

	// Call measure with both rect and circle instances
	measure(r)
	measure(c)
}
