package main

import (
	"bytes"
	"testing"
)

func TestWriteCsv(t *testing.T) {
	var tests = []struct {
		output   *bytes.Buffer
		input    [][]string
		expected string
	}{
		{&bytes.Buffer{},
			[][]string{{"col1", "col2"}, {"Robert", "Griesemer"}},
			"col1,col2\nRobert,Griesemer\n"},
	}

	for _, test := range tests {
		WriteCsv(test.output, test.input)
		actual := test.output.String()
		if string(actual) != test.expected {
			t.Errorf("want %q\ngot %q", test.expected, actual)
		}
	}
}
