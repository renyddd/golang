package main

/* package
给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。



示例 1：

输入：nums = [1,2,3]
输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
示例 2：

输入：nums = [0,1]
输出：[[0,1],[1,0]]
示例 3：

输入：nums = [1]
输出：[[1]]


提示：

1 <= nums.length <= 6
-10 <= nums[i] <= 10
nums 中的所有整数 互不相同

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/permutations
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

// 不错的题解：
//- https://leetcode-cn.com/problems/permutations/solution/hui-su-suan-fa-python-dai-ma-java-dai-ma-by-liweiw/
//- https://leetcode-cn.com/problems/permutations/solution/dai-ma-sui-xiang-lu-dai-ni-xue-tou-hui-s-mfrp/

func permute(nums []int) [][]int {
	res := make([][]int, 0)
	nl := len(nums)
	rootPath := make([]int, 0)
	rootMet := make([]bool, nl)
	for i := 0; i < nl; i++ {
		rootMet[i] = false
	}

	var backtrack func([]int, []bool)
	backtrack = func(path []int, meetnums []bool) {
		if len(path) == nl {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
		}

		for i, v := range nums {
			if !meetnums[i] {
				// 此处类似树的前序操作，为下次递归作准备
				meetnums[i] = true
				path = append(path, v)

				// backtrack
				backtrack(path, meetnums)

				// 此处类似树的后序操作，撤销上次递归的影响
				path = path[:len(path)-1]
				meetnums[i] = false
			}
		}
	}

	backtrack(rootPath, rootMet)

	return res
}

// func permute(nums []int) [][]int {
// 	res := make([][]int, 0)
// 	nl := len(nums)

// 	numsM := make(map[int]struct{})
// 	for _, v := range nums {
// 		numsM[v] = struct{}{}
// 	}

// 	map2slice := func(m map[int]struct{}) []int {
// 		n := make([]int, 0)
// 		for k, _ := range m {
// 			n = append(n, k)
// 		}

// 		return n
// 	}

// 	copyMap := func(dst, src map[int]struct{}) { // avoid copy pointer
// 		for k, v := range src {
// 			dst[k] = v
// 		}
// 	}

// 	var backtrack func(map[int]struct{}, map[int]struct{})
// 	backtrack = func(made, choice map[int]struct{}) {
// 		if len(made) == nl {
// 			res = append(res, map2slice(made)) // 不可行，这里不能保证顺序准确
// 			return
// 		}

// 		tmpChoice := make(map[int]struct{}, 0)
// 		copyMap(tmpChoice, choice)

// 		for key, _ := range choice {
// 			// 此处类似树的前序操作
// 			made[key] = struct{}{}
// 			delete(tmpChoice, key)

// 			backtrack(made, tmpChoice)
// 		}
// 	}

// 	emptyMade := make(map[int]struct{}, 0)
// 	backtrack(emptyMade, numsM)

// 	return res
// }
