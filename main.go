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
	statusLogFile := flag.String("statusOut", "", "Output status json file for following progress")
	flag.Parse()

	if len(*statusLogFile) > 0 {
		gr_util.InitLogger(*statusLogFile)
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
