package main

import "fmt"

func main() {
	// Arrays
	var exampleArray [2]string // Declare an array of strings with two values
	// exampleArray := [2string]{"Item 1", "Item 2" // More consise way to declare an array with values  }
	// Assign values to exampleArray (Go uses 0 indexing)
	exampleArray[0] = "Item 1"
	exampleArray[1] = "Item 2"
	fmt.Println(exampleArray[1]) // Print the second item of exampleArray

	// Slices
	exampleSlice := []string{"Item 1", "Item 2", "Item 3", "Item 4", "Item 5"} // Declare a slice of strings with five values. Unlike arrays, the number of items can change
	fmt.Println(exampleSlice[3])                                               // Print the fourth item of exampleSlice

}
