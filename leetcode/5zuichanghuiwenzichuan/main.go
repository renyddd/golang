package main

/*
给你一个字符串 s，找到 s 中最长的回文子串。

示例 1：
输入：s = "babad"
输出："bab"
解释："aba" 同样是符合题意的答案。
示例 2：
输入：s = "cbbd"
输出："bb"
示例 3：
输入：s = "a"
输出："a"
示例 4：
输入：s = "ac"
输出："a"

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-palindromic-substring
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// longestPalindrome 以每一个遍历的字符为中心，向两边散开来判断并更新最长回文传
func longestPalindrome(s string) string {
	longest := ""
	for i := 0; i < len(s); i++ {
		s1 := palindrome(s, i, i)
		s2 := palindrome(s, i, i+1)

		if len(s1) > len(longest) {
			longest = s1
		}
		if len(s2) > len(longest) {
			longest = s2
		}
	}

	return longest
}

// palindrome 查找子回文串的方法
func palindrome(s string, l, r int) string {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}

	// 返回的是当前失败后，中间的包围串
	return s[l+1 : r]
}
