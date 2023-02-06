package main

import "fmt"

const (
	N = 510
	M = 1e4 + 10
)

var n, m, k int // k代表最短路径最多包含k条边

// Edge 把每个边保存下来即可。这个算法，边怎么存都行
type Edge struct {
	a, b, w int
}

var edges [M]Edge
var dist, backup [N]int // backup 备份数组防止串联

func main() {
	fmt.Scanf("%d%d%d", &n, &m, &k)
	for i := 0; i < m; i++ {
		var x, y, z int
		fmt.Scanf("%d%d%d", &x, &y, &z)
		edges[i] = Edge{x, y, z}
	}
	t := BellmanFord()
	if t > 0x3f3f3f3f/2 {
		fmt.Printf("impossible")
	} else {
		fmt.Printf("%d", t)
	}
}

func BellmanFord() int {
	// init dist[]
	for i := 0; i < N; i++ {
		dist[i] = 0x3f3f3f3f // infinity
	}
	dist[1] = 0
	for i := 0; i < k; i++ { //k次循环
		backup = dist
		for _, e := range edges { //遍历所有边
			a, b, w := e.a, e.b, e.w
			//使用backup: 每次只用上次的dist更新
			dist[b] = min(dist[b], backup[a]+w)
		}
	}
	return dist[n]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
