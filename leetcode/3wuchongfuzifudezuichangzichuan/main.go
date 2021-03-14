package main

// https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/

func FindLongest(S string) int {
	// 临时性存储结果
	res := 0

	for i, _ := range S {
		tmpRes := make([]uint8, 0)
		existMap := make(map[uint8]struct{}, 0)

		pi := i
		for pi < len(S) {
			if _, exist := existMap[S[pi]]; exist {
				break
			}
			existMap[S[pi]] = struct{}{}
			tmpRes = append(tmpRes, S[pi])
			pi++
		}

		if len(tmpRes) > res {
			res = len(tmpRes)
		}
	}

	return res
}

// FindLongest2 用指针指出临时变量的内容
func FindLongest2(S string) []uint8 {
	// 临时性存储结果
	var res *[]uint8
	t := make([]uint8, 0)
	res = &t

	for i, _ := range S {
		tmpRes := make([]uint8, 0)
		existMap := make(map[uint8]struct{}, 0)

		pi := i
		for pi < len(S) {
			if _, exist := existMap[S[pi]]; exist {
				break
			}
			existMap[S[pi]] = struct{}{}
			tmpRes = append(tmpRes, S[pi])
			pi++
		}

		if len(tmpRes) > len(*res) {
			res = &tmpRes
		}
	}

	return *res
}

// FindLongest3 使用滑动窗口
// ref: https://github.com/labuladong/fucking-algorithm/blob/master/%E7%AE%97%E6%B3%95%E6%80%9D%E7%BB%B4%E7%B3%BB%E5%88%97/%E6%BB%91%E5%8A%A8%E7%AA%97%E5%8F%A3%E6%8A%80%E5%B7%A7.md
func FindLongest3(s string) int {
	var left, right, res int
	// existWindow 需要的是出现第二次的信息，如果单单只是 bool 类型的 map value，无法做出次区分
	existWindow := make(map[uint8]int, 0)

	for right < len(s) {
		pC := s[right]
		existWindow[pC]++
		right++

		for existWindow[pC] > 1 {
			tC := s[left]
			existWindow[tC]--
			left++
		}

		if v := right - left; v > res {
			res = v
		}
	}

	return res
}
