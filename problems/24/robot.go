package main

const N int = 3000

var res int
var n, m, k int
var dx = []int{-1, 0, 1, 0}
var dy = []int{0, 1, 0, -1}
var st [N][N]bool
var q [N]PII
var hh, tt int = 0, -1

type PII struct {
	x int
	y int
}

func movingCount(threshold, rows, cols int) int {
	if rows == 0 || cols == 0 {
		return 0
	}
	k = threshold
	n = rows
	m = cols
	bfs(0, 0)
	return res
}

func bfs(sx, sy int) {
	tt++
	q[tt] = PII{sx, sy}
	st[sx][sy] = true
	for hh <= tt {
		t := q[hh]
		hh++
		res++
		for i := 0; i < 4; i++ {
			x := t.x + dx[i]
			y := t.y + dy[i]
			if x < 0 || x >= n || y < 0 || y >= m {
				continue
			}
			if st[x][y] {
				continue
			}
			if sumDigit(x, y) > k {
				continue
			}
			tt++
			q[tt] = PII{x, y}
			st[x][y] = true
		}
	}
}

func sumDigit(x, y int) int {
	s := 0
	for x != 0 {
		s += x % 10
		x = x / 10
	}
	for y != 0 {
		s += y % 10
		y = y / 10
	}
	return s
}
