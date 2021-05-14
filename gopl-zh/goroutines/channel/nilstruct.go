package main

import (
	"fmt"
)

func main() {
	go main1()
}

func main1() {
	ch := make(chan struct{})
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("Processing", i)
			ch <- struct{}{}
		}(i)
	}

	for range ch {
		<-ch
	}
}
