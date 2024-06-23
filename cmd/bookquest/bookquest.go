package bookquest

import (
	"fmt"
	"github.com/miolfo/BookQuest-cmd/internal/util"
	"github.com/miolfo/BookQuest-cmd/pkg/finna"
	"github.com/miolfo/BookQuest-cmd/pkg/goodreads"
	"log"
	"time"
)

type BookPair struct {
	finnaBook finna.Book
	grBook    goodreads.Book
}

func Run(path string, building string, outPath string) {
	records := util.ReadCsvFromPath(path)
	books := goodreads.ParseBooks(records)
	booksToRead := goodreads.FilterByShelf(books, "to-read")
	var searchParams []finna.SearchParameters
	for _, book := range booksToRead {
		searchParams = append(searchParams, finna.SearchParameters{
			Title:    book.Title,
			Building: building,
			Author:   book.Author,
		})
	}

	bookPairs := findBooks(searchParams, booksToRead)
	result := util.BookSearchResults{Results: bookPairs}
	if util.IsScraperRunning() {
		log.Printf("Backend is running, fetching info from backend apis...")
		result = addScrapingResult(result)
		result = addEbooksComResult(result)
	} else {
		log.Printf("Scraper not running, not scraping availability info")
	}
	util.WriteResultsToPath(result, outPath)
	log.Printf("Wrote results to file %s", outPath)
}

func findBooks(searchParams []finna.SearchParameters, booksToRead []goodreads.Book) []util.BookSearchResult {
	var bookSearchResults []util.BookSearchResult
	for i, searchParam := range searchParams {
		log.Printf("Looking for a book with params %s", searchParam)
		foundBooks, err := finna.FindBookByTitle(searchParam)
		if err != nil {
			log.Print("No book found for title " + searchParam.Title)
			bookSearchResults = append(bookSearchResults, util.BookSearchResult{
				Title:     booksToRead[i].Title,
				Author:    booksToRead[i].Author,
				Available: []bool{},
				Urls:      []string{},
			})
		} else {
			log.Print("Book found for title " + searchParam.Title)
			var urls []string
			var finnaIds []string
			for _, foundBook := range foundBooks {
				urls = append(urls, foundBook.Url())
				finnaIds = append(finnaIds, foundBook.Id)
			}
			bookSearchResults = append(bookSearchResults, util.BookSearchResult{
				Title:     booksToRead[i].Title,
				Author:    booksToRead[i].Author,
				Available: []bool{},
				FinnaIds:  finnaIds,
				Urls:      urls,
			})
		}
		//Avoid spamming Finna api too much
		time.Sleep(500 * time.Millisecond)
	}
	return bookSearchResults
}

func addScrapingResult(results util.BookSearchResults) util.BookSearchResults {
	res := util.BookSearchResults{Results: []util.BookSearchResult{}}
	for _, result := range results.Results {
		var statuses []bool
		for _, finnaId := range result.FinnaIds {
			status := util.IsBookAvailable(finnaId)
			statuses = append(statuses, status)
		}
		res.Results = append(res.Results, util.BookSearchResult{
			Title:     result.Title,
			Author:    result.Author,
			FinnaIds:  result.FinnaIds,
			Available: statuses,
			Urls:      result.Urls,
		})
	}
	return res
}

func addEbooksComResult(results util.BookSearchResults) util.BookSearchResults {
	for _, result := range results.Results {
		ebooksResponse := util.GetBookInfo(result.Title, result.Author)
		fmt.Println(ebooksResponse)
	}
	return results
}
