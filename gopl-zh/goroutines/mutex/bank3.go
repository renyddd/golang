package main

import (
	"sync"
)

var (
	mu       sync.RWMutex
	balance2 int
)

func deposit2(amount int) {
	balance2 += amount
}

func Deposit2(amount int) {
	mu.Lock()
	defer mu.Unlock()
	deposit2(amount)
}

func Balance2() int {
	mu.Lock()
	defer mu.Unlock()
	return balance2
}

func WithDraw2(amount int) bool {
	mu.Lock()
	defer mu.Unlock()
	deposit2(-amount)
	if balance2 < 0 {
		deposit2(amount)
		return false
	}
	return true
}
