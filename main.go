package main

import (
	"flag"
	"github.com/miolfo/BookQuest-cmd/cmd/bookquest"
	"log"
)

func main() {
	inputFile := flag.String("in", "", "Input file path and name")
	library := flag.String("library", "0/Helmet/", "Finna library code (i.e. 0/Helmet/)")
	outputFile := flag.String("out", "", "Output json file name")
	flag.Parse()
	if len(*inputFile) == 0 {
		log.Printf("Missing in -flag")
		return
	}
	if len(*outputFile) == 0 {
		log.Printf("Missing out -flag")
		return
	}
	bookquest.Run(*inputFile, *library, *outputFile)
}
