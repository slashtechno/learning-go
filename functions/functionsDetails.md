# Functions  
* Functions in Go are very similar to functions in other languages  
* A function must be declared outside of `main()`

### Declaring a function  
A function can be declared with the `func` keyword. In addition, parameters can be declared by adding parameters inside the parentheses following the function name. However, the parameter type must also be declared. Similarly, to other languages, these paremeters act as variables inside of the function. The return type must also be declared if data is being returned. It's important to note that functions must return, but data does not need to be returned. If this is the case, a return type is not required.  
For example:  
```go  
func functionName(parameter1 string, parameter2 string) string {
    fmt.Printf("The value of parameter1 is %v and the value of parameter2 is %v\n", parameter1, parameter2)
    return "The value of parameter1 is " + parameter1 + "and the value of parameter2 is " + parameter2
}
```  

### Using a function  
A function can be called with `functionName()`. Inside the parentheses, parameters can be added. For example, `exampleFunction(parameter1, parameter2)`  

### Closures  
Closures allow a function that's inside another function to accept the other function's parameters.  

For example, in the following code, the function `exampleClosure` accepts a parameter which is passed on to a function inside of it which prints out the parameter. The function using a parameter from the function it was declared in is a closure. Also, because the function `exampleClosure` returns a function, a second pair of parentheses needs to be added.  

```go
func exampleClosure(x string) func() {
	return func() { fmt.Println(x) } // This function, which is declared in another function, uses the other function's pareameter. Thus, its a variable
}
func main() {
	exampleFunction("Hello, world!")()
}
```  