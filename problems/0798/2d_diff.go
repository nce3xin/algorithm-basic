package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const N int = 1000 + 10

var a [][]int = make([][]int, N)
var b [][]int = make([][]int, N)

func main() {
	// init a[] and b[]
	for i := 0; i < N; i++ {
		a[i] = make([]int, N)
		b[i] = make([]int, N)
	}
	// input
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := strings.Split(scanner.Text(), " ")
	var n, m, q int
	n, _ = strconv.Atoi(line[0])
	m, _ = strconv.Atoi(line[1])
	q, _ = strconv.Atoi(line[2])
	for i := 1; i <= n; i++ {
		scanner.Scan()
		line = strings.Split(scanner.Text(), " ")
		for j := 1; j <= m; j++ {
			a[i][j], _ = strconv.Atoi(line[j-1])
		}
	}

	// get b[]
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			insert(i, j, i, j, a[i][j])
		}
	}

	for ; q > 0; q-- {
		var x1, y1, x2, y2, c int
		scanner.Scan()
		line = strings.Split(scanner.Text(), " ")
		x1, _ = strconv.Atoi(line[0])
		y1, _ = strconv.Atoi(line[1])
		x2, _ = strconv.Atoi(line[2])
		y2, _ = strconv.Atoi(line[3])
		c, _ = strconv.Atoi(line[4])

		insert(x1, y1, x2, y2, c)
	}

	// calculate matrix after adding
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			// b[][] 是 a[][] 的差分，a[][] 就是b[][] 的前缀和
			// 也就是 a[i][j]等于其左上角所有b[][]的和
			// 这里用b[i][j]来存a[i][j]的值，可以使代码更简洁
			// 类似容斥原理，多加的要减去
			b[i][j] += b[i-1][j] + b[i][j-1] - b[i-1][j-1]
		}
	}

	// output
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			fmt.Printf("%d ", b[i][j])
		}
		fmt.Printf("\n")
	}
}

func insert(x1, y1, x2, y2, c int) {
	b[x1][y1] += c
	b[x2+1][y1] -= c
	b[x1][y2+1] -= c
	b[x2+1][y2+1] += c
}
