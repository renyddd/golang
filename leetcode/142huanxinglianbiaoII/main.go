package main

/*
给定一个链表，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。
为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。注意，pos 仅仅是用于标识环的情况，并不会作为参数传递到函数中。
说明：不允许修改给定的链表。

进阶：
你是否可以使用 O(1) 空间解决此题？

示例 1：
输入：head = [3,2,0,-4], pos = 1
输出：返回索引为 1 的链表节点
解释：链表中有一个环，其尾部连接到第二个节点。

示例 2：
输入：head = [1,2], pos = 0
输出：返回索引为 0 的链表节点
解释：链表中有一个环，其尾部连接到第一个节点。

示例 3：
输入：head = [1], pos = -1
输出：返回 null
解释：链表中没有环。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/linked-list-cycle-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

// detectCycle 首先采用 map 的方式肯定是能解决掉的
func detectCycle(head *ListNode) *ListNode {
	m := make(map[*ListNode]struct{}, 0)

	for p := head; p != nil; p = p.Next {
		if _, ok := m[p]; ok {
			return p
		}
		m[p] = struct{}{}
	}

	return nil
}

// detectCycle2 尝试速度比 1:2 的快慢指针是否可行，
// 之后再是这探讨下其中的数学原因
func detectCycle2(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	slow, fast := head, head

	for fast != nil || fast.Next != nil {
		if fast == slow {
			return fast
		}

		slow, fast = slow.Next, fast.Next.Next
	}

	return nil
}
