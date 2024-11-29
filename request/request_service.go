package request

import (
	"fmt"
	"net/http"
	netUrl "net/url"
	"strings"
	"sync"
	"time"
)

type RequestService struct{}

func NewRequestService() *RequestService {
	return &RequestService{}
}

func (rs *RequestService) MakeRequests(url string, totalRequests, concurrency int) {
	if !isValidUrl(url) {
		return
	}
	report := StartReport()
	wg := sync.WaitGroup{}
	wg.Add(totalRequests)
	requests := make(chan int, concurrency)
	results := make(chan int, totalRequests)

	for i := 0; i < totalRequests; i++ {
		requests <- i
		fmt.Printf("\r%d/%d", i+1, totalRequests)
		go func() {
			defer wg.Done()
			rs.makeRequest(url, requests, results)
		}()
	}

	wg.Wait()
	finishedAt := time.Now()
	close(requests)
	close(results)

	resultsMap := make(map[int]int)
	for result := range results {
		resultsMap[result]++
	}

	report.FinishReport(finishedAt, totalRequests, resultsMap)
	report.BeautyPrint()
}

func (rs *RequestService) makeRequest(url string, requests <-chan int, results chan<- int) {
	defer func() {
		<-requests
	}()

	resp, err := http.DefaultClient.Get(url)
	if err != nil {
		results <- 0
		return
	}
	defer resp.Body.Close()
	results <- resp.StatusCode
}

func isValidUrl(url string) bool {
	if _, err := netUrl.ParseRequestURI(url); err != nil {
		if strings.HasPrefix("http://", url) || strings.HasPrefix("https://", url) {
			fmt.Println("Please provide a valid URL")
			return false
		}
		fmt.Println("Please provide a valid URL with http:// or https://")
		return false
	}
	return true
}
