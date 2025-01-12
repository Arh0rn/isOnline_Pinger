package cli

import (
	"fmt"
	"github.com/Arh0rn/isOnline_Pinger/models"
)

func PrintUrls(urls []models.Url) {
	for _, url := range urls {
		fmt.Print(url)
	}
}

func NewParameters(timeout, interval, workers int) models.Parameters {
	return models.Parameters{
		Timeout:  timeout,
		Interval: interval,
		Workers:  workers,
	}
}
