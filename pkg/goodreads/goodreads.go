package goodreads

type Book struct {
	Title  string
	Author string
	Shelf  string
}

func ParseBooks(records [][]string) []Book {
	var books []Book
	for i, bookRecord := range records {
		if i == 0 {
			continue
		}
		book := Book{
			Title:  bookRecord[1],
			Author: bookRecord[3],
			Shelf:  bookRecord[18],
		}
		books = append(books, book)
	}
	return books
}

func FilterByShelf(books []Book, shelf string) []Book {
	correctShelf := []Book{}
	for _, book := range books {
		if book.Shelf == shelf {
			correctShelf = append(correctShelf, book)
		}
	}
	return correctShelf
}
