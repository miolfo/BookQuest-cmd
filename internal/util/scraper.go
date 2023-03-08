package util

import (
	json2 "encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type ping struct {
	Message string
}

type availabilityResponse struct {
	//Possible types are 'ORDERED' | 'UNAVAILABLE' | 'AVAILABLE' | 'WAITING' | 'IN_TRANSIT' | 'UNKNOWN'
	Statuses []string
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

func IsBookAvailable(finnaId string) bool {
	if finnaId == "" {
		return false
	}
	log.Printf("Checking availability for book %s", finnaId)
	url := fmt.Sprintf("http://localhost:8080/api/availability?id=%s", finnaId)

	request := get(url)
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Printf("Error checking availability: %s", err)
		return false
	}
	defer response.Body.Close()
	json, err := io.ReadAll(response.Body)
	if err != nil {
		log.Printf("Error checking availability: %s", err)
		return false
	}
	var result availabilityResponse
	err2 := json2.Unmarshal(json, &result)
	if err2 != nil {
		log.Printf("Error checking availability: %s", err)
		return false
	}

	log.Printf("Availability results for %s: %s", finnaId, result)

	for _, status := range result.Statuses {
		if status == "AVAILABLE" {
			return true
		}
	}

	return false
}

func get(url string) *http.Request {
	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Set("Accept", "text/json")
	return request
}
