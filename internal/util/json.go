package util

import (
	"encoding/json"
	"log"
	"os"
)

// TODO: Maybe append goodreads id (if there is such) and finna id for easier matching in results?
type BookSearchResult struct {
	Title    string
	Author   string
	FinnaIds []string
	//TODO: Change to string after refactoring util to return actual status instead of a boolean
	Statuses []bool
	Urls     []string
}

type BookSearchResults struct {
	Results []BookSearchResult
}

func WriteResultsToPath(results BookSearchResults, path string) {
	res, marshalErr := json.MarshalIndent(results, "", "\t")
	if marshalErr != nil {
		log.Fatal("Unable to marshall results: %s", marshalErr)
	}
	writeErr := os.WriteFile(path, res, 0644)
	if writeErr != nil {
		log.Fatal("Unable to write results: %s", writeErr)
	}
}
