package main // Needs to be in every Go program

// Import packages
import (
	"fmt"     // Used to output text
	"strconv" // Used to convert integers to strings
)

// All variables declared at the package level need full declaration syntax

func main() {
	fmt.Println("Hello, world!") // Print Hello, world!
	var x float32 = 1            // Declare value x as a float32 (allows decimals) with the value of 1
	fmt.Println(x)               // print x
	y := 2                       // More consise way to declare a variable with a value and type
	fmt.Println(y)               // print y
	y = 3                        // Change the value of y
	fmt.Printf("%v, %T", x, x)   // Print a formatting string with the value of x (%v) and type (%T)
	fmt.Print("\n")              // print a newline (fmt.Print doesn't make a newline unlike Println)
	// Declare multiple variables at once
	var (
		language        string = "go"
		editor          string = "VS Code"
		secondaryEditor string = "Vim"
	)
	fmt.Printf("I am learning %s and I am using %s as my editor. For quick edits, I use %s\n", language, editor, secondaryEditor)
	x = float32(y)      // Convert the value of y to a float and assign it to x
	var z string        // Declare z as a string
	z = strconv.Itoa(y) // Assign y (an integer) as a string to z
	fmt.Println(z)      // Print z
	fmt.Println("/ \\\n |\n |\n |\n the basics of variables in Go")
}
