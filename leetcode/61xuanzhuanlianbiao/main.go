package main

/*
给你一个链表的头节点 head ，旋转链表，将链表每个节点向右移动 k 个位置。

示例 1：

输入：head = [1,2,3,4,5], k = 2
输出：[4,5,1,2,3]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/rotate-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

// rotateRight 官方解先使链表成环，再断开
// 这样做能避免上很么样的问题呢？能带来什么样的好处？
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}

	tail, l := head, 1
	for ; tail.Next != nil; tail, l = tail.Next, l+1 {
	}

	if k%l == 0 {
		return head
	}

	// 形成环形
	tail.Next = head

	for d := l - k%l - 1; d > 0; head, d = head.Next, d-1 {
	}

	newHead := head.Next
	head.Next = nil

	return newHead
}

// rotateRight2  快慢指针来解决
// https://leetcode-cn.com/problems/rotate-list/solution/kuai-man-zhi-zhen-ru-he-fen-bu-zou-jie-j-ns7u/
func rotateRight2(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}

	l := 0
	for p := head; p != nil; p = p.Next {
		l++
	}

	k = k % l
	if k == 0 {
		return head
	}

	fast, slow := head, head
	for k > 0 {
		fast = fast.Next
		k--
	}

	for fast.Next != nil {
		fast, slow = fast.Next, slow.Next
	}

	newH := slow.Next
	fast.Next = head
	slow.Next = nil

	return newH
}

// 如下自想代码保留，注意参照上方与官方代码的思维不同之处

// rotateRight 分析题目需求，生的解法一
// 首先获取到链表总长度，向右移动 k 个位置相当于是，从倒数第 k 个节点处断开链接
// 并且拼接到 head 处
// func rotateRight(head *ListNode, k int) *ListNode {
// 	if k == 0 || head == nil {
// 		return head
// 	}

// 	var tail *ListNode
// 	l := 0
// 	for p := head; p != nil; p = p.Next {
// 		l++
// 		if p.Next == nil {
// 			tail = p
// 		}
// 	}

// 	if l <= 1 {
// 		return head
// 	}

// 	// 还需处理 k 超出链表长度的情况
// 	t := k % l

// 	// pD 即为倒数第 k 个节点
// 	pD, d := head, l-t-1
// 	for d > 0 {
// 		pD, d = pD.Next, d-1
// 	}

// 	pTmp := pD.Next
// 	pD.Next = nil
// 	tail.Next = head
// 	head = pTmp

// 	return head
// }
