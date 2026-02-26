/*
A slice is a dynamically-sized, flexible view over an underlying array.
👉 In real-world Go programs, you use slices much more than arrays.

A slice contains three components:
• Pointer: The pointer is used to points to the first element of the array that is accessible through the slice.
Here, it is not necessary that the pointed element is the first element of the array.
• Length: The length is the total number of elements present in the array.
Capacity: The capacity represents the maximum size upto which it can expand.
*/
package main

import (
	"fmt"
)

func main() {
	//Declaring a Slice
	var s []int // Empty slice

	// Using literal
	s1 := []int{1, 2, 3}

	//Using make
	s2 := make([]int, 5)     // length = 5, capacity = 5
	s3 := make([]int, 5, 10) // length = 5, capacity = 10

	//Length and Capacity
	fmt.Println(len(s2)) // 5
	fmt.Println(cap(s2)) // 10

	//Appending to Slice
	s5 := []int{1, 2}
	s5 = append(s, 3)
	//Append multiple:

	s = append(s, 4, 5)
	//Append another slice:

	a := []int{6, 7}
	s = append(s, a...)

	//Slicing (Creating Sub-slices)
	arr := []int{1, 2, 3, 4, 5}

	sub := arr[1:4] // [2 3 4]

	/*Syntax:
	  slice[start:end]
	   Includes start
	   Excludes end
	*/

	//Slices Share Underlying Array
	a6 := []int{1, 2, 3}
	b := a6

	b[0] = 100

	fmt.Println(a) // [100 2 3]

	//Copying a Slice (Important), To avoid modifying original

	a7 := []int{1, 2, 3}
	b7 := make([]int, len(a7))
	copy(b7, a7)
	//Now b is independent.

	//Looping Over Slice
	for i, v := range s {
		fmt.Println(i, v)
	}

	fmt.Println(s, s1, s2, s3, s5, sub)
}
