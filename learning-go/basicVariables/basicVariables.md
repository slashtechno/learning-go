# Basic Variables in Go  

## There are three ways to declare a variable in Go

### Declaring without a value   
`var x string`

### Declaring with a value and specific type    
`var x string = "Hello, world!"`  
* This is the only way to declare a variable outside of `func main(){}`  

### Declaring a variable with only a value (inferring the type)    
`x := "Hello, world!"`

## Changing the value of a variable  
`x = "Hello, Mars!"`  

## Printing text on a new line  
`fmt.Println("Hello, world!")`  
* fmt needs to be imported  