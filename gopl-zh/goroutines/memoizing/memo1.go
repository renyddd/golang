package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

// Memo caches the result of calling a Func.
type Memo struct {
	f           Func
	mu          sync.Mutex
	cache       map[string]result
	cacheStatus chan string
}

// Func is the type of the function to memoize.
type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

func New(f Func) *Memo {
	return &Memo{
		f:           f,
		cache:       make(map[string]result),
		cacheStatus: make(chan string, 1),
	}
}

func (m *Memo) Get(key string) (interface{}, error) {
	m.mu.Lock()
	res, ok := m.cache[key]
	m.mu.Unlock()
	if !ok {
		res.value, res.err = m.f(key)
		m.mu.Lock()
		m.cache[key] = res
		m.mu.Unlock()
		m.cacheStatus <- "Uncached -> "
	} else {
		m.cacheStatus <- "Cached -> "
	}
	return res.value, res.err
}

func httpGetBody(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func main() {
	m := New(httpGetBody)
	var wg sync.WaitGroup
	for _, url := range incomingURLs() {
		wg.Add(1)
		go func(url string) {
			start := time.Now()
			value, err := m.Get(url)
			if err != nil {
				log.Fatalln(err)
			}
			fmt.Println(<-m.cacheStatus, url, len(string(value.([]byte))), time.Since(start))
			wg.Done()
		}(url)
	}
	wg.Wait()
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
