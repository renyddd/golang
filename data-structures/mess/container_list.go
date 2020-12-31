package main

import (
	"container/list"
	"fmt"
)

func main() {
	l := list.New()
	l.PushBack(4)
	l.PushFront(1)
	// l.InsertBefore(3, e4)
	// l.InsertAfter(2, e1)

	l.PushBack(5)
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
