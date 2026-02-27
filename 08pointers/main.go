/*
In Go, a pointer is a variable that stores the memory address of another variable.
Go supports pointers, but does not allow pointer arithmetic (unlike C/C++). This makes Go safer and simpler.
*/
package main

import (
	"fmt"
)

// Avoid copying large structs
type User struct {
	Name string
	Age  int
}

func main() {
	x := 10
	p := &x // p stores the address of x

	fmt.Println(x)  // 10
	fmt.Println(p)  // memory address (example: 0xc000018030)
	fmt.Println(*p) // 10 (value at that address)

	change(x)
	change1(&x)

	/*
		Explanation:
			&x → gives address of x
			p → pointer to x
			*p → dereference (get value stored at address)
	*/

	//1. Creating a Pointer to an Existing Struct
	p1 := User{Name: "Alice", Age: 30}
	p2 := &p1 // p2 is a pointer to p1

	fmt.Println(p1.Name) // Output: Alice
	fmt.Println(p2.Name) // Output: Alice (automatic dereferencing)

	p2.Age = 31         // Modifies the original p1 struct
	fmt.Println(p1.Age) // Output: 31

	//Creating a Struct Pointer Directly
	// Using the & operator during initialization
	//p2 := &User{Name: "Bob", Age: 42}

	// Using the new() function (fields will have their zero values)
	p3 := new(User)
	p3.Name = "Charlie"
	p3.Age = 25

	//Using Pointers in Functions/Methods
	//Passing a pointer to a function allows that function to modify the original struct,
	// which is not possible if you pass by value (which creates a copy
	person := User{Name: "John", Age: 30}
	modifyPerson(&person)    // Pass the address of the struct
	fmt.Println(person.Name) // Output: John Doe

}

// Modify original value inside a function
// Without pointer (value copy):
func change(num int) {
	num = 50
	fmt.Println("without pointer:-", num)
}

// With pointer:
func change1(num *int) {
	*num = 50
	fmt.Println("with pointer:-", num)
}

func modifyPerson(u *User) {
	u.Name = "John Doe" // Modifies the original Person struct
}
