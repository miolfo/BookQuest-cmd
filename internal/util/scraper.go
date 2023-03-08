package util

import (
	json2 "encoding/json"
	"io"
	"log"
	"net/http"
)

type ping struct {
	Message string
}

func IsScraperRunning() bool {
	url := "http://localhost:8080/api/ping"
	request := get(url)
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Printf("Scraper not responding in url %s", url)
		return false
	}
	defer response.Body.Close()
	json, err := io.ReadAll(response.Body)
	if err != nil {
		log.Printf("Scraper not responding with proper data in url %s", url)
		return false
	}
	var pingresult ping
	err2 := json2.Unmarshal(json, &pingresult)
	if err2 != nil {
		log.Printf("Scraper not responding with proper data in url %s", url)
		return false
	}
	return pingresult.Message == "pong"
}

func get(url string) *http.Request {
	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Set("Accept", "text/json")
	return request
}
