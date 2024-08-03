package main

import "fmt"

func main() {
	fmt.Println("---If Else---")
	// Basic if-else example
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	// Using logical operators in conditions
	if 8%2 == 0 || 7%2 == 0 {
		fmt.Println("either 8 or 7 are even")
	}

	// Variable declaration in if statement
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

	cnt := 17
	var res string
	if cnt < 17 {
		fmt.Println("ok")
		res = "less"
	} else if cnt == 17 {
		res = "equal"
	} else {
		fmt.Println("more")
		res = "more"
	}

	fmt.Println(res)

	if cnt = 25; cnt < 17 {
		fmt.Println("ok")
		res = "lesssec"
	} else if cnt >= 17 {
		res = "moresec"
	}

	fmt.Println(res)
}
