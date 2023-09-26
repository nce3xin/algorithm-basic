package main

import "fmt"

const N int = 21 // 最大单词数
var n int        // 单词数
var word [N]string
var g [N][N]int // 两个单词重叠部分的最小长度，要大于0且严格小于2个串的长度
var used [N]int // 每个单词当前用了多少次，因为规定每个单词最多用2次

var ans int

func main() {
	fmt.Scanf("%d", &n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%s", &word[i])
	}
	var start byte
	fmt.Scanf("%c", &start)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			a := word[i]
			b := word[j]
			// 重叠部分的长度越短越好，这样才能使接龙后的长度更长，所以这里从小到大搜
			for k := 1; k < min(len(a), len(b)); k++ {
				if a[len(a)-k:len(a)] == b[:k] {
					g[i][j] = k
					break
				}
			}
		}
	}

	// 爆搜
	for i := 0; i < n; i++ {
		// 如果单词的开头字母等于start
		if word[i][0] == start {
			// 从word[i]开始搜
			dfs(word[i], i) // 记一下当前最后一个单词是i，方便判断下一个单词能不能接
		}
	}
	fmt.Printf("%d\n", ans)
}

// last: 上一个单词编号
func dfs(dragon string, last int) {
	ans = max(ans, len(dragon))

	used[last]++
	// 下一个单词该接哪个
	for i := 0; i < n; i++ {
		if g[last][i] != 0 && used[i] < 2 {
			overlapLen := g[last][i]
			dfs(dragon+word[i][overlapLen:], i)
		}
	}
	// 恢复现场
	used[last]--
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
