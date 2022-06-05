package main

// Importing a single package
import (
	"fmt"
	"math"
)

// Importing multiple packages

func main() {
	// Examples of using built-in packages  
	fmt.Println("fmt can be used to output text.\nThis text is outputted using fmt.Println") // Example of using fmt to output text
	fmt.Println(math.Floor(2.7))                                                             // Use math.Floor to round down 2.7
	fmt.Println(math.Ceil(2.7))                                                              // Use math.Ceil to round up 2.7
	fmt.Println(math.Sqrt(16))                                                               // Use math.Sqrt to find the square root of 16
	
}
