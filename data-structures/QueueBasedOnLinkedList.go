package main

// https://github.com/wangzheng0822/algo/blob/master/go/09_queue/QueueBasedOnLinkedList.go
import "fmt"

type ListNode struct {
	val  interface{}
	next *ListNode
}

type LinkedListQueue struct {
	head   *ListNode
	tail   *ListNode
	length int
}

func main() {
	q := NewLinkedListQueue()

	q.EnQueue(1)
	q.EnQueue(3)
	q.EnQueue(13)
	q.EnQueue(9)
	q.EnQueue(0)

	fmt.Println(q.DeQueue())
	fmt.Println(q.DeQueue())
	fmt.Println(q.DeQueue())
	fmt.Println(q.DeQueue())
}

func NewLinkedListQueue() *LinkedListQueue {
	return &LinkedListQueue{nil, nil, 0}
}

func (this *LinkedListQueue) EnQueue(v interface{}) {
	node := &ListNode{v, nil}
	if nil == this.tail {
		this.tail = node
		this.head = node
	} else {
		this.tail.next = node
		this.tail = node
	}
	this.length++
}

func (this *LinkedListQueue) DeQueue() interface{} {
	if this.head == nil {
		return nil
	}
	v := this.head.val
	this.head = this.head.next
	this.length--
	return v
}
