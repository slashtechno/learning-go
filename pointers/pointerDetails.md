# Pointers  
* Pointers are the memory address of a variable  
* Printing a pointer will return the memory address  
* Printing a pointer with an asterisk before it will return the value of the original variable   
* Changing the pointer with an asterisk before it will change the value of the original variable  

#### Example code  
Declaring a pointer:  
```go
x := 5
y := &a
```  
Getting the original value for a variable from a pointer:  
```go
fmt.Println(*y) // Assumes y was declared as a pointer
```  
Changing the value of the original variable using a pointer:  
```go
*y = 5 // Assumes was a declared as a pointer
```  