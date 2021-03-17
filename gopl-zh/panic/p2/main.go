package main

import (
	"log"
	"time"
)

// ref: https://stackoverflow.com/questions/47808360/how-does-a-caller-function-to-recover-from-child-goroutines-panics
// 为什么父进程服务 recover 子进程的 panic
// A goroutine cannot recover from a panic in another goroutine.

func fson() {
	panic("fson die")
}

func ffather() {
	defer func() {
		if p := recover(); p != nil {
			log.Println(p)
		}
	}()

	go fson()

	time.Sleep(1*time.Second)
}

func protect(g func()) {
	defer func() {
		log.Println("done")  // Println executes normally even if there is a panic
		if x := recover(); x != nil {
			log.Printf("run time panic: %v", x)
		}
	}()
	log.Println("start")
	g()
}

func main() {
	protect(func() {
		a := []int{1,2}
		_ = a[5]
	})
}