package finnagr

import (
	"fmt"
	"github.com/miolfo/goodreads-finna/internal/util"
	"github.com/miolfo/goodreads-finna/pkg/finna"
	"github.com/miolfo/goodreads-finna/pkg/goodreads"
	"log"
)

func Finnagr(path string) {
	records := util.ReadCsvFromPath(path)
	books := goodreads.ParseBooks(records)
	var searchParams []finna.SearchParameters
	for _, book := range books {
		searchParams = append(searchParams, finna.SearchParameters{
			Title:    book.Title,
			Building: "0/Helmet/",
		})
	}

	searchedBook := searchParams[0]
	book, err := finna.FindBookByTitle(searchedBook)
	if err != nil {
		log.Print("No book found for title " + searchedBook.Title)
	}
	fmt.Println(book)
}
