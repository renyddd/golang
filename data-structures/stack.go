package main

import (
	"errors"
	"fmt"
)

// https://studygolang.com/articles/10481
type Stack []interface{}

func main() {
	s := Stack{2, "dfdsfsd"}

	fmt.Println(s)

	s.Push(3)
	fmt.Println(s)

	v, _ := s.Pop()
	fmt.Println(v)

	if value, err := s.Pop(); err == nil {
		fmt.Println(value)
	}
}

func (stack Stack) Len() int {
	return len(stack)
}

func (stack *Stack) Push(value interface{}) {
	*stack = append(*stack, value)
}

func (stack Stack) Top() (interface{}, error) {
	if len(stack) == 0 {
		return nil, errors.New("Out of index, len is 0")
	}
	return stack[len(stack)-1], nil
}

func (stack *Stack) Pop() (interface{}, error) {
	theStack := *stack
	if len(theStack) == 0 {
		return nil, errors.New("Out of index, len is 0")
	}
	value := theStack[len(theStack)-1]
	*stack = theStack[:len(theStack)-1]
	return value, nil
}

func (stack Stack) IsEmpty() bool {
	return len(stack) == 0
}
