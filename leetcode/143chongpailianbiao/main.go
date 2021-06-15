package main

/*
给定一个单链表L：L0→L1→…→Ln-1→Ln ，
将其重新排列后变为： L0→Ln→L1→Ln-1→L2→Ln-2→…

你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

示例1:
给定链表 1->2->3->4, 重新排列为 1->4->2->3.
示例 2:
给定链表 1->2->3->4->5, 重新排列为 1->5->2->4->3.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reorder-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

// reorderList1 首先构建线性表再按序来构建链表
// 仔细审题，题目的意思为头一个尾一个重新建成链表
func reorderList1(head *ListNode) {
	if head == nil {
		return
	}

	tables := make([]*ListNode, 0)
	for p := head; p != nil; p = p.Next {
		tables = append(tables, p)
	}

	i, j := 0, len(tables)-1

	for i < j {
		tables[i].Next = tables[j]
		i++

		// 以简单的例子做验证该条件即可，此条件判断仅省略了最后一次跳出循环时的重复赋值
		// 而函数最后的赋 nil 值才保证了链表的非环状特征
		if i == j {
			break
		}

		tables[j].Next = tables[i]
		j--
	}

	tables[i].Next = nil
}

// reorderList2
// https://leetcode-cn.com/problems/reorder-list/solution/zhong-pai-lian-biao-by-leetcode-solution/
// 这里主要是想讨论一下思维方式上的差异，解法一的思维是头一个尾一个的新建链表，解法二是寻找中点后反转后半部分来合并
func reorderList2(head *ListNode) {
	if head == nil {
		return
	}

	mid := middleNode(head)
	l1 := head
	// 寻找的为左中点
	l2 := mid.Next
	// 断开原有连接避免成环
	mid.Next = nil
	l2 = reverseList(l2)
	mergeList(l1, l2)
}

func middleNode(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	// 注意当想要寻找偶数个节点的左边中点时的初始化方式
	slow, fast := head, head.Next

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	return slow
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var prev *ListNode
	p := head

	for p != nil {
		tmp := p.Next
		p.Next = prev
		prev = p
		p = tmp
	}

	return prev
}

func mergeList(l1, l2 *ListNode) {
	var l1Next, l2Next *ListNode

	for l1 != nil && l2 != nil {
		l1Next, l2Next = l1.Next, l2.Next

		l1.Next = l2
		l1 = l1Next

		l2.Next = l1
		l2 = l2Next
	}
}
