package main

import (
	"fmt"
	"strconv"
)

// Define person struct
type Person struct {
	firstname string
	lastname  string
	// firstname lastname string
	age int
}

// Greeting method (value reciever)
func (p Person) greet() string { // Declare a method named greet with p as an identifier and the name of the struct (Person) with a string as a return value
	return "Hello, my name is " + p.firstname + " " + p.lastname + ". I am " + strconv.Itoa(p.age) // Return "Hello, my name is <firstname> <lastname>. I am <age>"
}

// hasBirthday method (pointer reciever)
func (p *Person) hasBirthday() {
	p.age++ // Add 1 to age
}

// changesName method (pointer reciever)
func (p *Person) changeFirstName(newName string) {
	// Return if the new name is equal to the old name
	if p.firstname == newName {
		fmt.Println("New name is equal to old name")
		return
	} else {
		p.firstname = newName // Change the first name to the new name
	}
}

func main() {
	// Initialize person using struct
	person1 := Person{firstname: "John", lastname: "Doe", age: 30} // short form: person1 := Person{"John", "Doe", 30}
	fmt.Println(person1)                                           // Output: {John Doe 30}
	fmt.Println(person1.firstname)                                 // Get an individual value of a struct
	person1.age = 40                                               // Change value of individual struct value
	person1.hasBirthday()                                          // Increase the age by 1 using the hasBirthday method
	person1.changeFirstName("xyz")
	fmt.Println(person1.greet()) // Print the result of person1.greet (greet is a method)
}
