package main

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

/*
	ref: https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree/
	此题进行查找的是一个颗普通的二叉树，并无法利用 235 题中的 bst 的性质

	官方解答中，定义函数 fx 为，(flson && frson) || (x == p || x == q) && (flson || frson)
	- 左子树且右子树
	- q 或 p 为 x 节点
	- 左子树或右子树
*/

func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	var res *TreeNode

	var fDFS func(*TreeNode, *TreeNode, *TreeNode) bool
	/*
		**注意** 内部函数的内部变量名！
		这个坑你想踩几次？
	*/
	fDFS = func(curNode, p, q *TreeNode) bool {
		if curNode == nil {
			return false
		}

		lson := fDFS(curNode.Left, p, q)
		rson := fDFS(curNode.Right, p, q)

		if (lson && rson) || ((curNode.Val == p.Val || curNode.Val == q.Val) && (lson || rson)) {
			res = curNode
		}

		return lson || rson || (curNode.Val == p.Val || curNode.Val == q.Val)
	}

	fDFS(root, p, q)

	return res
}

/*
	官方的另一种解法
	利用该二叉树的值不重复的特性，可用一个 map[int]{*TreeNode, bool} 来存储每一个 val 值的父节点信息
	再先利用 p 来遍历整个 map 路径并表示已访问过的节点，再当 q 来遍历整个 map 是一旦遇见访问过的节点便是最小公共祖先
*/

func LowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	type Pair struct {
		Node    *TreeNode
		Visited bool
	}
	// 存储 val 的父节点与是否被访问过的信息
	pathMap := make(map[int]*Pair, 0)
	var DFS func(curNode *TreeNode)
	DFS = func(curNode *TreeNode) {
		if curNode == nil {
			return
		}
		if curNode.Left != nil {
			pathMap[curNode.Left.Val] = &Pair{curNode, false}
			DFS(curNode.Left)
		}
		if curNode.Right != nil {
			pathMap[curNode.Right.Val] = &Pair{curNode, false}
			DFS(curNode.Right)
		}
	}

	// 记得初始化 root 节点的 *TreeNode 非非空，否则后续将导致解引用空指针
	pathMap[root.Val] = &Pair{nil, false}

	DFS(root)

	// Visited 信息在 p 遍历时更新
	for p != nil {
		pathMap[p.Val].Visited = true
		p = pathMap[p.Val].Node
	}

	for q != nil {
		if pathMap[q.Val].Visited {
			return q
		}
		q = pathMap[q.Val].Node
	}

	//for v := p; v != nil; v = pathMap[v.Val].Node {
	//	if _, ok := pathMap[v.Val]; ok {
	//		pathMap[v.Val].Visited = true
	//	}
	//}
	//for v := q; v != nil; v = pathMap[v.Val].Node {
	//	if pathMap[v.Val].Visited {
	//		return v
	//	}
	//}
	return nil
}

func main() {
	left := &TreeNode{2, nil, nil}
	right := &TreeNode{3, nil, nil}
	root := &TreeNode{1, left, right}
	LowestCommonAncestor(root, left, right)
}
