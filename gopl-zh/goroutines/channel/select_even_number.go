package main

import "fmt"

func main() {
	ch := make(chan struct{}, 1)

	for i := 1; i < 1000; i++ {
		select {
		case <-ch:
			fmt.Println(i)
		case ch <- struct{}{}:
		}
	}
}
