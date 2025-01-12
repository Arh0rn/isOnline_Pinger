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
func PrintInfo() {
	fmt.Println("isOnline_Pinger")
	fmt.Println("Choose an option:")
	fmt.Println("1. show urls")
	fmt.Println("2. add url")
	fmt.Println("3. delete url")
	fmt.Println("4. show parameters")
	fmt.Println("5. edit parameters")
	fmt.Println("6. start")
}
