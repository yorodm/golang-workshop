package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	numPiecesOfWork := 20
	numWorkers := 5
	workCh := make(chan int)
	wg := &sync.WaitGroup{}
	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go worker(workCh, wg)
	}
	for i := 0; i < numPiecesOfWork; i++ {
		work := i % 10
		workCh <- work
	}
	close(workCh)
	v := <-workCh
	fmt.Println(v)
	wg.Wait()
	fmt.Println("done")
}

func worker(workCh <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for work := range workCh {
		doWork(work)
	}
}

func doWork(work int) {
	time.Sleep(time.Duration(work) * time.Millisecond)
	fmt.Println("slept for", work, "milliseconds")
}
