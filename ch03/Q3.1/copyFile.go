package main

import (
	"io"
	"os"
)

func main() {
	src := os.Args[1]
	dst := os.Args[2]
	file, err := os.Open(src)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	newFile, _ := os.Create(dst)
	if err != nil {
		panic(err)
	}
	defer newFile.Close()

	_, err = io.Copy(newFile, file)
	if err != nil {
		panic(err)
	}
}
