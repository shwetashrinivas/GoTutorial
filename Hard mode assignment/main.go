/*
Create a program that reads the contents of a text file then prints its contents to the terminal.package main

The file to open should be provided as an argument to the program when it is executed at the terminal.

To read in the arguments provided to a program , you can reference the variable 'os.Args', which is a slice of string.

To open a file, check out the documentation for the 'Open' function in the 'os' package

The file type implements both the Reader and Writer interface.

Run: go run main.go myfile.txt

ALSO: For displaying contents of main.go-
go build main.go
main.exe main.go

*/

package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	//fmt.Println(os.Args[1]) //ouput : myfile.txt
	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, f) //Reading from my file
}
