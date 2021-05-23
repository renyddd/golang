package main

/*
输入两个链表，找出它们的第一个公共节点。


输入：intersectVal = 8, listA = [4,1,8,4,5], listB = [5,0,1,8,4,5], skipA = 2, skipB = 3
输出：Reference of the node with value = 8
输入解释：相交节点的值为 8 （注意，如果两个列表相交则不能为 0）。从各自的表头开始算起，链表 A 为 [4,1,8,4,5]，链表 B 为 [5,0,1,8,4,5]。在 A 中，相交节点前有 2 个节点；在 B 中，相交节点前有 3 个节点。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/liang-ge-lian-biao-de-di-yi-ge-gong-gong-jie-dian-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode-cn.com/problems/liang-ge-lian-biao-de-di-yi-ge-gong-gong-jie-dian-lcof/solution/shuang-zhi-zhen-fa-lang-man-xiang-yu-by-ml-zimingm/
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	n1, n2 := headA, headB

	for n1 != n2 {
		// 注意需要让 n1 去判别是否进行到了 nil 值，以最简单的两个不相交的但节点链表为例
		// 若综艺 n1.Next 去判断，则会处在无限的循环当中
		// 同时也是有机会让两条不相交的链表有机会，共同走向自己的 nil，从而包含这种情况
		if n1 == nil {
			n1 = headB
		} else {
			n1 = n1.Next
		}

		if n2 == nil {
			n2 = headA
		} else {
			n2 = n2.Next
		}
	}

	return n1
}

// getIntersectionNode2 .
func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	ma := make(map[*ListNode]struct{}, 0)

	for p := headA; p != nil; p = p.Next {
		ma[p] = struct{}{}
	}

	for p := headB; p != nil; p = p.Next {
		if _, ok := ma[p]; ok {
			return p
		}
	}

	return nil
}

func getIntersectionNode3(headA, headB *ListNode) *ListNode {
	var getLen func(*ListNode) int
	getLen = func(h *ListNode) (l int) {
		for h != nil {
			h, l = h.Next, l+1
		}
		return
	}

	la, lb := getLen(headA), getLen(headB)

	// 使较两链表剩余同样的长度，注意这种同步的方法
	for la != lb {
		if la > lb {
			headA = headA.Next
			la--
		} else {
			headB = headB.Next
			lb--
		}
	}

	// 此时剩余长度已相同，故碰见不相交时也会共同达到 nil 节点
	for headA != headB {
		headA, headB = headA.Next, headB.Next
	}

	return headA
}
