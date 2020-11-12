package main

import "fmt"

func main() {
	fmt.Println(lengthOfLastWord("a"))
}

func lengthOfLastWord(s string) int {
	var cnt int
	// 去除字符串以空格结尾的情况 为 true 则表示已遇见一个单词
	end := len(s) - 1
	if end < 0 {
		return 0
	}
	for s[end] == ' ' && end > 0 {
		end--
	}

	for i := end; i >= 0; i-- {
		if s[i] != ' '{
			cnt++
		} else {
			break
		}
	}
	return cnt
}