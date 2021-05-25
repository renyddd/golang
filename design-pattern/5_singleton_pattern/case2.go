package main

import (
	"fmt"
	"sync"
)

var once sync.Once

type single2 struct {
}

var singleInstance2 *single2

func getInstance2() *single2 {
	if singleInstance2 == nil {
		once.Do(
			func() {
				fmt.Println("creating single2")
				singleInstance2 = &single2{}
			})
	} else {
		fmt.Println("already exist")
	}

	return singleInstance2
}
