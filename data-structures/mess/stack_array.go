package main

import "fmt"

const arraySize = 10

type Stack3 struct {
	top  int
	data [arraySize]int
}

func (s *Stack3) Push(i int) bool {
	if s.top == len(s.data) {
		return false
	}
	s.data[s.top] = i
	s.top += 1
	return true
}

func (s *Stack3) Pop() (int, bool) {
	if s.top == 0 {
		return 0, false
	}
	i := s.data[s.top-1]
	s.top -= 1
	return i, true
}

func (s *Stack3) Peek() int {
	return s.data[s.top-1]
}

func (s *Stack3) IsEmpty() bool {
	return s.top == 0
}

func main() {
	var s Stack3
	s.Push(3)
	s.Push(2)
	s.Push(5)
	for !s.IsEmpty() {
		fmt.Println(s.Peek())
		s.Pop()
	}
}
