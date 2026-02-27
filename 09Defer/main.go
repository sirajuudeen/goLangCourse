/*
defer schedules a function to be executed just before the surrounding function returns,
whether the return is normal or due to a panic.
*/
package main

import (
	"fmt"
	"os"
)

func main() {
	f := createfile("test.txt")
	defer closefile(f)
	WriteFile(f)

}

func createfile(p string) *os.File {
	fmt.Println("creating file")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f

}
func WriteFile(f *os.File) {
	fmt.Println("writing file")
	//fmt.Println(f, "data")

}
func closefile(f *os.File) {
	fmt.Println("closing file")
	err := f.Close()
	if err != nil {
		//fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}
}
