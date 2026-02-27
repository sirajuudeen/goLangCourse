/*
map is a built-in, unordered collection of unique key-value pairs, implemented as a hash table.
They are used for fast data lookups, retrieval, updates, and deletions based on keys
*/
package main

import (
	"fmt"
)

func main() {
	//Declaration and Initialization:
	//Using make(): The idiomatic way to create an empty map, ready for use.

	//ages := make(map[string]int)

	//Using a Map Literal:
	ages := map[string]int{
		"Alice": 30,
		"Bob":   25,
	}

	//Adding or Updating Elements: Use the assignment syntax with a key in square brackets.
	ages["Alice"] = 29 // Add a new key-value pair
	ages["Alice"] = 31 // Update the value for an existing key

	//Retrieving Values: Access a value by specifying its key
	age := ages["Alice"]

	//Checking for Key Existence: Use the "comma ok" idiom to check if a key is present and
	// distinguish between a missing key and a zero value
	val, ok := ages["Bob"]
	if ok {
		// Key exists, val is the value
	} else {
		// Key does not exist
	}
	//Deleting Elements: Use the built-in delete function
	delete(ages, "Bob")

	//Iterating over a Map: Use a for...range loop to iterate over key-value pairs.
	for key, value := range ages {
		fmt.Println("Key:", key, "Value:", value)
	}

	fmt.Println(ages, age, val)
}
