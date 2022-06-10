# Structs  
Structs are similar to classes in Python. They are templates to store sets of data.  

### Define a struct  
Before a struct can be created, a struct needs to be defined. While defining, the fields and the field type must be defined.  
Example:  
```go
type Person struct {
	firstname string
	lastname  string
    // Define multiple fields of the same type: firstname lastname string
	age int
}
```  

### Inializing a struct  
Once a struct is defined, the struct can be initialized. There are two primary ways to initialize a struct. Both accomplish the same.  
* `var exampleInitializedStruct = exampleUn Struct{value1: "x", value2: "y", value3: 3}`  - Long form  
* `exampleInitializedStruct := exampleUninitializedStruct{value1: "x", value2: "y", value3: 3}`  
Also, when initializing, the field names are optional. For example, `x := y{value1: "x"}` is the same as `x := y{"x"}`  

### Accessing and manipulating struct data  
* `fmt.Println(exampleInitializedStruct.field)` - Example of getting a field's value and printing it  
* `exampleInitializedStruct.field = 'new value` - Example of changing the value of a struct field  
#### Methods  
Methods are functions specific to a struct type  
##### Method types  
There are two method types, value recievers and pointer recievers  
* Value reciever - accesses struct data, but does not change data  
* Pointer reciever - Can access and change struct data  
##### Defining a method (value reciever)  
Defining a method is similar to defining a function. However, the struct type and an identifier must be defined  
For example, `func (p Person) greet() string {return "example return value"}`  
In the above example, `p` in the identifier, `Person` is the struct type, `greet()` is the method name, and `string` is the return type  
##### Accessing struct data from a method  
Using a method is like using a normal function. However, struct data can be accessed. To access struct data inside of a method, use the following format: `x = i.exampleField`. In this example, `x` is is set to the value of the field `exampleField`. `i` in this case, is the identifier set during method declaration.  
##### Running a method  
`exampleStruct.exampleMethod(parameter1, parameter2)` - Run a method that uses `exampleStruct`  
##### Defining a method (pointer reciever)  
Defining a method that's a pointer reciever is similar to a defining a value reciever. However, the struct type must be a pointer.  
For example, `func (p *Person) greet() string {return "example return value"}` (notice the asterisk before `Person`)  
##### Using a method (pointer reciever)  
Using a pointer reciever is similar to using a value reciever. However, it's possible to change values in the struct.  
Example of changing a value: `i.exampleField = x`. In this example, exampleField is set to x. Also,`i` is the identifier set during  method declaration. Running a pointer reciever is the same as running a value reciever (`exampleStruct.exampleMethod(parameter1, parameter2)`)  
