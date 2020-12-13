package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

type result struct {
	value interface{}
	err   error
}

type entry struct {
	res result
	ready chan struct{}
}

type Func func(key string) (interface{}, error)

// Memo caches the result of calling a Func.
type Memo struct {
	f           Func
	mu          sync.Mutex
	cache       map[string]*entry
	cacheStatus chan string
	usedTime chan string
}

func New(f Func) *Memo {
	return &Memo{
		f:           f,
		cache:       make(map[string]*entry),
		cacheStatus: make(chan string, 1),
		usedTime: make(chan string, 1),
	}
}

func (memo *Memo) Get(key string) (interface{}, error) {
	memo.mu.Lock()
	e := memo.cache[key]
	var start time.Time
	if e == nil {
		e = &entry{ready: make(chan struct{})}
		memo.cache[key] = e
		memo.mu.Unlock()

		memo.cacheStatus <- "Uncached -> "
		start = time.Now()
		e.res.value, e.res.err = memo.f(key)
		close(e.ready)
	} else {
		memo.mu.Unlock()
		start = time.Now()
		memo.cacheStatus <- "Cached -> "
		<-e.ready
	}
	memo.usedTime <- time.Since(start).String()
	return e.res.value, e.res.err
}

func main() {
	m := New(httpGetBody)
	var wg sync.WaitGroup
	for _, url := range incomingURLs() {
		wg.Add(1)
		go func(url string) {
			value, err := m.Get(url)
			if err != nil {
				log.Fatalln(err)
			}
			fmt.Println(<-m.cacheStatus, <-m.usedTime, url,
				len(string(value.([]byte))))
			wg.Done()
		}(url)
	}
	wg.Wait()
}

func httpGetBody(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func incomingURLs() []string {
	return []string{
		"https://www.bytedance.com",
		"https://www.qq.com",
		"https://www.163.com",
		"https://www.bilibili.com",
		"https://www.bytedance.com",
		"https://www.bytedance.com",
		"https://www.bilibili.com",
		"https://www.bilibili.com",
	}
}