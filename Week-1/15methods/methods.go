package main

import "fmt"

// Define the User struct
type User struct {
	Name  string
	Score int
	Live  bool
}

// Method with a value receiver
func (u User) getStatus() {
	fmt.Println("is alive:", u.Live)
}

type rect struct {
	width, height int
}

// Method with a pointer receiver
func (r *rect) area() int {
	return r.width * r.height
}

// Method with a value receiver
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	fmt.Println("---Methods---")

	// Create an instance of User
	dixit := User{"ld", 17, true}
	fmt.Println(dixit)
	fmt.Printf("dixit : %+v\n", dixit)
	fmt.Printf("dixit : %v\n", dixit.Name)

	// Call the method on the User instance
	dixit.getStatus()

	// Create an instance of rect
	r := rect{width: 10, height: 5}

	// Call methods on the struct
	fmt.Println("area: ", r.area())  // Output: area: 50
	fmt.Println("perim:", r.perim()) // Output: perim: 30

	// Methods can also be called on a pointer to the struct
	rp := &r
	fmt.Println("area: ", rp.area()) // Output: area: 50
	fmt.Println("perim:", rp.perim())
}
