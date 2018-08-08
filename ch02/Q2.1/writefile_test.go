package main

import (
	"io/ioutil"
	"testing"
)

func TestWriteFile(t *testing.T) {
	WriteFile("test.txt")
	actual, err := ioutil.ReadFile("test.txt")
	if err != nil {
		t.Fatal(err)
	}

	expected := "Hello World 5 123.456000"
	if string(actual) != expected {
		t.Errorf("want %q\ngot %q", expected, actual)
	}
}
