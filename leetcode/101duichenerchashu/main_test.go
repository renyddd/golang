package main

import (
	"../treenode"
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestIsSymmetric(t *testing.T) {
	root := treenode.BuildA163TmpTree()
	t.Log(isSymmetric(root))

	t.Log(isSymmetric(treenode.BuildSymmetric811Tree()))
}

func TestBC(t *testing.T) {
	num := "154800 574752 33480 174816 78432 125184 120960 12096 148128 25440 692928 288000"

	n := strings.Split(num, " ")
	var res int
	for _, v := range n {
		add, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		res += add
		fmt.Println(res)
	}
	t.Log(res)
}
