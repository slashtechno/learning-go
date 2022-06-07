# Arrays and Slices  
* `len()` can be used to check the length of a slice or an array  
* To access a range, use `example[1:3]` (this would return items 1 and 2)  

## Arrays  
* Arrays need to be a fixed length  
* The type of an array item must be declared  
* Because of these limitations, slices are often used instead of arrays  
* Arrays use 0 indexing  
* The append function can be used to append to an array.
    * For example, `example = append("an appended item", "another appended item")`
Example of array usage:  
```go  
exampleArray[1] = "Item 2"
var exampleArray [2]string // Declare an array of strings with two values
// exampleArray := [2string]{"Item 1", "Item 2" // More consise way to declare an array with values  }

// Assign values to exampleArray (Go uses 0 indexing)
exampleArray[0] = "Item 1"
fmt.Println(exampleArray[1]) // Print the second item of exampleArray
```  

## Slices  
* Slices are very similar to arrays  
* The length of a slice does not need to be specified  
* To keep the order of items in an array after removing an item, similar code to the following can be used (following code is from Linuxhint)  
    ```go
    slice := []string{"a", "b", "c", "d", "e", "f"}
    index := 1
    copy(slice[index:], slice[index+1:]) // shift valuesafter the indexwith a factor of 1
    slice[len(slice)-1] = "" // remove element
    slice = slice[:len(slice)-1]  // truncate slice
    ```  
Example of slice usage:  
```go
	// Slices
	exampleSlice := []string{"Item 1", "Item 2", "Item 3", "Item 4", "Item 5"} // Declare a slice of strings with five values. Unlike arrays, the number of items can change
	fmt.Println(exampleSlice[3])       
```  