package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println("Commencing countdown.")
	abort := make(chan struct{})

	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()

	select {
	// select 只是会简答的阻塞在这里而已
	case <-time.After(5 * time.Second):
	case <-abort:
		fmt.Println("aborted")
		return

	}

}
