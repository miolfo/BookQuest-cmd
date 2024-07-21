package ebookscom

import (
	json2 "encoding/json"
	"fmt"
	"github.com/miolfo/BookQuest-cmd/pkg/util"
	"io"
	"log"
	"net/http"
	url2 "net/url"
)

type EbooksComResponse struct {
	TotalResults int
	Results      []EbooksComResult
}

type EbooksComResult struct {
	Title         string
	StorefrontUrl string
	Authors       []EbooksComAuthor
	Price         EbooksComPrice
}

type EbooksComAuthor struct {
	Name string
}

type EbooksComPrice struct {
	Currency string
	Value    float32
}

func GetEbooksComInfo(title string, author string) EbooksComResponse {
	strippedTitle := util.StripTitle(title, false)
	url := fmt.Sprintf("https://api.ebooks.com/v2/FI/book/search?title=%s&author=%s", url2.QueryEscape(strippedTitle), url2.QueryEscape(author))
	log.Printf("Searching ebooks.com for %s, %s from path %s", strippedTitle, author, url)
	request := getEbook(url)
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Printf("Error checking ebooks.com: %s", err)
		return EbooksComResponse{}
	}
	defer response.Body.Close()
	json, err := io.ReadAll(response.Body)
	if err != nil {
		log.Printf("Error checking ebooks.com: %s", err)
		return EbooksComResponse{}
	}
	var result EbooksComResponse
	err2 := json2.Unmarshal(json, &result)
	if err2 != nil {
		log.Printf("Error checking ebooks.com: %s", err)
		return EbooksComResponse{}
	}
	return result
}

func getEbook(url string) *http.Request {
	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")
	return request
}
