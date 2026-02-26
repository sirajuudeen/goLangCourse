/*
An array in Go is a fixed-size sequence of elements of the same type.
*/

package main

import (
	"fmt"
)

func main() {
	//Declaration
	var arr [5]int

	/*
		[5] → length is part of the type
		int → element type
		Default values are zero values (0 for int)
	*/

	//Initialization
	arr1 := [5]int{1, 2, 3, 4, 5}

	//	Let Go count the size
	arr2 := [...]int{1, 2, 3} //Go automatically sets the length to 3.

	//Accessing Elements
	arr3 := [3]string{"Go", "Java", "Python"}

	//Updating Elements
	arr4 := [3]int{1, 2, 3}
	arr[1] = 10

	//Looping Over Array
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
	//Using range
	for index, value := range arr {
		fmt.Println(index, value)
	}
	//Ignore index if not needed:

	for _, value := range arr {
		fmt.Println(value)
	}

	fmt.Println(arr)
	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println(arr4)
}
