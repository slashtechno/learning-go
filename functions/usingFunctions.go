package main

import "fmt"

func exampleFunction(paremeter1 string, paremeter2 string) {
	fmt.Printf("The value of paremeter1 is %v and the value of paremeter2 is %v\n", paremeter1, paremeter2)
	return
}

func main() {
	exampleFunction("Hello, Earth!", "Hello, Mars!")
}
