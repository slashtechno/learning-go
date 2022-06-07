# Maps   
Maps are used to store key-value pairs  

## Declaring a map  
When declaring a map, the first type is for specifying key type while the second is for specifying the value type  
* `var exampleMap map[string]string` - Delcare a nil map 
    * Needs to be initialized  
* `exampleMap := map[int]string{}` - Declare an initialized map  
* `var exampleMap = make(map[string]string)`  - Declare an initialized map  

## Assigning values
* `exampleMap["key"] = value` - Add a key-value pair, or set an existing key to a value  
* `exampleMap = map[string]string { "x": "a", "y": "b"}` - Inialize a map with data  
    * Will overwrite existing data  

## Accessing Map Data  
#### Retrieve value from a key  
`fmt.Println(exampleMap["key"])`  
#### Iterating over map data  
```go
for k, v := range exampleMap {
  fmt.Println("Key", k)
  fmt.Println("Value", v)
}
```  
#### Deleting a key-value pair  
`delete(exampleMap["key"])`  
