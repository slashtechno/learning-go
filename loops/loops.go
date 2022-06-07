package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Long method
	i := 1
	fmt.Println("The following is outputted from the long method to create a loop")
	for i <= 10 { // While i is less than or equal to 10, repeat (10 times)
		fmt.Println("This loop has repeated " + strconv.Itoa(i) + " times")
		i++ // i = i + 1
	}
	// Short method
	fmt.Println("The following is outputted from the short method to create a loop")
	for i := 1; i <= 10; i++ {
		fmt.Printf("This loop has repeated %d times\n", i)
	}

	//FizzBuzz (loop through 100, every multiple of 3 print out "Fizz" every multiple of 4, "Buzz", common multiple "FizzBuzz")
	fmt.Println("The following is the FizzBuzz challenge, a common interview question")
	for i := 1; i <= 100; i++ {
		if i%15 == 0 { // Check if i is divisible by 15
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 { // Check if i is divisible by 3
			fmt.Println("Fizz")
		} else if i%5 == 0 { // Check if i is divisible by 5
			fmt.Println("Buzz")
		} else {
			fmt.Println("Not divisible, current repetition is " + strconv.Itoa(i))
		}
	}
	// Loop through a slice
	fmt.Println("The following is the result of looping through a slice")
	exampleSlice := []string{"Item 1", "Item 2", "Item 3", "Item 4"}
	for i := 0; i < len(exampleSlice); i++ {
		fmt.Println(exampleSlice[i])
	}
}
