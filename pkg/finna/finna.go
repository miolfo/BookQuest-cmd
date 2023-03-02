package finna

import (
	"fmt"
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

	url := fmt.Sprintf("%s?%s%s&%s%s&sort=relevance:id asc&page=1",
		finnaApiUrl,
		finnaApiLookForTitle,
		url2.QueryEscape(searchParameters.Title),
		finnaApiFilterBuilding,
		searchParameters.Building)
	fmt.Println(url)
}
