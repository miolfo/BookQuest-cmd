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

	bookPairs := findBookPairs(searchParams, booksToRead)

	if util.IsScraperRunning() {
		log.Printf("Scraper is running, scraping availability info from finna...")
		bookPairs = addScrapingResult(bookPairs)
	} else {
		log.Printf("Scraper not running, not scraping availability info")
	}

	csvRecords := convertToRecords(bookPairs)
	util.WriteRecordsToPath(csvRecords, outPath)
	log.Printf("Wrote results to file %s", outPath)
}

func addScrapingResult(pairs []BookPair) []BookPair {
	var scrapedPairs []BookPair
	for _, pair := range pairs {
		isAvailable := util.IsBookAvailable(pair.finnaBook.Id)
		scrapedPairs = append(scrapedPairs, BookPair{
			finnaBook: finna.Book{
				Title:               pair.finnaBook.Title,
				Id:                  pair.finnaBook.Id,
				NonPresenterAuthors: pair.finnaBook.NonPresenterAuthors,
				Available:           isAvailable,
			},
			grBook: pair.grBook,
		})
	}
	return scrapedPairs
}

func findBookPairs(searchParams []finna.SearchParameters, booksToRead []goodreads.Book) []BookPair {
	var bookPairs []BookPair
	for i, searchParam := range searchParams {
		log.Printf("Looking for a book with params %s", searchParam)
		foundBook, err := finna.FindBookByTitle(searchParam)
		if err != nil {
			log.Print("No book found for title " + searchParam.Title)
			bookPairs = append(bookPairs, BookPair{
				finnaBook: finna.Book{},
				grBook:    booksToRead[i],
			})
		} else {
			log.Print("Book found for title " + searchParam.Title)
			bookPairs = append(bookPairs, BookPair{
				finnaBook: foundBook,
				grBook:    booksToRead[i],
			})
		}
		//Avoid spamming Finna api too much
		time.Sleep(500 * time.Millisecond)
	}
	return bookPairs
}

func convertToRecords(bookPairs []BookPair) [][]string {
	var records [][]string
	for _, pair := range bookPairs {
		records = append(records, convertToRecord(pair))
	}
	return records
}

func convertToRecord(bookPair BookPair) []string {
	availabilityString := ""
	if bookPair.finnaBook.Available {
		availabilityString = "AVAILABLE"
	}
	if bookPair.finnaBook.Id != "" {
		return []string{bookPair.grBook.Title, bookPair.grBook.Author, bookPair.finnaBook.Url(), availabilityString}
	} else {
		return []string{bookPair.grBook.Title, bookPair.grBook.Author, "", availabilityString}
	}
}
