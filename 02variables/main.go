/*
	Variable is used to store data that can be changed during program execution
*/

package main

import (
	"fmt"
)

// Global Variables , Declared outside functions
var globalVar = "I am global"

func main() {

	//Declaring Variables with types
	var age int = 25
	var name string = "John"

	//Using var (type inferred)
	var age1 = 25      // Go infers int
	var name1 = "John" // Go infers string

	//Short Declaration (Most Common) Only works inside functions
	age2 := 25
	name2 := "John"

	//Declaring Multiple Variables

	var a, b, c int = 1, 2, 3
	//or
	a1, b1, c1 := 1, 2, 3

	//Zero Values (Default Values)
	//If you declare without initializing, Go assigns a default value:

	var count int     // 0
	var price float64 // 0
	var isActive bool // false
	var text string   // ""

	fmt.Println(globalVar)
	fmt.Println(age, name)
	fmt.Println(age1, name1)
	fmt.Println(age2, name2)
	fmt.Println(a, b, c)
	fmt.Println(a1, b1, c1)
	fmt.Println(count, price, isActive, text)

}
