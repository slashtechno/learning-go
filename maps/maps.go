package main

import "fmt"

func main() {
	// Declare a map
	emails := make(map[string]string) // Full form var emails map[string]string

	// Assign values
	emails["x"] = "x@example.com"
	emails["y"] = "y@example.com"

	fmt.Print("Current email map: ")
	fmt.Println(emails)

	// Accessing values
	fmt.Println(emails["x"]) // Print email for user x

	// Delete from map
	delete(emails, "y") // Delete y from map

	// Initialize map with key value pairs
	phones := map[string]int{"x": 1005551000, "y": 2005552000}
	fmt.Print("Current phone map: ")
	fmt.Println(phones)
}
