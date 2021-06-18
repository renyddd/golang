package main

/*
给你一个单链表的引用结点 head。链表中每个结点的值不是 0 就是 1。已知此链表是一个整数数字的二进制表示形式。

请你返回该链表所表示数字的 十进制值 。
输入：head = [1,0,1]
输出：5
解释：二进制数 (101) 转化为十进制数 (5)

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/convert-binary-number-in-a-linked-list-to-integer
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func getDecimalValue(head *ListNode) int {
	if head == nil {
		return 0
	}

	var res int

	for p := head; p != nil; p = p.Next {
		res *= 2

		// res <<= 1

		res += p.Val
	}

	return res
}
