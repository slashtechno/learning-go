package main

import "fmt"

func main() {
	a := 5
	b := &a // Point b to the memory address (location in memory) of a
	fmt.Println(a, b)

	// Use  * to read value from address
	fmt.Printf("b type: %T\n", b) // An asterisk before a type repersents a pointer
	fmt.Printf("Value of *b: %v\n", *b)

	// Change value with pointer
	*b = 10
	fmt.Printf("Because *b was changed to 10, so was a since b is a pointer of a.\na: %v\nb: %v\n*b: %v", a, b, *b)

}
