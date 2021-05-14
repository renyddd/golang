package main

import "fmt"

func Traverse(a [][]int) {
	if a == nil {
		return
	}
	rows, colums := len(a), len(a[0])
	totola := rows * colums
	mVisited := make([][]bool, rows)
	cur, i, j := 0, 0, 0

	for cur < totola {
		//right
		for ; i < colums && !mVisited[j][i]; i++ {
			cur++
			fmt.Println(a[j][i])
		}
		//down
		for ; j < rows && !mVisited[j][i]; j++ {
			cur++
			fmt.Println(a[j][i])
		}
	}
}
