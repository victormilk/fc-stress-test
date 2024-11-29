package request

import (
	"fmt"
	"strings"
	"time"
)

type Report struct {
	StartedAt     time.Time
	FinishedAt    time.Time
	TotalTime     float64
	TotalRequests int
	ResultsStatus map[int]int
}

func StartReport() *Report {
	return &Report{
		StartedAt: time.Now(),
	}
}

func (r *Report) FinishReport(finishedAt time.Time, totalRequests int, results map[int]int) {
	r.FinishedAt = time.Now()
	r.TotalTime = r.FinishedAt.Sub(r.StartedAt).Seconds()
	r.TotalRequests = totalRequests
	r.ResultsStatus = results
}

func (r *Report) BeautyPrint() {
	fmt.Printf("\n\n%s\n", strings.Repeat("-", 30))
	fmt.Printf("Total execution time: %.2fs\n", r.TotalTime)
	fmt.Printf("Total requests: %d\n", r.TotalRequests)
	for status, count := range r.ResultsStatus {
		if status == 0 {
			fmt.Printf("Requests with unknown status: %d\n", count)
		} else {
			fmt.Printf("Requests with status code %d: %d\n", status, count)
		}
	}
	fmt.Println(strings.Repeat("-", 30))
}
