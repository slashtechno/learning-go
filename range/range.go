package main

import "fmt"

func main() {
	// Slices/Arrays
	ids := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Loop through IDs
	for i, id := range ids { // i is the current index, id is the current element/item, and ids is the slice
		// for _, id := range ids{fmt.Printf("%d - ID: %d\n", i, id)} - excludes the index since there is an underscore
		fmt.Printf("%d - ID: %d\n", i, id) // Print a list of the index and ID in the format of "<index> - ID <value>"
	}

	// Add ids together
	sum := 0
	for _, id := range ids { // For each value in ids, as id:
		sum += id // same as sum = sum + id
		fmt.Println("Sum: ", sum)
	}

	// Maps
	// Declare a map
	emails := make(map[string]string) // Full form var emails map[string]string

	// Assign values
	emails["x"] = "x@example.com"
	emails["y"] = "y@example.com"

	// Iterate over map
	for k, v := range emails {
		fmt.Printf("%s: %s\n", k, v) // Same as fmt.Println(k + ": " + v)
	}

	for k := range emails {
		fmt.Println("Name: " + k) // The keys in the emails map are single letters. Thus, those will be outputted
	}
}
