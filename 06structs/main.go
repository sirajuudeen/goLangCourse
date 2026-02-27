/*
User-defined composite data type that groups together related fields (variables) under a single name.
Fields within a struct can be of different data types, allowing for the modeling of real-world entities
like a person with a name (string) and age (int).
*/
package main

import (
	"fmt"
)

// Define a struct named 'Person'
type Person struct {
	FirstName string
	LastName  string
	Age       int
}

// Define a struct named 'Student'
type Student struct {
	name     string
	rollNo   int
	subjects []string
}

// Define a struct named 'User'
type User struct {
	Name  string `json:"name"` // Tag used by the encoding/json package
	Email string `json:"email"`
}

func main() {
	// Initialize a struct using a composite literal (recommended syntax)
	p1 := Person{
		FirstName: "John",
		LastName:  "Doe",
		Age:       30,
	}
	fmt.Println("Person 1:", p1.FirstName, p1.LastName, p1.Age)

	student1 := Student{
		name:   "Siraj",
		rollNo: 5,
		subjects: []string{
			"Maths", "Science", "IT",
		},
	}
	student1.rollNo = 6
	fmt.Println("New Age:", student1)

	// Access and modify fields using the dot operator
	p1.Age = 31
	fmt.Println("New Age:", p1.Age)

	// Anonymous struct (used for one-off data groupings)
	dog := struct {
		name   string
		isGood bool
	}{"Rex", true}

	fmt.Println("Dog:", dog.name, dog.isGood)

	/*
		Struct Embedding (Composition) in Go

		In Go, struct embedding is a way to achieve composition (reusing code) instead of traditional inheritance.
		Go does not support inheritance, but it supports composition using embedded structs.
	*/

	type Person struct {
		Name string
		Age  int
	}

	type Employee struct {
		Person // Embedded struct
		ID     int
		Salary int
	}

	e := Employee{
		Person: Person{
			Name: "Siraj",
			Age:  20,
		},
		ID:     101,
		Salary: 50000,
	}

	fmt.Println(e.Name) // Promoted field
	fmt.Println(e.Age)
	fmt.Println(e.ID)
	fmt.Println(e.Salary)
}
