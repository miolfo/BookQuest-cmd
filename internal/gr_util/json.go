package gr_util

import (
	"encoding/json"
	"log"
	"os"
)

type BookSearchResult struct {
	Title     string
	Author    string
	FinnaIds  []string
	Available []bool
	Urls      []string
	Price     PriceWithCurrency
}

type PriceWithCurrency struct {
	Currency string
	Value    float32
}

type BookSearchResults struct {
	Results []BookSearchResult
}

func WriteResultsToPath(results BookSearchResults, path string) {
	res, marshalErr := json.MarshalIndent(results, "", "\t")
	if marshalErr != nil {
		log.Fatalf("Unable to marshall results: %s", marshalErr)
	}
	writeErr := os.WriteFile(path, res, 0644)
	if writeErr != nil {
		log.Fatalf("Unable to write results: %s", writeErr)
	}
}
