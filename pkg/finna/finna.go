package finna

import (
	json2 "encoding/json"
	"fmt"
	util "github.com/miolfo/BookQuest-cmd/pkg/util"
	"go/types"
	"io"
	"log"
	"net/http"
	url2 "net/url"
	"strings"
)

type SearchParameters struct {
	Title    string
	Building string
	Author   string
}

type Book struct {
	Title               string
	Id                  string
	NonPresenterAuthors []Authors
	Available           bool
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

func FindBookByTitle(searchParameters SearchParameters) ([]Book, error) {

	//TODO: If first search return no results, try again with author -> author2, or without author
	url := getUrl(searchParameters)
	log.Println(url)
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
		return []Book{}, types.Error{
			Msg: "No matching book found",
		}
	}

	//Remove all whitespace
	strippedSearchTitle := util.StripTitle(searchParameters.Title, true, false)
	var results []Book
	for _, record := range searchResult.Records {
		strippedTitle := util.StripTitle(record.Title, true, false)
		if strings.Contains(strippedTitle, strippedSearchTitle) || strings.Contains(strippedSearchTitle, strippedTitle) {
			results = append(results, record)
		}
	}

	return results, nil
}

func getUrl(searchParameters SearchParameters) string {
	return fmt.Sprintf("https://api.finna.fi/api/v1/search?lookfor=title:%s&filter[]=building:%s&filter[]=author:%s&field[]=title&field[]=id&field[]=nonPresenterAuthors&filter[]=format:0/Book/",
		url2.QueryEscape(searchParameters.Title),
		url2.QueryEscape(searchParameters.Building),
		url2.QueryEscape(searchParameters.Author))
}

func get(url string) *http.Request {
	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Set("Host", "api.finna.fi")
	request.Header.Set("Accept", "text/json")
	return request
}
