package miniSDK

import (
	"fmt"
	"github.com/Arh0rn/isOnline_Pinger/models"
	"github.com/Arh0rn/isOnline_Pinger/workerpool"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func RunPool(u []models.Url, p models.Parameters) {
	fmt.Println("Pool started")
	defer fmt.Println("Exit")

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
