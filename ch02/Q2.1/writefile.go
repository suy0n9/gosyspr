package main

import (
	"fmt"
	"os"
)

var filename = "example.txt"

func main() {
	WriteFile(filename)
}

// WriteFile wirte formatted string in file
func WriteFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	s := "Hello World"
	d := 5
	f := 123.456

	fmt.Fprintf(file, "%s %d %f", s, d, f)
}
