package main

// some ref:
// - https://refactoringguru.cn/design-patterns/singleton/go/example
// - http://marcio.io/2015/07/singleton-pattern-in-go/

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

type single struct {
}

var singleInstance *single

func getInstance() *single {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()

		if singleInstance == nil {
			fmt.Println("creating single instance.")
			singleInstance = &single{}
		} else {
			fmt.Println("already created")
		}
	} else {
		fmt.Println("already exist")
	}

	return singleInstance
}
