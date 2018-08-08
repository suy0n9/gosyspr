package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
)

var records = [][]string{
	{"first_name", "last_name", "username"},
	{"Rob", "Pike", "rob"},
	{"Ken", "Thompson", "ken"},
}

func main() {
	WriteCsv(os.Stdout, records)
}

// WriteCsv write csv
func WriteCsv(output io.Writer, records [][]string) {
	w := csv.NewWriter(output)

	for _, record := range records {
		if err := w.Write(record); err != nil {
			log.Fatalln("error writing record to csv:", err)
		}
	}

	w.Flush()

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
}
