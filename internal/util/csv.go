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

func WriteRecordsToPath(bookPairs [][]string, path string) {
	csvFile, createErr := os.Create(path)
	if createErr != nil {
		log.Fatalf("Unable to create file %s", path)
	}
	csvWriter := csv.NewWriter(csvFile)
	writeErr := csvWriter.WriteAll(bookPairs)
	if writeErr != nil {
		log.Fatalf("Unable to write record: %s", writeErr)
	}
	defer csvWriter.Flush()
}
