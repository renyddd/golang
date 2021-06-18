package main

/*
给你链表的头节点 head 和一个整数 k 。

交换 链表正数第 k 个节点和倒数第 k 个节点的值后，返回链表的头节点（链表 从 1 开始索引）。
示例 1：
输入：head = [1,2,3,4,5], k = 2
输出：[1,4,3,2,5]
示例 2：
输入：head = [7,9,6,6,7,8,3,0,9,5], k = 5
输出：[7,9,6,6,8,7,3,0,9,5]
示例 3：
输入：head = [1], k = 1
输出：[1]
示例 4：
输入：head = [1,2], k = 1
输出：[2,1]
示例 5：
输入：head = [1,2,3], k = 2
输出：[1,2,3]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/swapping-nodes-in-a-linked-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapNodes(head *ListNode, k int) *ListNode {
	var l int
	var left, right *ListNode

	for p := head; p != nil; p = p.Next {
		l++
		if l == k {
			left = p
		}
	}

	right = head
	for l != k {
		right = right.Next
		l--
	}

	left.Val, right.Val = right.Val, left.Val
	return head
}

// swapNodes2 快慢指针拉开需要的差距，当快指针到达时满指针即为所求节点
func swapNodes2(head *ListNode, k int) *ListNode {
	fast, slow := head, head
	for cnt := 1; cnt != k; cnt++ {
		fast = fast.Next
	}

	var left, right *ListNode
	left = fast

	// fast 此时只是索引正确，但需要再多一个以增加间隔距离
	for ; fast.Next != nil; fast = fast.Next {
		slow = slow.Next
	}
	right = slow

	left.Val, right.Val = right.Val, left.Val
	return head
}
