package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// N 如果是1e5,就会报错在32行，panic: runtime error: index out of range [1] with length 1
const N int = 1e6 + 10

var a []int = make([]int, N)
var b []int = make([]int, N)

func main() {
	var n, m int
	// fmt.Scanf() 会TLE
	scanner := bufio.NewScanner(os.Stdin)
	// 必须重设scanner的buf大小，默认最大64k，而n最大是1e5, 大于64k
	buf := make([]byte, N)
	scanner.Buffer(buf, N)

	scanner.Scan()
	line := strings.Split(scanner.Text(), " ")
	n, _ = strconv.Atoi(line[0])
	m, _ = strconv.Atoi(line[1])
	scanner.Scan()
	line = strings.Split(scanner.Text(), " ")
	for i := 1; i <= n; i++ {
		a[i], _ = strconv.Atoi(line[i-1])
	}
	for i := 1; i <= n; i++ {
		insert(i, i, a[i])
	}
	for ; m > 0; m-- {
		var l, r, c int
		scanner.Scan()
		line = strings.Split(scanner.Text(), " ")
		l, _ = strconv.Atoi(line[0])
		r, _ = strconv.Atoi(line[1])
		c, _ = strconv.Atoi(line[2])
		insert(l, r, c)
	}
	for i := 1; i <= n; i++ {
		b[i] += b[i-1]
	}
	for i := 1; i <= n; i++ {
		fmt.Printf("%d ", b[i])
	}
}

func insert(l, r, c int) {
	b[l] += c
	b[r+1] -= c
}
