// https://www.digitalocean.com/community/tutorials/how-to-use-json-in-go
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// Generating JSON from a map
	mapToConvert := map[string]interface{}{ // interface{} makes it possible to use any data type as a value
		"intValue":    1234,
		"boolValue":   true,
		"stringValue": "hello!",
		"objectValue": map[string]interface{}{
			"arrayValue": []int{1, 2, 3, 4},
		},
		"nullValue": nil, // nil is converted to a null value when converting to JSON
	}
	newJsonData, err := json.Marshal(mapToConvert) // Declare and set jsonData to the the map as JSON. If there is an error, err will be set to the error
	// json.Marshal outputs bytes. This can be converted to a string with string() or the %s verb in a formatting string
	if err != nil { // If there is an error, print the error
		fmt.Printf("Could not marshal data to JSON. Error: %v\n", err)
	}

	fmt.Printf("jsonData: %s\n", newJsonData) // Output the JSON data as a string

	// Converting JSON to a map
	jsonDataToLoad := `
		{
			"intValue":1234,
			"boolValue":true,
			"stringValue":"hello!",
			"dateValue":"2022-03-02T09:10:00Z",
			"objectValue":{
				"arrayValue":[1,2,3,4]
			},
			"nullStringValue":null,
			"nullIntValue":null
		}
	`
	var mapFromJSON map[string]any                             // Declare a map that can have string key type and any value type (any is the same as interface{})
	err = json.Unmarshal([]byte(jsonDataToLoad), &mapFromJSON) // Unmarshal the JSON data to a map. err is set to this as the map to unmarshal to is specified in the function as a pointer
	if err != nil {                                            // If there is an error, print the error
		fmt.Printf("Could not unmarshal data to a map. Error:\n %v\n", err)
	}
	fmt.Printf("Map:\n%v\nMap Type:\n%T\n", mapFromJSON, mapFromJSON)

	// Write the JSON data to a file
	file, err := os.Create("example.json") // Create a new file
	if err != nil {                        // If there is an error, print the error
		fmt.Printf("Could not create file. Error:\n %v\n", err)
	}
	defer file.Close()                    // Close file
	file.WriteString(string(newJsonData)) // Write the JSON data to a file

	// Load JSON data from the file
	var mapFromFile map[string]any
	fileContent, err := ioutil.ReadFile("example.json")
	if err != nil { // If there is an error, print the error}
		fmt.Printf("Could not read file. Error:\n %v\n", err)
	}
	err = json.Unmarshal([]byte(fileContent), &mapFromFile) // Unmarshal file contentz
	if err != nil {                                         // If there is an error, print the error
		fmt.Printf("Could not unmarshal data to a map. Error:\n %v\n", err)

	}
	fmt.Printf("mapFromFile: %v\n", mapFromFile)
}
