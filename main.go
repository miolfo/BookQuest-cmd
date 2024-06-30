package main

import (
	"flag"
	"github.com/miolfo/BookQuest-cmd/cmd/bookquest"
	"github.com/miolfo/BookQuest-cmd/internal/gr_util"
	"log"
)

func main() {
	inputFile := flag.String("in", "", "Input file path and name")
	library := flag.String("library", "0/Helmet/", "Finna library code (i.e. 0/Helmet/)")
	outputFile := flag.String("out", "", "Output json file name")
	noStatusLog := flag.Bool("nostatuslog", false, "Suppress status log")

	flag.Parse()
	log.Printf("Logger: %t", *noStatusLog)
	if *noStatusLog {
		gr_util.DisableLogger()
	}

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
