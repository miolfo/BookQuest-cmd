# BookQuest-cmd
Integration between finna and goodreads to check what books are available in a specific location. Can be used to read an exported .csv file from goodreads,
and then checking if they can be found from a local library using Finna public api, as well as finding a price for the books from ebooks.com api. The tool only checks books in your want-to-read -shelf.

Usage

`
go run . ./resources/goodreads_library_export.csv 0/Helmet/ ./resources/result.json
`

First parameter is the name of the input file. Second parameter is the building to look for (for example, 0/Helmet/ includes all Helsinki city libraries).
Third parameter is the output file.
Output generates json objects such as 

`
{
	"Title": "Seveneves",
	"Author": "Stephenson, Neal",
	"FinnaIds": [
		"helmet.2539843",
		"helmet.2252762"
	],
	"Available": [
		false,
		true
	],
	"Urls": [
		"https://www.finna.fi/Record/helmet.2539843",
		"https://www.finna.fi/Record/helmet.2252762",
		"https://www.ebooks.com/en-fi/book/1745874/seveneves/neal-stephenson/"
	],
	"Price": {
		"Currency": "EUR",
		"Value": 12.29
	}
}
`

Urls contain all relevant URLs for the title, Price contains the price fetched from ebooks.com api, 
and Available contains scraped info about book's library availability in finna -urls

