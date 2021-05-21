package main

/*
输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回）。

示例 1：

输入：head = [1,3,2]
输出：[2,3,1]
输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回）。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/cong-wei-dao-tou-da-yin-lian-biao-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

数组反转：https://stackoverflow.com/questions/19239449/how-do-i-reverse-an-array-in-go
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

// reversePrint 反转数组
func reversePrint(head *ListNode) []int {
	res := make([]int, 0)

	for p := head; p != nil; p = p.Next {
		res = append(res, p.Val)
	}

	doSliceReverse(res)

	return res
}

// reversePrint2 首先获取链表长度后反向填充
func reversePrint2(head *ListNode) []int {
	l := 0
	for p := head; p != nil; p = p.Next {
		l++
	}

	res := make([]int, l)
	l--

	for p := head; p != nil; p = p.Next {
		res[l] = p.Val
		l--
	}

	return res
}

func reversePrint3(head *ListNode) []int {
	var iter func(*ListNode) []int
	iter = func(ln *ListNode) []int {
		if ln == nil {
			return make([]int, 0)
		}
		return append(iter(ln.Next), ln.Val)
	}

	return iter(head)
}
