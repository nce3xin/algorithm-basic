package main

const N int = 1010

func longestSubstringWithoutDuplication(s string) int {
	n := len(s)
	res := 0
	// 每个字母出现的次数
	ss := make(map[byte]int)
	for i, j := 0, 0; i < n; i++ {
		ss[s[i]]++
		for ss[s[i]] > 1 {
			ss[s[j]]--
			j++
		}
		res = max(res, i-j+1)
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
