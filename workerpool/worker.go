package workerpool

import (
	"net/http"
	"time"
)

type worker struct {
	client *http.Client
}

func newWorker(timeout time.Duration) *worker {
	return &worker{
		client: &http.Client{
			Timeout: timeout,
		},
	}
}

func (w worker) process(url string) Result {
	result := Result{URL: url, Time: time.Now(), isOk: true}
	resp, err := w.client.Get(url)

	if err != nil || resp == nil || resp.StatusCode != http.StatusOK {
		result.Status = "Error"
		result.isOk = false
	} else {
		err = resp.Body.Close()
		result.Status = resp.Status
	}

	return result
}
