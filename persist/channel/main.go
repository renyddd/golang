package main

import (
	"log"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	errorChan := make(chan int, 0)

	go func() {
		GenerateDateWithError(errorChan)
		wg.Done()
	}()

	go func() {
		HandleError(errorChan)
		wg.Done()
	}()

	wg.Wait()
}

func GenerateDateWithError(ch chan<- int) {
	for i := 0; i < 10000; i++ {
		if i % 1340 == 0 {
			time.Sleep(100000 * time.Microsecond)
			ch <- i
			log.Println("Error occur with:", i)
		}
	}
	close(ch)
}

func HandleError(ch <-chan int) {
	for v := range ch {
		log.Println("Handled error:", v)
	}
}