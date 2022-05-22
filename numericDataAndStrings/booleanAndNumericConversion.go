package main

import "fmt"

func main() {
	var x bool = true                                                                                // Declare a boolean variable as true
	x = false                                                                                        // Change x from true to false
	var y int = 0                                                                                    // Declare y as an integer with value set to 0
	x = y == y                                                                                       // If y is equal to y, set x to true
	fmt.Printf("x (%T) is %v because y (%T) is equal to %v and y is equal to itself \n", x, x, y, y) // Print details about x's value
	var z int32 = 20                                                                                 // Declare z as an int32 variable with value set to 20
	fmt.Println(y + int(z))                                                                          // Print the result of y + z. Variable z is an int32 while y is not. Thus, z needs to be converted to a default interger
}
