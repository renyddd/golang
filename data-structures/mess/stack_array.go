package main

import "fmt"

const arraySize = 10

type Stack struct {
	top  int
	data [arraySize]int
}

func (s *Stack) Push(i int) bool {
	if s.top == len(s.data) {
		return false
	}
	s.data[s.top] = i
	s.top += 1
	return true
}

func (s *Stack) Pop() (int, bool) {
	if s.top == 0 {
		return 0, false
	}
	i := s.data[s.top-1]
	s.top -= 1
	return i, true
}

func (s *Stack) Peek() int {
	return s.data[s.top-1]
}

func (s *Stack) IsEmpty() bool {
	return s.top == 0
}

func main() {
	var s Stack
	s.Push(3)
	s.Push(2)
	s.Push(5)
	for !s.IsEmpty() {
		fmt.Println(s.Peek())
		s.Pop()
	}
}
