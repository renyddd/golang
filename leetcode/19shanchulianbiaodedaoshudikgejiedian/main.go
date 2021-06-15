package main

/*
给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。

进阶：你能尝试使用一趟扫描实现吗？

示例 1：
输入：head = [1,2,3,4,5], n = 2
输出：[1,2,3,5]

示例 2：
输入：head = [1], n = 1
输出：[]

示例 3：
输入：head = [1,2], n = 1
输出：[1]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/solution/shan-chu-lian-biao-de-dao-shu-di-nge-jie-dian-b-61/
可参照题解来在草稿纸上完成
*/

// removeNthFromEnd 首先设置表头节点，然后遍历至待删节点的前驱节点
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	l := 0
	for p := head; p != nil; p = p.Next {
		l++
	}

	// 首先明确设置表头节点（哑节点）的意义：
	// 去除对链表头节点删除的特殊性
	dummy := &ListNode{-1, head}
	p := dummy

	// 注意循环参数的设置，因为需要找的使待删节点的前驱节点，故从哑节点开始遍历
	for i := 0; i < l-n; i++ {
		p = p.Next
	}

	p.Next = p.Next.Next

	// 因以取出头节点删除时的特殊性，故只需固定返回哑节点的 Next 域即可
	return dummy.Next
}

// removeNthFromEnd2 用栈的思想来模拟，其意义就是空间换时间
// 能通过单次遍历的时间来迅速找到待删除节点的前驱
func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	if head == nil {
		return head
	}

	// 将 dummy 算进链表长度里，也是为了取出头节点的特殊性
	dummy := &ListNode{-1, head}
	nodes := make([]*ListNode, 0)

	for p := dummy; p != nil; p = p.Next {
		nodes = append(nodes, p)
	}

	// 注意：同样找的是前驱节点
	pD := nodes[len(nodes)-n-1]
	pD.Next = pD.Next.Next

	return dummy.Next
}

// removeNthFromEnd3 采用快慢指针来完成删除
// 当然也需要 dummy 节点的帮助
func removeNthFromEnd3(head *ListNode, n int) *ListNode {
	// Tip: fast 从 head 开始先跑 n 个节点
	fast := head
	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	// Tip： slow 从 dummy 开始，以便迅速找出待删除节点的前驱
	dummy := &ListNode{-1, head}
	slow := dummy

	// Tip：当 fast 直接为 nil 时则代表所删除的为 head 节点
	for fast != nil {
		slow, fast = slow.Next, fast.Next
	}

	slow.Next = slow.Next.Next

	return dummy.Next
}

// https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/comments/229867
func removeNthFromEnd4(head *ListNode, n int) *ListNode {
	num := 0

	var iter func(*ListNode) *ListNode
	iter = func(ln *ListNode) *ListNode {
		if ln == nil {
			return ln
		}
		ln.Next = iter(ln.Next)
		num++

		if num == n {
			return ln.Next
		}
		return ln
	}

	return iter(head)
}
