# BookQuest-cmd
Integration between finna and goodreads to check what books are available in a specific location. Can be used to read an exported .csv file from goodreads,
and then checking if they can be found from a local library using Finna public api. The tool only checks books in your want-to-read -shelf.

Usage

`
go run . ./resources/goodreads_library_export.csv 0/Helmet/ ./resources/result.json
`

First parameter is the name of the input file. Second parameter is the building to look for (for example, 0/Helmet/ includes all Helsinki city libraries).
Third parameter is the output file.
Output generates json objects such as 

`
		{
			"Title": "Children of Dune (Dune, #3)",
			"Author": "Herbert, Frank",
			"Status": false,
			"Urls": [
				"https://www.finna.fi/Record/helmet.2404461",
				"https://www.finna.fi/Record/helmet.2511280"
			]
		}
`

with Urls containing all found locations for the book

## Scraper

If you are also running the Selenium based finna -scraper (https://github.com/miolfo/finna-scraper-server)
on the same machine in port 8080, finna-goodreads tool will also check if the book is available
for loaning. If the book is available, the example row will look like this

`
"Children of Dune (Dune, #3)",Frank Herbert,https://www.finna.fi/Record/helmet.2404461,AVAILABLE
`
