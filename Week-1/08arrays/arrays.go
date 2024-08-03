package main

import "fmt"

func main() {
	fmt.Println("---Arrays---")
	// Example 1: Creating and initializing an array
	var a [5]int
	fmt.Println("emp:", a)

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	fmt.Println("len:", len(a))

	// Declare and initialize an array in one line
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// Compiler counts the number of elements
	b = [...]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// Specifying indices
	b = [...]int{100, 3: 400, 500}
	fmt.Println("idx:", b)

	// Multi-dimensional arrays
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)

	// Initialize multi-dimensional arrays
	twoD = [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println("2d:", twoD)

	var namelist [3]string
	namelist[0] = "dixit"
	namelist[1] = "lukhi"
	fmt.Println("names:", namelist)

	namelist[2] = "shrut"
	fmt.Println("names:", namelist)
	fmt.Println("names length:", len(namelist))

	var alphabetList = [5]string{"a", "b", "c", "d"}
	fmt.Println("alphabets:", alphabetList)
	fmt.Println("alphabets length:", len(alphabetList))
}
