package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const N int = 1e5 + 10 // 点数
var n, m int
var p [N]int // 并查集

type Edge struct {
	a, b, w int
}

var edges []Edge // 边数是点数的2倍, 但这儿不能直接设长度是2*N，因为sort函数只能用于slice，不能用于array

func main() {
	r := bufio.NewReader(os.Stdin)
	fmt.Fscan(r, &n, &m)
	for i := 0; i < m; i++ {
		var a, b, c int
		fmt.Fscan(r, &a, &b, &c)
		edges = append(edges, Edge{
			a: a,
			b: b,
			w: c,
		})
	}

	// sort edges ascend
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].w < edges[j].w
	})

	// init p (union find)
	for i := 1; i <= n; i++ {
		p[i] = i
	}
	res := 0 // 最小生成树中所有边的权重总和
	cnt := 0 // 最小生成数中边的总数，用于判断是否找到最小生成树
	// traverse all edges ascend
	for i := 0; i < m; i++ {
		a := edges[i].a
		b := edges[i].b
		c := edges[i].w
		pa := find(a)
		pb := find(b)
		if pa != pb { // a和b不连通
			res += c
			cnt++
			p[pa] = pb
		}
	}

	if cnt < n-1 { // 不存在最小生成树
		fmt.Printf("impossible\n")
	} else {
		fmt.Printf("%d", res)
	}
}

func find(x int) int {
	if p[x] != x { // 不是根节点
		p[x] = find(p[x])
	}
	return p[x]
}
