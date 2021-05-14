package main

import "testing"

func TestBFS(t *testing.T) {
	root := &TreeNode{Element: 6}
	var i ElementType = 0
	for ; i < 13; i++ {
		root = InsertIteration(i, root)
	}

	/*
		PrintTree(root) - all right
			0       1       2       3       4       5       6       7       8       9       10      11      12      --- PASS: TestBFS (0.00s)
			PASS
	*/

	// 有次递增的给出数据树结构完全会退还成链表
	res := BFS(root)
	t.Log(res)
}

func TestBFS2(t *testing.T) {
	type testData struct {
		Root *TreeNode
		Nums []ElementType
	}
	TS := []testData{
		{&TreeNode{Element: 4}, []ElementType{1, 2, 3, 4, 5, 6, 7}},
		{&TreeNode{Element: 4}, []ElementType{7, 6, 5, 4, 3, 2, 1}},
		{&TreeNode{Element: 4}, []ElementType{3, 6, 7, 4, 2, 1, 5}},
	}

	for _, v := range TS {
		for _, e := range v.Nums {
			v.Root = InsertIteration(e, v.Root)
		}
	}

	for _, v := range TS {
		t.Log(BFS(v.Root))
	}
}
