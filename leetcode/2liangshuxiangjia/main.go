package main

/*
给你两个非空 的链表，表示两个非负的整数。它们每位数字都是按照逆序的方式存储的，并且每个节点只能存储一位数字。

请你将两个数相加，并以相同形式返回一个表示和的链表。

你可以假设除了数字 0 之外，这两个数都不会以 0开头。

输入：l1 = [2,4,3], l2 = [5,6,4]
输出：[7,0,8]
解释：342 + 465 = 807.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/add-two-numbers
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

// addTwoNumbers
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var tail, head *ListNode
	carry := 0

	for l1 != nil || l2 != nil {
		var n1, n2 int
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		sum := n1 + n2 + carry
		sum, carry = sum%10, sum/10

		if head == nil {
			head = &ListNode{Val: sum}
			tail = head
		} else {
			tail.Next = &ListNode{Val: sum}
			tail = tail.Next
		}
	}

	// 最后的进位值要生成新的节点
	if carry > 0 {
		tail.Next = &ListNode{Val: carry}
		tail = tail.Next
	}

	return head
}

// addTwoNumbers2
// ref: https://leetcode-cn.com/problems/add-two-numbers/solution/di-gui-si-lu-jian-dan-dai-ma-duan-by-dnanki/
func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	var dfs func(l1, l2 *ListNode, carry int) *ListNode
	dfs = func(l1, l2 *ListNode, carry int) *ListNode {
		if l1 == nil && l2 == nil && carry == 0 {
			return nil
		}
		var sum, n1, n2 int
		var nl1, nl2 *ListNode
		if l1 != nil {
			n1 = l1.Val
			nl1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			nl2 = l2.Next
		}
		sum = n1 + n2 + carry
		carry = sum / 10
		sum = sum % 10
		retNode := &ListNode{}
		retNode.Val = sum
		retNode.Next = dfs(nl1, nl2, carry)
		return retNode
	}

	return dfs(l1, l2, 0)
}
