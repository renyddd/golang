package main

import (
	"container/list"
	"log"
)

func main() {
	l := list.New()
 	l.PushBack(666)
	l.PushBack(777)
	l.PushFront(555)

	for v := l.Front(); v != nil; v = v.Next() {
		log.Printf("%v\n", v.Value)
	}
}