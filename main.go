package main

import (
	"github.com/miolfo/goodreads-finna/cmd/finnagr"
	"os"
)

func main() {
	finnagr.Finnagr(os.Args[1])
}
