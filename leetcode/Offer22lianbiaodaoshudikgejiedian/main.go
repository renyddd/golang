package main

/*
输入一个链表，输出该链表中倒数第k个节点。为了符合大多数人的习惯，本题从1开始计数，即链表的尾节点是倒数第1个节点。

例如，一个链表有 6 个节点，从头节点开始，它们的值依次是 1、2、3、4、5、6。这个链表的倒数第 3 个节点是值为 4 的节点。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/lian-biao-zhong-dao-shu-di-kge-jie-dian-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

// getKthFromEnd_0 遍历两次
func getKthFromEnd_0(head *ListNode, k int) *ListNode {
	// 注意这里遍历时的取值
	l := 0
	for p := head; p != nil; p = p.Next {
		l++
	}

	// 如果上面选 l := 1
	// 那么这里会面临单一节点链表返回为空的情况
	if k > l {
		return nil
	}

	// 就是需要往后走 l-k+1 步
	t := 1
	for ; t < l-k+1; head = head.Next {
		t++
	}

	return head
}

// getKthFromEnd_1 遍历两次使用快慢指针拉开差距
func getKthFromEnd_1(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return nil
	}
	// 让 fast 比 slow 快 k 步长即可
	fast := head
	for i := 0; i < k; i++ {
		if fast.Next != nil {
			fast = fast.Next
		} else {
			return nil
		}
	}

	slow := head
	for ; fast != nil; fast = fast.Next {
		slow = slow.Next
	}

	return slow
}
