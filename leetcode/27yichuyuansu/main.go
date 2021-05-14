package main

import "fmt"

func main() {
	n := []int{1, 2, 2, 3, 3, 3, 2, 3, 2, 3, 2, 0}
	removeElement(n, 3)
	fmt.Println(n)
}

func removeElement(nums []int, val int) int {
	var fast, slow int
	for fast < len(nums) {
		if nums[fast] == val {
			fast++
		} else {
			nums[slow] = nums[fast]
			fast++
			slow++
		}
	}
	return slow
}
