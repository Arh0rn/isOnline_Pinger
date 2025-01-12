package cli

import (
	"fmt"
	"github.com/Arh0rn/isOnline_Pinger/config"
	"github.com/Arh0rn/isOnline_Pinger/models"
	"github.com/Arh0rn/isOnline_Pinger/storage"
	"github.com/Arh0rn/isOnline_Pinger/workerpool"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const configPath = "config/config.json"

func RunCLI() {
	conf, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatal(err)
	}

	u, p := startMenu(conf)
	RunPool(u, p)
}

func startMenu(conf *config.Config) ([]models.Url, models.Parameters) {
	var db storage.DB
	db.NewDBfrom(conf)
	err := db.ConnectDB(conf)
	if err != nil {
		log.Fatal(err)
	}

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

	urls, err = db.GetUrls()
	if err != nil {
		log.Fatal(err)
	}
	parameters, err := db.GetParameters()
	if err != nil {
		log.Fatal(err)
	}
	return urls, parameters
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

func RunPool(u []models.Url, p models.Parameters) {
	fmt.Println("Pool started")
	defer fmt.Println("Exiting...")

	results := make(chan workerpool.Result)
	wp := workerpool.NewPool(p.Workers, p.Timeout, results)
	wp.Init()

	go generateJobs(wp, u, p.Interval)
	go processResults(results)

	gracefulShutdown(wp)
}

func generateJobs(wp *workerpool.Pool, urls []models.Url, interval int) {
	for {
		for _, url := range urls {
			wp.Push(url.URL)
		}
		time.Sleep(time.Duration(interval) * time.Second)
	}
}
func processResults(results chan workerpool.Result) {
	go func() {
		for result := range results {
			fmt.Println(result)
		}
	}()
}
func gracefulShutdown(wp *workerpool.Pool) {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop
	wp.Stop()
}
