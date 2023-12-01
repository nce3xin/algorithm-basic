package main

func numIslands(grid [][]byte) int {
	n := len(grid)
	m := len(grid[0])
	res := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == '0' {
				continue
			}
			dfs(i, j, n, m, grid)
			res++
		}
	}
	return res
}

func dfs(sx, sy, n, m int, grid [][]byte) {
	if sx < 0 || sx >= n || sy < 0 || sy >= m {
		return
	}
	if grid[sx][sy] == '0' {
		return
	}
	grid[sx][sy] = '0'
	dx := []int{-1, 0, 1, 0}
	dy := []int{0, -1, 0, 1}
	for i := 0; i < 4; i++ {
		x := sx + dx[i]
		y := sy + dy[i]
		dfs(x, y, n, m, grid)
	}
}
