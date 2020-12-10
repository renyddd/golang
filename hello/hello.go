package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	live5seconds := func() {
		log.Println("sleep")
		time.Sleep(5*time.Second)
	}

	go func() {
		defer wg.Done()
		err := http.ListenAndServe("0.0.0.0:6060", nil)
		if err != nil {
			log.Fatalln(err)
		}
	}()

	done := make(chan bool)
	ticker := time.NewTicker(3*time.Second)
	defer ticker.Stop()
	for {
		go live5seconds()
		select {
		case <- ticker.C:
			log.Println("tick ...")
			continue
		case <- done:
			break
		}
	}



	wg.Wait()
}