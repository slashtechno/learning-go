# Conditionals  
* Conditionals in Go are similar to conditionals in other languages, especially C based languages  
* Parentheses are are rarely used in if statements, but an error will not be thrown if they are  
* To check if two values are equal, `==` should be used, not `=`  
* The or operator is `||`  
* The and operator is `&&`  

### Simple if statements  
```go
// If x(5) is less than y(10), print "5 is less than 10"
if x < y {
	fmt.Printf("%d is less than %d\n", x, y)
}
```  
### Else statements  
An else statement can be used after an if statement to run code if the if statement is not true.  
Example of an else statement:  
```go
// Print which variable is less than the other
if x < y {
	fmt.Printf("%d is less than %d\n", x, y)
} else { 
    fmt.Printf("%d is less than %d\n", y, x)
}
```  
### Else if statements  
Else if statements can be used to check if a different condition is true after an if statement.  
For example:  
```go
// Print which value is less than the other value. Also check if the values are equal, and output if they are.  
if x < y {
	fmt.Printf("%d is less than %d\n", x, y)
} else if x == y{
    fmt.Printf("%d is equal to %d\n", x, y)
}else { 
    fmt.Printf("%d is less than %d\n", y, x)
}
```  
### Switches  
Switches can be used to easily execute code depending on the state of a variable.  
The following is an example of a switch:  
```go
switch color {
    case "blue": 
        fmt.Println("Color is blue")
    case "green": 
        fmt.Println("Color is green")
    case "red": 
        fmt.Println("Color is red")
    default: 
    fmt.Println("Color is some other value")
}
```  
