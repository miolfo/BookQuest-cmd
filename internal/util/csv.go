package util

import (
	"encoding/csv"
	"log"
	"os"
)

func ReadCsvFromPath(path string) [][]string {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal("Error opening file")
	}
	defer f.Close()
	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse csv file")
	}
	return records
}
