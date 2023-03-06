package finnagr

import (
	"fmt"
	"github.com/miolfo/goodreads-finna/internal/util"
	"github.com/miolfo/goodreads-finna/pkg/finna"
	"github.com/miolfo/goodreads-finna/pkg/goodreads"
	"log"
	"time"
)

type BookPair struct {
	finnaBook finna.Book
	grBook    goodreads.Book
}

func Finnagr(path string) {
	records := util.ReadCsvFromPath(path)
	books := goodreads.ParseBooks(records)
	booksToRead := goodreads.FilterByShelf(books, "to-read")
	var searchParams []finna.SearchParameters
	for _, book := range booksToRead {
		searchParams = append(searchParams, finna.SearchParameters{
			Title:    book.Title,
			Building: "0/Helmet/",
		})
	}

	//bookPairs := []BookPair{}
	for _, searchParam := range searchParams {
		log.Printf("Looking for a book with params %s", searchParam)
		foundBook, err := finna.FindBookByTitle(searchParam)
		if err != nil {
			log.Print("No book found for title " + searchParam.Title)
		} else {
			log.Print("Book found for title " + searchParam.Title)
			fmt.Println(foundBook)
		}
		//Avoid spamming Finna api too much
		time.Sleep(500 * time.Millisecond)
	}
}
