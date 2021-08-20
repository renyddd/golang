package main

/*
给定一个整数数组  nums，求出数组从索引 i 到 j（i ≤ j）范围内元素的总和，包含 i、j 两点。

实现 NumArray 类：

NumArray(int[] nums) 使用数组 nums 初始化对象
int sumRange(int i, int j) 返回数组 nums 从索引 i 到 j（i ≤ j）范围内元素的总和，包含 i、j 两点（也就是 sum(nums[i], nums[i + 1], ... , nums[j])）

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/range-sum-query-immutable
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

type NumArray struct {
	data []int
}

func Constructor(nums []int) NumArray {
	tmp := make([]int, len(nums)+1)
	var n NumArray

	for i, v := range nums {
		tmp[i+1] = tmp[i] + v
	}

	n.data = tmp

	return n
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.data[right+1] - this.data[left]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
