package main

/*
定义一个函数，输入一个链表的头节点，反转该链表并输出反转后链表的头节点。

示例:

输入: 1->2->3->4->5->NULL
输出: 5->4->3->2->1->NULL

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/fan-zhuan-lian-biao-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/* 题解心得
先完成局部小模块的验证，比如假设当前已反转至链表的中部，接着以此种情况来构造算法步骤
然后再将该算法步骤推广至链表尾部、首部
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

// reverseList
func reverseList(head *ListNode) *ListNode {
	var prev, p *ListNode
	p = head

	for p != nil {
		tmp := p.Next
		p.Next = prev
		prev = p
		p = tmp
	}

	return prev
}
