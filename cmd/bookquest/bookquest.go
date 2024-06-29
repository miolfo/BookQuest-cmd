package bookquest

import (
	"fmt"
	"github.com/miolfo/BookQuest-cmd/internal/util"
	"github.com/miolfo/BookQuest-cmd/pkg/finna"
	"github.com/miolfo/BookQuest-cmd/pkg/goodreads"
	"log"
	"slices"
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
	result = addScrapingResult(&result)
	result = addEbooksComResult(&result)
	fmt.Println(result)
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

func addScrapingResult(results *util.BookSearchResults) util.BookSearchResults {
	var finnaIdList []string
	for _, result := range results.Results {
		for _, finnaId := range result.FinnaIds {
			finnaIdList = append(finnaIdList, finnaId)
		}
	}

	availabilities := util.AreBooksAvailable(finnaIdList)
	for _, availabilityResult := range availabilities {
		for bsIdx, bookSearchResult := range results.Results {
			finnaIdIndex := slices.Index(bookSearchResult.FinnaIds, availabilityResult.FinnaId)
			if finnaIdIndex > -1 {
				results.Results[bsIdx].Available = append(results.Results[bsIdx].Available, availabilityResult.Available)
			}
		}
	}

	return *results
}

func addEbooksComResult(results *util.BookSearchResults) util.BookSearchResults {
	for i, result := range results.Results {
		ebooksResponse := util.GetEbooksComInfo(result.Title, result.Author)
		if ebooksResponse.TotalResults > 0 {
			ebooksBook := ebooksResponse.Results[0]
			results.Results[i].Urls = append(result.Urls, ebooksBook.StorefrontUrl)
			results.Results[i].Price = util.PriceWithCurrency{
				Value:    ebooksBook.Price.Value,
				Currency: ebooksBook.Price.Currency,
			}
		}
	}
	return *results
}
