package main

import (
	"fmt"
	"math"
	"packagesInGo/strutil"
)

func main() {
	// Examples of using built-in packages
	fmt.Println("fmt can be used to output text.\nThis text is outputted using fmt.Println") // Example of using fmt to output text
	fmt.Println("2.7 rounded down is", math.Floor(2.7))                                      // Use math.Floor to round down 2.7
	fmt.Println("2.7 rounded up is", math.Ceil(2.7))                                         // Use math.Ceil to round up 2.7
	fmt.Println("The square root of 16 is", math.Sqrt(16))
	fmt.Println("\"Hello, world!\" reversed is", strutil.Reverse("Hello, world!")) // Use math.Sqrt to find the square root of 16

}
