# Range  
Range, in Go, is used with for loops to iterate over slices, maps, etc  
While I used `id`, `i`, `k`, and `v` as variable names for the for loop, it can be changed.   

### Iterate over slices and arrays  
Example slice: `ids := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}`  
To iterate over this slice, the following code could be used:  
```go
for i, id := range ids {
fmt.Printf("%d - Current Element/Item\n", id)
fmt.Printf("%d - Current Index\n", i)
}
```  
In the above example, `i` is the current index, and `id` is the current element. Thus, outputting `id` will return the current element/item in the slice, while `i` will return the current index  

### Iterate over maps  
Example map: `map[x:x@example.com y:y@example.com]`  
Example Code:  
```go
	for k, v := range emails {
		fmt.Printf("%s: %s\n", k, v) // Same as fmt.Println(k + ": " + v)
	}
```  
* Iterating over maps is similar to slices, except that instead of the index, there is a key  
* `k` is the key while `v` is the current value  