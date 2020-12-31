package main

// https://github.com/floyernick/Data-Structures-and-Algorithms/blob/master/LinkedList/LinkedList.go

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (list *LinkedList) InsertFirst(i int) {
	p := &Node{data: i}
	if list.head != nil {
		p.next = list.head
	}
	list.head = p
}

func (list *LinkedList) InsertLast(i int) {
	p := &Node{data: i}
	if list.head == nil {
		list.head = p
		return
	}
	current := list.head
	for ; current.next != nil; current = current.next {

	}
	current.next = p
}

func (list *LinkedList) GetItems() []int {
	var items []int
	current := list.head
	for current != nil {
		items = append(items, current.data)
		current = current.next
	}
	return items
}

func (list *LinkedList) RemoveByValue(i int) bool {
	if list.head == nil {
		return false
	}
	if list.head.data == i {
		list.head = list.head.next
		return true
	}
	current := list.head
	for current.next != nil {
		if current.next.data == i {
			current.next = current.next.next
			return true
		}
		current = current.next
	}
	return false
}

func (list *LinkedList) RemoveByIndex(i int) bool {
	if list.head == nil {
		return false
	}
	if i < 0 {
		return false
	}
	if i == 0 {
		list.head = list.head.next
		return true
	}
	current := list.head
	for u := 1; u < i; u++ {
		if current.next.next == nil {
			return false
		}
		current = current.next
	}
	current.next = current.next.next
	return true
}

func (list *LinkedList) SearchValue(i int) bool {
	if list.head == nil {
		return false
	}
	for current := list.head; current != nil; current = current.next {
		if current.data == i {
			return true
		}
	}
	return false
}

func (list *LinkedList) GetFirst() (int, bool) {
	if list.head == nil {
		return 0, false
	}
	return list.head.data, true
}

func (list *LinkedList) GetLast() (int, bool) {
	if list.head == nil {
		return 0, false
	}
	current := list.head
	for ; current.next != nil; current = current.next {
	}
	return current.data, true
}

func (list *LinkedList) GetSize() int {
	count := 0
	for current := list.head; current.next != nil; current = current.next {
		count++
	}
	return count
}

func main() {
	//	l1 := new(LinkedList)
	var l1 LinkedList
	l1.InsertFirst(3)
	l1.InsertFirst(2)
	l1.InsertFirst(1)
	l1.InsertLast(4)
	l1.InsertLast(5)
	fmt.Println(l1.GetItems())

	if l1.RemoveByValue(21) {
		fmt.Println(l1.GetItems())
	}

	if l1.RemoveByIndex(3) {
		fmt.Println(l1.GetItems())
	}

	fmt.Println(l1.SearchValue(3))
	fmt.Println(l1.GetFirst())
	fmt.Println(l1.GetLast())
	fmt.Println(l1.GetSize())
}
