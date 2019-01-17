package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

var out io.Writer = os.Stdout // modified during testing

func main() {
	for index, arg := range os.Args {
		s := "[" + strconv.Itoa(index) + "] " + arg // Convert integert to string
		fmt.Fprintln(out, s)
	}
}
