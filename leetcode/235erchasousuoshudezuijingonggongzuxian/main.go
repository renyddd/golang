package main

//ref: https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-search-tree/

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// lowestCommonAncestorNotGraceful 遍历两次树存储两次去找到 p q 的路径节点
// 在从两个切片中提取相同元素进行返回即可
func lowestCommonAncestorNotGraceful(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	var pPath, qPath []*TreeNode
	var DFS func(*TreeNode, int, *[]*TreeNode)
	DFS = func(curNode *TreeNode, pivot int, path *[]*TreeNode) {
		if curNode == nil {
			return
		}
		*path = append(*path, curNode)
		if curNode.Val > pivot {
			DFS(curNode.Left, pivot, path)
		} else if curNode.Val < pivot {
			DFS(curNode.Right, pivot, path)
		} else {
			return
		}
	}

	DFS(root, p.Val, &pPath)
	DFS(root, q.Val, &qPath)

	var res *TreeNode
	//for i := range pPath {
	//	for j := range qPath {
	//		if pPath[i].Val == qPath[j].Val {
	//			res = pPath[i]
	//			break
	//		}
	//	}
	//}
	// 寻找共同节点的遍历可进行优化，毕竟两个切片都是从根节点进行的
	for i := 0; i < len(pPath) && i < len(qPath) && pPath[i] == qPath[i]; i++ {
		res = pPath[i]
	}

	return res
}

/*
official solve: 只需要一次遍历但同时进行对 p 与 q 的比较，若当前节点
同时小于或大于 p 于 q，则按照 bst 的性质进行遍历即可，若当前节点
不同时符合上述条件，则说明到了两节点的公公祖先处即分岔点
 */
func lowestCommonAncestorOfficial(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	ancestor := root
	for ancestor != nil {
		if p.Val < ancestor.Val && q.Val < ancestor.Val {
			ancestor = ancestor.Left
		} else if p.Val > ancestor.Val && q.Val > ancestor.Val {
			ancestor = ancestor.Right
		} else {
			return ancestor
		}
	}

	return ancestor
}

// 当然如上的过程也能用递归来改写出来
func lowestCommonAncestorRecursion(root, p, q *TreeNode) *TreeNode {
	if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestorRecursion(root.Left, p, q)
	} else if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestorRecursion(root.Right, p, q)
	} else {
		return root
	}
}