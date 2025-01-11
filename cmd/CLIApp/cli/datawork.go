package cli

import (
	"fmt"
	"github.com/Arh0rn/isOnline_Pinger/storage"
)

func PrintUrls(urls []storage.Url) {
	for _, url := range urls {
		fmt.Print(url)
	}
}

func NewParameters(timeout, interval, workers int) storage.Parameters {
	return storage.Parameters{
		Timeout:  timeout,
		Interval: interval,
		Workers:  workers,
	}
}
