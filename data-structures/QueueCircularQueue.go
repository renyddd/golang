package main

// https://github.com/wangzheng0822/algo/blob/master/go/09_queue/CircularQueue.go
import "fmt"

type CircularQueue struct {
	q        []interface{}
	capacity int
	head     int
	tail     int
}

func main() {
	cq := NewCircularQueue(5)

	cq.EnQueue(1)
	cq.EnQueue(2)
	cq.EnQueue(3)
	cq.EnQueue(4)
	cq.EnQueue(5)
	cq.EnQueue(6)

	for !cq.IsEmpty() {
		fmt.Println(cq.DeQueue())
	}

	cq.EnQueue(4)
	cq.EnQueue(5)
	cq.EnQueue(6)

	for !cq.IsEmpty() {
		fmt.Println(cq.DeQueue())
	}
}

func NewCircularQueue(n int) *CircularQueue {
	if n == 0 {
		return nil
	}
	return &CircularQueue{make([]interface{}, n), n, 0, 0}
}

// 栈空条件为 head == tail
func (this *CircularQueue) IsEmpty() bool {
	if this.head == this.tail {
		return true
	}
	return false
}

// 栈满条件为 (tail+1) % capacity == head
func (this *CircularQueue) IsFull() bool {
	if this.head == (this.tail+1)%this.capacity {
		return true
	}
	return false
}

func (this *CircularQueue) EnQueue(v interface{}) bool {
	if this.IsFull() {
		return false
	}
	this.q[this.tail] = v
	this.tail = (this.tail + 1) % this.capacity
	return true
}

func (this *CircularQueue) DeQueue() interface{} {
	if this.IsEmpty() {
		return nil
	}
	v := this.q[this.head]
	this.head = (this.head + 1) % this.capacity
	return v
}
