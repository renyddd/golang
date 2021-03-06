package main

// ref: https://leetcode-cn.com/problems/valid-parentheses/

// isValidOfficial 来自官方解答
func isValidOfficial(s string) bool {
	// storeLeftStack 用栈来存储遇见的左半边括号
	storeLeftStack := make([]byte, 0)
	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, v := range s {
		// 遇见左括号
		if val, ok := pairs[byte(v)]; !ok {
			storeLeftStack = append(storeLeftStack, val)
		} else {
			storeLeftStack = storeLeftStack[:len(storeLeftStack)-1]
			if byte(v) != val {
				return false
			}
		}
	}

	return true
}

/*
 */

func isValid(s string) bool {
	if len(s)%2 > 0 {
		return false
	}
	nums := new(Kuohao)
	for _, v := range s {
		switch v {
		case '(':
			nums.Push(v)
		case '[':
			nums.Push(v)
		case '{':
			nums.Push(v)
		case ')':
			ok, item := nums.Pop()
			if ok {
				if item != '(' {
					return false
				}
			} else {
				return false
			}
		case ']':
			ok, item := nums.Pop()
			if ok {
				if item != '[' {
					return false
				}
			} else {
				return false
			}
		case '}':
			ok, item := nums.Pop()
			if ok {
				if item != '{' {
					return false
				}
			} else {
				return false
			}
		}

	}
	if nums.IsEmpty() {
		return true
	}
	return false
}

type Kuohao struct {
	data []rune
}

func (k *Kuohao) IsEmpty() bool {
	if len(k.data) == 0 {
		return true
	}
	return false
}

func (k *Kuohao) Push(x rune) {
	k.data = append(k.data, x)
}

func (k *Kuohao) Pop() (bool, rune) {
	if !k.IsEmpty() {
		v := k.data[len(k.data)-1]
		k.data = k.data[:len(k.data)-1]
		return true, v
	}
	return false, -1
}
