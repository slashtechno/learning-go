package main // Needs to be in every Go program

import "fmt" // fmt allows for text to be outputed

// Import packages

func main() {
	fmt.Println("Hello, world!") // Print Hello, world!
	var x int = 1                // Declare value x as an integer with the value of 1
	fmt.Println(x)               // print x
	y := 2                       // More consise way to declare a variable with a value and type
	fmt.Println(y)               // print y
	y = 3                        // Change the value of y
	fmt.Printf("%v, %T", x, x)   // Print a formatting string with the value of x (%v) and type (%T)

}
