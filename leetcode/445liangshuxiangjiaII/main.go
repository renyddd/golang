package main

/*
给你两个 非空 链表来代表两个非负整数。数字最高位位于链表开始位置。它们的每个节点只存储一位数字。将这两数相加会返回一个新的链表。

你可以假设除了数字 0 之外，这两个数字都不会以零开头。

进阶：
如果输入链表不能修改该如何处理？换句话说，你不能对列表中的节点进行翻转。

示例：

输入：(7 -> 2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 8 -> 0 -> 7

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/add-two-numbers-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var genStack func(*ListNode) []int
	genStack = func(h *ListNode) []int {
		s := make([]int, 0)
		for p := h; p != nil; p = p.Next {
			s = append(s, p.Val)
		}
		return s
	}

	s1, s2 := genStack(l1), genStack(l2)
	var head *ListNode

	carry := 0
	//for len(s1) > 0 && len(s2) > 0 {
	// 注意判断条件
	for len(s1) != 0 || len(s2) != 0 || carry != 0 {
		var v1, v2 int

		if len(s1) > 0 {
			v1 = s1[len(s1)-1]
			s1 = s1[:len(s1)-1]
		}
		if len(s2) > 0 {
			v2 = s2[len(s2)-1]
			s2 = s2[:len(s2)-1]
		}

		newV := (v1 + v2 + carry) % 10
		carry = (v1 + v2 + carry) / 10

		curNode := &ListNode{Val: newV, Next: head}
		head = curNode
	}

	return head
}
