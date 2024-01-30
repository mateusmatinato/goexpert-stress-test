package http

import (
	"net/http"
	"time"
)

func ExecuteRequest(url string) (int, error) {
	client := http.Client{
		Timeout: 3 * time.Second,
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return 0, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}

	return resp.StatusCode, nil
}
