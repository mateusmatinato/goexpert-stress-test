package stresstest

import (
	"fmt"
	"time"
)

type Params struct {
	URL         string
	Concurrency int
	Requests    int
}

func (p *Params) PrintStart() {
	printLine()
	fmt.Printf("Starting Stress Test\n")
	fmt.Printf("URL: %s\n", p.URL)
	fmt.Printf("Concurrency: %d\n", p.Concurrency)
	fmt.Printf("Request Amount: %d\n", p.Requests)
	fmt.Printf("Wait for results...\n")
}

type HTTPStatusGroup map[int]int

type Response struct {
	TotalRequests   int
	Duration        time.Duration
	HTTPStatusGroup HTTPStatusGroup
}

type Job struct {
	URL string
}

type Result struct {
	HTTPStatus int
}

func (r *Response) PrintResult() {
	fmt.Printf("\n")
	printLine()
	fmt.Printf("Stress Test Results\n")
	fmt.Printf("Request Amount: %d\n", r.TotalRequests)
	fmt.Printf("Total Duration: %s\n", r.Duration)
	for status, count := range r.HTTPStatusGroup {
		fmt.Printf("[Status %d]: %d requests\n", status, count)
	}
	printLine()
}

func printLine() {
	fmt.Printf("------------------------------------\n")
}
