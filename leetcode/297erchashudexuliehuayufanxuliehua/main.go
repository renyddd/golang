package main

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (c *Codec) serialize(root *TreeNode) string {
	return do_serialize(root)
}

func do_serialize(root *TreeNode) string {
	if root == nil {
		return "#"
	}
	return fmt.Sprintf("%d,%s,%s", root.Val, do_serialize(root.Left), do_serialize(root.Right))
}

// Deserializes your encoded data to tree.
func (c *Codec) deserialize(data string) *TreeNode {
	slice := strings.Split(data, ",")
	return do_deserialize(&slice)
}

func do_deserialize(s *[]string) *TreeNode {
	preRoot := (*s)[0]
	// 注意字符串减一操作必须放在这，否则将会在第一次遇见 # 后一直返回 nil
	*s = (*s)[1:]
	if preRoot == "#" {
		return nil
	}
	rootVal, _ := strconv.Atoi(preRoot)
	root := &TreeNode{rootVal, nil, nil}

	root.Left = do_deserialize(s)
	root.Right = do_deserialize(s)

	return root
}


/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */