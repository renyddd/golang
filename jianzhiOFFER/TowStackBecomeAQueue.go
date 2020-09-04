package main

import "fmt"

type Stack struct {
	Items []int
}

type Queue struct {
	in  Stack
	out Stack
}

func main() {
	// 使用 new 来为结构体变量分配内存，并零值化
	// q := &Queue{} 的底层还是在调用 new()
	q := new(Queue)
	q.AppenTail(5)
	q.AppenTail(2)
	q.AppenTail(1)

	fmt.Println(q.DeleteHead())
	fmt.Println(q.DeleteHead())
	fmt.Println(q.DeleteHead())
	q.AppenTail(1)
	fmt.Println(q.DeleteHead())
	fmt.Println(q.DeleteHead())
}

func (queue *Queue) AppenTail(val int) {
	queue.in.Push(val)
}

func (queue *Queue) DeleteHead() int {
	// 出队列时分为两种情况：
	// 1. queue.out 栈不为空时其栈顶元素便为最先入队
	// 2. queue.out 栈为空时需要将 in 栈元素全部彈栈并且压入 out
	//    以保证 out 栈顶元素为此刻队列最头

	var val int = -1

	if !queue.out.Isempty() {
		val = queue.out.Pop()
	} else {
		// 情况二 out 为空：in 全部彈栈并入 out 栈
		for !queue.in.Isempty() {
			queue.out.Push(queue.in.Pop())
		}
		val = queue.out.Pop()
	}
	return val
}

func (queue *Queue) Isempty() bool {
	if queue.out.Isempty() && queue.in.Isempty() {
		return true
	}
	return false
}

func (stack Stack) Isempty() bool {
	if len(stack.Items) == 0 {
		return true
	}
	return false
}

func (stack *Stack) Push(val int) {
	stack.Items = append(stack.Items, val)
}

func (stack *Stack) Pop() int {
	if stack.Isempty() {
		return -1
	}
	val := stack.Items[len(stack.Items)-1]
	// 冒号后的数字代表”第多少个元素“，而非索引
	stack.Items = stack.Items[:len(stack.Items)-1]

	return val
}
