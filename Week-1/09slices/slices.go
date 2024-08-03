package main

import (
	"fmt"
	"slices"
	"sort"
)

func main() {
	fmt.Println("---Slices---")
	// Example 1: Creating and initializing slices
	var s []string
	fmt.Println("uninit:", s, s == nil, len(s) == 0)

	s = make([]string, 3)
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))

	// Append elements to the slice
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	// Copying slices
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	// Slicing operations
	l := s[2:5]
	fmt.Println("sl1:", l)
	l = s[:5]
	fmt.Println("sl2:", l)
	l = s[2:]
	fmt.Println("sl3:", l)

	// Declare and initialize a slice in one line
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	// Using slices package for utility functions
	t2 := []string{"g", "h", "i"}
	if slices.Equal(t, t2) {
		fmt.Println("t == t2")
	}

	// Multi-dimensional slices
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)

	var nameList = []string{"dixit", "shrut"}
	fmt.Printf("Type : %T\n", nameList)
	fmt.Println("List : ", nameList)

	nameList = append(nameList, "M. Tech")
	fmt.Println("List : ", nameList)

	// Remove the second element by slicing
	nameList = append(nameList[0:1], nameList[2:]...)
	fmt.Println("List : ", nameList)

	// Working with integer slices
	scores := make([]int, 5)
	scores[0] = 1
	scores[1] = 14
	scores[2] = 4540
	scores[3] = 031
	scores[4] = 48450

	scores = append(scores, 4545, 78878)
	fmt.Println("scores : ", scores)

	sort.Ints(scores)
	fmt.Println(scores)
	fmt.Println(sort.IntsAreSorted(scores))

	// Remove an element by index
	var ind int = 2
	scores = append(scores[:ind], scores[ind+1:]...)
	fmt.Println("remove : ", scores)
}
