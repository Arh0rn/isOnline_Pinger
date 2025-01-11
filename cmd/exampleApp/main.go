package main

import (
	"fmt"
	"github.com/Arh0rn/isOnline_Pinger/workerpool"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	interval       = 5
	RequestTimeout = 2
	WorkersCount   = 5
)

var urls = []string{
	"https://www.google.com",
	"https://www.github.com",
	"https://www.stackoverflow.com",
	"https://www.youtube.com",
}

func main() {
	fmt.Println("Strart")
	defer fmt.Println("End")

	results := make(chan workerpool.Result)
	wp := workerpool.NewPool(WorkersCount, RequestTimeout, results)
	wp.Init()

	go generateJobs(wp)
	go processResults(results)

	gracefulShutdown(wp)
}

func processResults(results chan workerpool.Result) {
	go func() {
		for result := range results {
			fmt.Println(result)
		}
	}()
}

func generateJobs(wp *workerpool.Pool) {
	for {
		for _, url := range urls {
			wp.Push(url)
		}
		time.Sleep(interval * time.Second)
	}
}

func gracefulShutdown(wp *workerpool.Pool) {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop
	wp.Stop()
}
