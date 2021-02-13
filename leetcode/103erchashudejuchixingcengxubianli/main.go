package main

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// main fo gdb debug
func main() {
	zigzagLevelOrder(nil)
}

// 良好的注释好简洁明了的变量明明习惯简直太重要了
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	// 期望的结果二维数组
	res := make([][]int, 0)
	// 因为当层序的深度优先遍历进行至本层最末尾的节点时，当前层的元素已全部出队并且下一层的元素也全部入队成功
	curLevelNodes := make([]*TreeNode, 0)
	curLevelNodes = append(curLevelNodes, root)

	for level := 0; len(curLevelNodes) > 0; level++ {
		nextLevelNodes := make([]*TreeNode, 0)
		curNums := make([]int, 0)

		//for _, v := range curLevelNodes {
		//	curNums = append(curNums, v.Val)
		//	if v.Left != nil {
		//		nextLevelNodes = append(nextLevelNodes, v.Left)
		//	}
		//	if v.Right != nil {
		//		nextLevelNodes = append(nextLevelNodes, v.Right)
		//	}
		//}
		for _, v := range curLevelNodes {
			if v != nil {
				curNums = append(curNums, v.Val)
				nextLevelNodes = append(nextLevelNodes, v.Left)
				//  如果在此没有排除非空元素的添加，将导致多循环判断一层叶子节点
				// [nil, nil, nil] 的列表可不是为 nil 的列表，还是需要用长度来判断
				nextLevelNodes = append(nextLevelNodes, v.Right)
			}
		}

		if level%2 == 1 {
			for i, l := 0, len(curNums); i < l/2; i++ {
				curNums[i], curNums[l-1-i] = curNums[l-1-i], curNums[i]
			}
		}

		// 一定要记得在此修改循环的判断条件
		// 当进行当叶子节点是，下一层级存储的都是 nil 将引发循环的退出
		curLevelNodes = nextLevelNodes
		if len(curNums) > 0 {
			res = append(res, curNums)
		}
	}

	return res
}
