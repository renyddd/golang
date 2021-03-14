package main

import (
	"fmt"
	"sync"
	"time"
)

//ref: https://learnku.com/docs/gobyexample/2020/work-pool-worker/6285

func worker(id int, jobs <-chan int, result chan<- int, wg *sync.WaitGroup) {
	for j := range jobs {
		fmt.Println("worker", id, "processing job", j)
		//time.Sleep(1*time.Second)
		result <- j * 2
	}
	wg.Done()
}

func main() {
	// 注意 buffer 限制
	jobs := make(chan int, 1000000)
	results := make(chan int, 1000000)
	wg := sync.WaitGroup{}

	for w := 1; w <= 100; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	for j := 1; j <= 1000; j++ {
		jobs <- j
	}
	close(jobs)

	go func() {
		for {
			<-results
		}
	}()

	time.Sleep(3*time.Second)
	//wg.Wait()
}
