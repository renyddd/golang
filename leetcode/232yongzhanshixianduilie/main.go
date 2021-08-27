package main

/* package

请你仅使用两个栈实现先入先出队列。队列应当支持一般队列支持的所有操作（push、pop、peek、empty）：

实现 MyQueue 类：

void push(int x) 将元素 x 推到队列的末尾
int pop() 从队列的开头移除并返回元素
int peek() 返回队列开头的元素
boolean empty() 如果队列为空，返回 true ；否则，返回 false

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/implement-queue-using-stacks
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// MyQueue sIn 用于接受新元素，sOut 用来维持当前非空栈的 FIFO 队列性质
type MyQueue struct {
	sIn, sOut []int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{
		sIn:  make([]int, 0),
		sOut: make([]int, 0),
	}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.sIn = append(this.sIn, x)
}

// moveSin2Sout in 栈的搬移函数
func (q *MyQueue) moveSin2Sout() {
	for len(q.sIn) > 0 {
		elem := q.sIn[len(q.sIn)-1]
		q.sIn = q.sIn[:len(q.sIn)-1]
		q.sOut = append(q.sOut, elem)
	}
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if len(this.sOut) == 0 {
		// 搬移空 s1 至 s2，再从 s2 弹出
		this.moveSin2Sout()
	}
	out := this.sOut[len(this.sOut)-1]
	this.sOut = this.sOut[:len(this.sOut)-1]

	return out
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if len(this.sOut) == 0 {
		this.moveSin2Sout()
	}

	if l := len(this.sOut); l > 0 {
		return this.sOut[l-1]
	}
	return 0
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.sOut) == 0 && len(this.sIn) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
