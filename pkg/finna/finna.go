package finna

import (
	json2 "encoding/json"
	"fmt"
	"go/types"
	"io"
	"log"
	"net/http"
	url2 "net/url"
)

type SearchParameters struct {
	Title    string
	Building string
}

type Book struct {
	Title               string
	Id                  string
	NonPresenterAuthors []Authors
}

func (book Book) Url() string {
	return fmt.Sprintf("https://www.finna.fi/Record/%s", book.Id)
}

type Authors struct {
	Name string
	Role string
}

type finnaSearchResult struct {
	ResultCount int
	Records     []Book
}

const finnaApiUrl = "https://api.finna.fi/api/v1/search"
const finnaApiFilterBuilding = "filter[]=building:"
const finnaApiLookForTitle = "lookfor=title:"
const finnaApiFields = "field[]=title&field[]=id&field[]=nonPresenterAuthors&filter[]=format:0/Book/"

func FindBookByTitle(searchParameters SearchParameters) (Book, error) {

	url := getUrl(searchParameters)
	fmt.Println(url)
	request := get(url)
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		log.Fatalln("Unable to fetch book " + searchParameters.Title)
	}
	defer response.Body.Close()

	json, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var searchResult finnaSearchResult
	err2 := json2.Unmarshal(json, &searchResult)
	if err2 != nil {
		log.Fatalln("Error unmarshalling json")
	}
	// Since books are ordered by relevance, first result should be the matching one
	if searchResult.ResultCount < 1 {
		return Book{}, types.Error{
			Msg: "No matching book found",
		}
	}
	return searchResult.Records[0], nil
}

func getUrl(searchParameters SearchParameters) string {
	return fmt.Sprintf("https://api.finna.fi/api/v1/search?lookfor=title:%s&filter[]=building:%s&field[]=title&field[]=id&field[]=nonPresenterAuthors&filter[]=format:0/Book/",
		url2.QueryEscape(searchParameters.Title),
		url2.QueryEscape(searchParameters.Building))
}

func get(url string) *http.Request {
	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Set("Host", "api.finna.fi")
	request.Header.Set("Accept", "text/json")
	return request
}
