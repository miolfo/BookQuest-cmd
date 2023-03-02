package finna

import (
	"fmt"
	"io"
	"log"
	"net/http"
	url2 "net/url"
)

type SearchParameters struct {
	Title    string
	Building string
}

const finnaApiUrl = "https://api.finna.fi/api/v1/search"
const finnaApiFilterBuilding = "filter[]=building:"
const finnaApiLookForTitle = "lookfor=title:"

func FindBookByTitle(searchParameters SearchParameters) {

	url := fmt.Sprintf("%s?%s%s&%s%s",
		finnaApiUrl,
		finnaApiLookForTitle,
		url2.QueryEscape(searchParameters.Title),
		finnaApiFilterBuilding,
		url2.QueryEscape(searchParameters.Building))
	fmt.Println(url)
	request := get(url)
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		log.Print("Unable to fetch book " + searchParameters.Title)
		return
	}
	defer response.Body.Close()

	json, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(json))
}

func get(url string) *http.Request {
	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Set("Host", "api.finna.fi")
	request.Header.Set("Accept", "text/json")
	return request
}
