package main

/*
给定单向链表的头指针和一个要删除的节点的值，定义一个函数删除该节点。
返回删除后的链表的头节点。

示例 1:
输入: head = [4,5,1,9], val = 5
输出: [4,1,9]
解释: 给定你链表中值为 5 的第二个节点，那么在调用了你的函数之后，该链表应变为 4 -> 1 -> 9.
示例 2:

输入: head = [4,5,1,9], val = 1
输出: [4,5,9]
解释: 给定你链表中值为 1 的第三个节点，那么在调用了你的函数之后，该链表应变为 4 -> 5 -> 9.

题目保证链表中节点的值互不相同

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/shan-chu-lian-biao-de-jie-dian-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

// deleteNode 实现参考如下题解，学习拆分算法步骤的能力
// https://leetcode-cn.com/problems/shan-chu-lian-biao-de-jie-dian-lcof/solution/mian-shi-ti-18-shan-chu-lian-biao-de-jie-dian-sh-2/
func deleteNode(head *ListNode, val int) *ListNode {
	// 1. head 为空无需处理
	if head == nil {
		return nil
	}

	// 2. 特殊处理删除头节点的情况
	if head.Val == val {
		return head.Next
	}

	// 3.
	// 排除删除头节点后，删除中间尾都可统一
	// 如下循环先定位待删除节点
	prev, pD := head, head.Next
	for pD != nil {
		if pD.Val == val {
			break
		}

		prev = pD
		pD = pD.Next
	}

	// 4. 如未找到则无需删除
	if pD != nil {
		prev.Next = pD.Next
	}

	return head
}
