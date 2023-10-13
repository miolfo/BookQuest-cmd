package finnagr

import (
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

func Finnagr(path string, building string, outPath string) {
	records := util.ReadCsvFromPath(path)
	books := goodreads.ParseBooks(records)
	booksToRead := goodreads.FilterByShelf(books, "to-read")
	var searchParams []finna.SearchParameters
	for _, book := range booksToRead {
		searchParams = append(searchParams, finna.SearchParameters{
			Title:    book.Title,
			Building: building,
		})
	}

	bookPairs := findBooks(searchParams, booksToRead)

	util.WriteResultsToPath(util.BookSearchResults{Results: bookPairs}, outPath)
	log.Printf("Wrote results to file %s", outPath)
}

func findBooks(searchParams []finna.SearchParameters, booksToRead []goodreads.Book) []util.BookSearchResult {
	var bookSearchResults []util.BookSearchResult
	for i, searchParam := range searchParams {
		log.Printf("Looking for a book with params %s", searchParam)
		foundBook, err := finna.FindBookByTitle(searchParam)
		if err != nil {
			log.Print("No book found for title " + searchParam.Title)
			bookSearchResults = append(bookSearchResults, util.BookSearchResult{
				Title:  booksToRead[i].Title,
				Author: booksToRead[i].Author,
				Status: false,
				Urls:   []string{},
			})
		} else {
			log.Print("Book found for title " + searchParam.Title)
			bookSearchResults = append(bookSearchResults, util.BookSearchResult{
				Title:  booksToRead[i].Title,
				Author: booksToRead[i].Author,
				Status: false,
				Urls:   []string{foundBook.Url()},
			})
		}
		//Avoid spamming Finna api too much
		time.Sleep(500 * time.Millisecond)
	}
	return bookSearchResults
}
