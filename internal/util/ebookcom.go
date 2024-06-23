package util

import (
	"bytes"
	json2 "encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type EbooksComResponse struct {
	Title  string
	Author string
}

func GetBookInfo(title string, author string) EbooksComResponse {
	url := "http://localhost:3000/api/search-ebookcom"
	log.Printf("Searching ebooks.com for %s, %s", title, author)
	request := post(url, title, author)
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
	log.Printf("Ebooks com response for %s, %s: %s", title, author, result)
	return result
}

func post(url string, title string, author string) *http.Request {
	body := fmt.Sprintf("{\"author\": \"%s\", \"title\": \"%s\"}", author, title)
	jsonStr := []byte(body)
	log.Printf(body)
	request, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")
	return request
}
