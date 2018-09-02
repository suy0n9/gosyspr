package main

import (
	"archive/zip"
	"io"
	"os"
	"strings"
)

func main() {
	// zipの内容を書き込むファイル
	file, err := os.Create("sample.zip")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// zipファイル
	zipWriter := zip.NewWriter(file)
	defer zipWriter.Close()

	// 書き込み
	writer, err := zipWriter.Create("newfile.txt")
	if err != nil {
		panic(err)
	}
	io.Copy(writer, strings.NewReader("zip file"))
}
