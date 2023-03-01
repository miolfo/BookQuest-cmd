package finnagr

import (
	"github.com/miolfo/goodreads-finna/internal/util"
	"github.com/miolfo/goodreads-finna/pkg/goodreads"
)

func Finnagr(path string) {
	records := util.ReadCsvFromPath(path)
	goodreads.ParseBooks(records)
}
