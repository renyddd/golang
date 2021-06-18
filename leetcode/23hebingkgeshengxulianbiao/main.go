package main

/*
给你一个链表数组，每个链表都已经按升序排列。
请你将所有链表合并到一个升序链表中，返回合并后的链表。

示例 1：
输入：lists = [[1,4,5],[1,3,4],[2,6]]
输出：[1,1,2,3,4,4,5,6]
解释：链表数组如下：
[
  1->4->5,
  1->3->4,
  2->6
]
将它们合并到一个有序链表中得到。
1->1->2->3->4->4->5->6

示例 2：
输入：lists = []
输出：[]

示例 3：
输入：lists = [[]]
输出：[]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/merge-k-sorted-lists
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

题解参照：https://leetcode-cn.com/problems/merge-k-sorted-lists/solution/he-bing-kge-pai-xu-lian-biao-by-leetcode-solutio-2/
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

// mergeKLists1 采用顺序合并，第 i 次循环将第 i 个链表合并到答案中
func mergeKLists1(lists []*ListNode) *ListNode {
	var mergeTowLists func(*ListNode, *ListNode) *ListNode
	mergeTowLists = func(h1 *ListNode, h2 *ListNode) *ListNode {
		if h1 == nil {
			return h2
		} else if h2 == nil {
			return h1
		}

		tail := &ListNode{Val: -1}
		dummyHead := tail

		for h1 != nil && h2 != nil {
			var p *ListNode
			if h1.Val < h2.Val {
				p = h1
				h1 = h1.Next
			} else {
				p = h2
				h2 = h2.Next
			}
			tail.Next = p
			tail = tail.Next
		}

		if h1 == nil {
			tail.Next = h2
		} else {
			tail.Next = h1
		}

		return dummyHead.Next
	}

	var res *ListNode

	for _, l := range lists {
		res = mergeTowLists(res, l)
	}

	return res
}

func merge2List(h1, h2 *ListNode) *ListNode {
	if h1 == nil {
		return h2
	} else if h2 == nil {
		return h1
	}

	tail := &ListNode{Val: -1}
	dummyHead := tail

	for h1 != nil && h2 != nil {
		if h1.Val < h2.Val {
			tail.Next = h1
			h1 = h1.Next
		} else {
			tail.Next = h2
			h2 = h2.Next
		}
		tail = tail.Next
	}

	if h1 == nil {
		tail.Next = h2
	} else {
		tail.Next = h1
	}

	return dummyHead.Next
}

// mergeKLists2 采用分治的思想来处理，这里采用二分的方式来分割
func mergeKLists2(lists []*ListNode) *ListNode {
	var merge func(int, int) *ListNode
	merge = func(l int, r int) *ListNode {
		if l == r {
			return lists[l]
		}
		if l > r {
			return nil
		}
		mid := (l + r) >> 1
		return merge2List(merge(l, mid), merge(mid+1, r))
	}

	return merge(0, len(lists))
}
