package main

import "fmt"

type Element struct {
	next, prev *Element
	list       *List
	Value      interface{}
}

type List struct {
	root Element
	len  int
}

func (e *Element) Prev() *Element {
	if p := e.prev; e.list != nil && p != &e.list.root {
		return p
	}
	return nil
}

func (e *Element) Next() *Element {
	if p := e.next; e.list != nil && p != &e.list.root {
		return p
	}
	return nil
}

func (l *List) Init() *List {
	l.root.next = &l.root
	l.root.prev = &l.root
	l.len = 0
	return l
}

func (l *List) lazyInit() {
	if l.root.next == nil {
		l.Init()
	}
}

func New() *List {
	return new(List).Init()
}

func (l *List) Len() int { return l.len }

func (l *List) insert(e, at *Element) *Element {
	n := at.next
	at.next = e
	e.prev = at
	e.next = n
	n.prev = e
	e.list = l
	l.len++
	return e
}

// a convenience wrapper fo inser
func (l *List) insertValue(v interface{}, at *Element) *Element {
	return l.insert(&Element{Value: v}, at)
}

func (l *List) PushBack(v interface{}) *Element {
	l.lazyInit()
	return l.insertValue(v, l.root.prev)
}

func main() {
	l := new(List)
	l.Init()
	l.PushBack(5)
	fmt.Println(l)
}
