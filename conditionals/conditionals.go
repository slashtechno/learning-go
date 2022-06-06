package main

import "fmt"

func main() {
	x := 5
	y := 10
	// Print which variable is less than the other
	if x < y {
		fmt.Printf("%d is less than %d\n", x, y) // %d is similar to %v, but it's used for base10 integers
	} else if x == y {
		fmt.Printf("%d is equal to %d\n", x, y)
	} else {
		fmt.Printf("%d is less than %d\n", y, x)
	}
}
