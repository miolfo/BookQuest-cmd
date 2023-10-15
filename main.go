package main

import (
	"github.com/miolfo/BookQuest-cmd/cmd/bookquest"
	"os"
)

func main() {
	bookquest.Run(os.Args[1], os.Args[2], os.Args[3])
}
