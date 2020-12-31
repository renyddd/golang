package main

import "fmt"

// https://github.com/wangzheng0822/algo/blob/master/go/09_queue/QueueBasedOnArray.go

type ArrayQueue struct {
	q        []interface{}
	capacity int
	head     int
	tail     int
}

func main() {
	q := NewArrayQueue(5)
	q.EnQueue(5)
	q.EnQueue(7)
	q.EnQueue(9)

	fmt.Println(q.DeQueue())
	fmt.Println(q.DeQueue())
}

func NewArrayQueue(n int) *ArrayQueue {
	return &ArrayQueue{make([]interface{}, n), n, 0, 0}
}

func (this *ArrayQueue) EnQueue(v interface{}) bool {
	if this.tail == this.capacity {
		return false
	}
	this.q[this.tail] = v
	this.tail++
	return true
}

func (this *ArrayQueue) DeQueue() interface{} {
	if this.head == this.tail {
		return nil
	}
	v := this.q[this.head]
	this.head++
	return v
}
