package main

import (
	"bufio"
	"fmt"
	"strings"
)

var source = `1行目
2行目
3行目 `

func main() {
	scanner := bufio.NewScanner(strings.NewReader(source))
	// 単語区切り
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		fmt.Printf("%#v\n", scanner.Text())
	}
}
