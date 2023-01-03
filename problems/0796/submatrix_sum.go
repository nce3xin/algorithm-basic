package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const N int = 1000 + 10

func main() {
	var n, m, q int
	var a, s [N][N]int
	fmt.Scanf("%d%d%d", &n, &m, &q)
	scanner := bufio.NewScanner(os.Stdin)
	for i := 1; i <= n; i++ {
		scanner.Scan()
		line := strings.Split(scanner.Text(), " ")
		for j := 1; j <= m; j++ {
			a[i][j], _ = strconv.Atoi(line[j-1])
		}
	}
	// calculate prefixSum[][]
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			s[i][j] = s[i-1][j] + s[i][j-1] - s[i-1][j-1] + a[i][j] // 求前缀和
		}
	}
	var x1, y1, x2, y2 int
	var ll [4]int
	for ; q > 0; q-- {
		scanner.Scan()
		temp := strings.Split(scanner.Text(), " ")
		for i, s := range temp {
			ll[i], _ = strconv.Atoi(s)
		}
		x1, y1, x2, y2 = ll[0], ll[1], ll[2], ll[3]

		res := s[x2][y2] - s[x2][y1-1] - s[x1-1][y2] + s[x1-1][y1-1] // 算子矩阵和

		fmt.Printf("%d\n", res)
	}
}
