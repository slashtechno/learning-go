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