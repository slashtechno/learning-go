# Loops  
Loops in Go are similar to JavaScript.  

## Longer method to create a loop  
The longer method to create a loop involves declaring a variable (often i) and after each repetition, adding 1 to the variable. The loop will only contiue while the variable is less than or equal to the number specified.  
Example:  
```go
i := 1
for i <= 10 { // While i is less than or equal to 10, repeat (10 times)
	fmt.Println(i)
	i++ // i = i + 1
}
```  

## Short method  
The shorter method to create a loop involves declaring and adding to the variable in the for statement rather than in parts of the code.  
Example:  
```go
for i := 1; i <= 10; i++ {
	fmt.Println(i)
}
```  