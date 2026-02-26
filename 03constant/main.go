/*
constant is a fixed value that cannot be changed after it is declared. Constants are declared using the const keyword.
*/
package main

import (
	"fmt"
)

func main() {
	//Basic Constant Declaration
	const Pi = 3.14
	const Name = "John"
	const IsActive = true
	//You can also declare multiple constants together:

	const (
		A = 10
		B = 20
		C = "Hello"
	)

	// Untyped Constants (default)
	const x = 10

	//Typed Constants
	const y int = 10 // type is fixed

	//Constant Expressions
	//Constants can be calculated at compile time
	const I = 10
	const J = 20
	const K = A + B

	//Special: iota
	//iota is a predeclared identifier used for generating successive constant values
	const (
		Monday = iota
		Tuesday
		Wednesday
	)

	/*
		Values:

			Monday = 0
			Tuesday = 1
			Wednesday = 2
	*/

	const (
		L = iota + 1
		O
		M
	)
	/*Results:

	  L = 1
	  O = 2
	  M = 3
	*/

	fmt.Println(Pi, Name, IsActive)
}
