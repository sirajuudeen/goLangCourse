/*
In Go, defer, panic, and recover are mechanisms used to handle exceptional, unrecoverable errors,
distinct from the typical explicit error handling via return values.

defer: schedules a function to be executed just before the surrounding function returns,
whether the return is normal or due to a panic.

panic: immediately stops the normal execution of the current function and begins unwinding the stack,
running any deferred functions along the way.

recover: is a built-in function used inside a deferred function to stop the panicking process, capture the
panic value (the argument passed to panic), and resume normal program execution
*/

package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("Start of main")
	// The deferred function is set up first, before any potential panic
	defer func() {
		if r := recover(); r != nil {
			// recover() works only inside a deferred function
			log.Println("Recovered from panic:", r)
		}
	}()

	riskyFunction(0)           // This will cause a panic
	fmt.Println("End of main") // This line won't be reached if a panic is unrecovered
}

func riskyFunction(b int) {
	// Defer statement inside this function runs during stack unwinding
	defer fmt.Println("Exiting riskyFunction")
	if b == 0 {
		panic("division by zero") // Program flow stops here
	}
	fmt.Println("Result:", 10/b)
}

/*
Output:
Start of main
Exiting riskyFunction
2026/02/27 07:21:10 Recovered from panic: division by zero
*/
