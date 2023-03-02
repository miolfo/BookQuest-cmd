package finnagr

import (
	"github.com/miolfo/goodreads-finna/internal/util"
	"github.com/miolfo/goodreads-finna/pkg/finna"
	"github.com/miolfo/goodreads-finna/pkg/goodreads"
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

	finna.FindBookByTitle(searchParams[0])
}
