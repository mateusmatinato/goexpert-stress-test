package stresstest

import (
	"time"

	"github.com/mateusmatinato/goexpert-stress-test/internal/platform/http"
	"github.com/schollz/progressbar/v3"
)

type UseCase interface {
	Execute(params Params) Response
}

type useCase struct {
}

func (u useCase) Execute(params Params) Response {
	start := time.Now()

	numJobs := params.Requests

	jobs := make(chan Job, numJobs)
	results := make(chan Result, numJobs)

	for w := 1; w <= params.Concurrency; w++ {
		go worker(jobs, results)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- Job{URL: params.URL}
	}
	close(jobs)

	var resp Response
	resp.HTTPStatusGroup = make(HTTPStatusGroup)

	bar := progressbar.New(numJobs)
	for a := 1; a <= numJobs; a++ {
		result := <-results
		resp.HTTPStatusGroup[result.HTTPStatus]++
		_ = bar.Add(1)
	}

	resp.TotalRequests = numJobs
	resp.Duration = time.Since(start)

	return resp
}

func worker(jobs chan Job, results chan Result) {
	for job := range jobs {
		status, err := http.ExecuteRequest(job.URL)
		if err != nil {
			status = 500
		}
		results <- Result{HTTPStatus: status}
	}
}

func NewUseCase() UseCase {
	return &useCase{}
}
