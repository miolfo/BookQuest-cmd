package goodreads

import "fmt"

type Book struct {
	title  string
	author string
	shelf  string
}

func ParseBooks(records [][]string) []Book {
	var books []Book
	for i, bookRecord := range records {
		if i == 0 {
			continue
		}
		book := Book{
			title:  bookRecord[1],
			author: bookRecord[2],
			shelf:  bookRecord[18],
		}
		books = append(books, book)
	}
	fmt.Println(books[1])
	return books
}
