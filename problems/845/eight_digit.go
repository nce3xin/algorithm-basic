package main

import (
	"fmt"
	"strings"
)

var q []string               // queue，节点用string表示（把3*3矩阵拉成一个字符串）
var d = make(map[string]int) // 每个节点到原点的距离，节点用string表示（把3*3矩阵拉成一个字符串）

func main() {
	var start string
	for i := 0; i < 9; i++ {
		var c string
		fmt.Scanf("%s", &c)
		start += c
	}
	fmt.Printf("%d", bfs(start))
}

func bfs(start string) int {
	end := "12345678x"
	// 起点入队
	q = append(q, start)
	d[start] = 0 // 起点距离是0
	dx := []int{-1, 0, 1, 0}
	dy := []int{0, 1, 0, -1}
	for len(q) > 0 { // queue is not empty
		// 取出队头元素
		t := q[0]
		q = q[1:]
		//记录当前状态的距离，如果是最终状态则返回距离
		// 这里必须要开一个变量存d[t]，如果不用distance用d[t]，会Wrong Answer
		distance := d[t]
		if t == end { // 如果走到目标
			return distance
		}
		k := find(t, "x") // k是队头节点的一维坐标
		x := k / 3        // 一维转二维坐标小技巧
		y := k % 3        // 一维转二维坐标小技巧
		// 扩展t的上下左右四个方向
		for i := 0; i < 4; i++ {
			// a,b是扩展后的点的坐标（上下左右四个点中的任意一个）
			a := x + dx[i]
			b := y + dy[i]
			if a >= 0 && a < 3 && b >= 0 && b < 3 { // 没有出界
				// swap (x,y) and (a,b), 状态转移。二维坐标转一维坐标
				t = swap(t, k, 3*a+b)
				//如果当前状态是第一次遍历，记录距离，入队
				if _, ok := d[t]; !ok {
					// 更新距离
					d[t] = distance + 1 // 这儿必须要用distance，不能用d[t] = d[t] + 1,因为交换之后t已经变了
					// 扩展后的元素入队
					q = append(q, t)
				}
				// 还原状态，为下一种转换情况做准备
				t = swap(t, k, 3*a+b)
			}
		}
	}
	return -1
}

// 在字符串s中寻找字符'x'出现的位置，返回其下标
func find(s string, c string) int {
	return strings.Index(s, c)
}

func swap(s string, i, j int) string {
	char := []byte(s)
	char[i], char[j] = char[j], char[i]
	return string(char)
}
