package workerpool

import (
	"sync"
	"time"
)

type Pool struct {
	worker      *worker
	workerCount int
	timeout     time.Duration

	urls    chan string
	results chan Result

	wg        *sync.WaitGroup
	isStopped bool
}

func NewPool(wc int, to int, results chan Result) *Pool {
	timeout := time.Duration(to) * time.Second
	return &Pool{
		worker:      newWorker(timeout),
		workerCount: wc,
		timeout:     timeout,

		urls:    make(chan string),
		results: results,

		wg:        new(sync.WaitGroup),
		isStopped: false,
	}
}

func (p *Pool) Init() {
	for i := 0; i < p.workerCount; i++ {
		go p.initWorker()
	}
}

func (p *Pool) Push(url string) {
	if p.isStopped {
		return
	}

	p.urls <- url
	p.wg.Add(1)
}

func (p *Pool) Stop() {
	p.isStopped = true
	p.wg.Wait()
	close(p.urls)
	close(p.results)
}

func (p *Pool) initWorker() {
	for url := range p.urls {
		p.results <- p.worker.process(url)
		p.wg.Done()
	}
}
