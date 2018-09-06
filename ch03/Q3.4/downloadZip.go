package main

import (
	"archive/zip"
	"io"
	"net/http"
	"strings"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", "attachemt; filename=ascii_sample.zip")

	zipWriter := zip.NewWriter(w)
	defer zipWriter.Close()

	writer, err := zipWriter.Create("newfile.txt")
	if err != nil {
		panic(err)
	}
	io.Copy(writer, strings.NewReader("zip file"))
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
