package main

import "fmt"

//  Declare a function that takes a string as a paremeter and returns a function
func exampleClosure(x string) func() {
	return func() { fmt.Println(x) } // This function, which is declared in another function, uses the other function's pareameter. Thus, its a variable
}

func exampleFunction(paremeter1 string, parameter2 string) {
	fmt.Printf("The value of paremeter1 is %v and the value of paremeter2 is %v\n", paremeter1, parameter2)
	return
}

func main() {
	exampleFunction("Hello, Earth!", "Hello, Mars!")
	exampleClosure("Hello, world!")()
}
