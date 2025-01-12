package cli

import (
	"fmt"
	"github.com/Arh0rn/isOnline_Pinger/config"
	"github.com/Arh0rn/isOnline_Pinger/models"
	"github.com/Arh0rn/isOnline_Pinger/storage"
	_ "github.com/Arh0rn/isOnline_Pinger/storage/mongo"
	_ "github.com/Arh0rn/isOnline_Pinger/storage/postgres"
	"log"
)

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
func OpenDB(conf *config.Config) (storage.DB, error) {
	db, err := storage.NewDBfrom(*conf)
	if err != nil {
		return nil, err
	}
	err = db.ConnectDB(*conf)
	if err != nil {
		return nil, err
	}
	return db, nil
}
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
func startMenu(db storage.DB) ([]models.Url, models.Parameters) {

	var input int
	PrintInfo()

	var urls []models.Url
	var out bool = false
	for {
		if out {
			break
		}
		_, err := fmt.Scan(&input)
		if err != nil {
			fmt.Println(err)
		}
		switch input {
		case 1:
			urls, err := db.GetUrls()
			if err != nil {
				log.Fatal(err)
			}
			PrintUrls(urls)
		case 2:
			var inputUrl string
			_, err := fmt.Scan(&inputUrl)
			if err != nil {
				log.Fatal(err)
			}
			err = db.AddUrl(inputUrl)
			if err != nil {
				log.Fatal(err)
			}
		case 3:
			var inputId int
			_, err := fmt.Scan(&inputId)
			if err != nil {
				log.Fatal(err)
			}
			err = db.DeleteUrl(inputId)
			if err != nil {
				log.Fatal(err)
			}
		case 4:
			parameters, err := db.GetParameters()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(parameters)
		case 5:
			var timeout, interval, workers int
			fmt.Println("Enter timeout, interval and workers separated by space")
			_, err := fmt.Scan(&timeout, &interval, &workers)
			if err != nil {
				log.Fatal(err)
			}
			p := NewParameters(timeout, interval, workers)
			err = db.SetParameters(p)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("Parameters updated")
		case 6:
			out = true
		}
	}

	urls, err := db.GetUrls()
	if err != nil {
		log.Fatal(err)
	}
	parameters, err := db.GetParameters()
	if err != nil {
		log.Fatal(err)
	}
	return urls, parameters
}
