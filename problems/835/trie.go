package main

import "fmt"

const N int = 1e5 + 10

// son[父亲的位置][儿子的名字]=儿子的位置
// 比如[0][1]=3, [0]表示根节点，[1]表示它有一个儿子'b',这个儿子的下标是3
// 下标是0的点，既是根节点，又是空节点
var son [N][26]int
var idx int    // 当前用到哪个下标了
var cnt [N]int // 以当前字符结尾的单词有多少个

func main() {
	var n int
	fmt.Scanf("%d", &n)
	for i := 0; i < n; i++ {
		var op, s string
		fmt.Scanf("%s%s", &op, &s)

		if op == "I" {
			insert(s)
		} else if op == "Q" {
			count := query(s)
			fmt.Printf("%d\n", count)
		}
	}
}

func insert(s string) {
	p := 0 // 从根节点开始
	for _, c := range s {
		u := c - 'a'
		// 如果不存在当前字母，创建一个
		if son[p][u] == 0 {
			idx++
			son[p][u] = idx
		}
		p = son[p][u]
	}
	cnt[p]++
}

func query(s string) int {
	p := 0 // 从根节点开始
	for _, c := range s {
		u := c - 'a'
		// 不存在当前字符
		if son[p][u] == 0 {
			return 0
		}
		p = son[p][u]
	}
	return cnt[p]
}
