package main

/* package

请你仅使用两个队列实现一个后入先出（LIFO）的栈，并支持普通栈的全部四种操作（push、top、pop 和 empty）。

实现 MyStack 类：

void push(int x) 将元素 x 压入栈顶。
int pop() 移除并返回栈顶元素。
int top() 返回栈顶元素。
boolean empty() 如果栈是空的，返回 true ；否则，返回 false 。


注意：

你只能使用队列的基本操作 —— 也就是 push to back、peek/pop from front、size 和 is empty 这些操作。
你所使用的语言也许不支持队列。 你可以使用 list （列表）或者 deque（双端队列）来模拟一个队列 , 只要是标准的队列操作即可。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/implement-stack-using-queues
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// MyStack 仅用一个队列来模拟，在插入时保持调整值的顺序
type MyStack struct {
	queue []int
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{
		queue: make([]int, 0),
	}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.queue = append(this.queue, x)
	l := len(this.queue) - 1

	// eg: x = 40; {30, 20, 10, 40}
	// 将新 append 的 40 转移至零号索引
	for ; l > 0; l-- {
		this.queue = append(this.queue, this.queue[0])
		this.queue = this.queue[1:]
	}
	// eg: {40, 30, 20, 10}
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	t := this.queue[0]
	this.queue = this.queue[1:]

	return t
}

/** Get the top element. */
func (this *MyStack) Top() int {
	return this.queue[0]
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return len(this.queue) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
