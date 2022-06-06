# Arrays and Slices  

## Arrays  
* Arrays need to be a fixed length  
* The type of an array item must be declared  
* Because of these limitations, slices are often used instead of arrays  
* Arrays use 0 indexing  
Example of array usage:  
```go  
exampleArray[1] = "Item 2"
var exampleArray [2]string // Declare an array with two values which are both string
// Assign values to exampleArray (Go uses 0 indexing)
exampleArray[0] = "Item 1"
```  

## Slices  
* Slices are very similar to arrays  
* The length of a slice does not need to be specified  
