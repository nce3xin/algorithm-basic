package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type PII pair

type pair struct {
	first  int
	second int
}

// N 坐标x的数量上限为1e5，两个坐标l,r的数量上限也为1e5,所以加起来为3*le5;
const N int = 3e5 + 10

// a[] 存储坐标插入的值
var a []int = make([]int, N)

// s[] 存储数组a的前缀和
var s []int = make([]int, N)

func main() {
	var n, m int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := strings.Split(scanner.Text(), " ")
	n, _ = strconv.Atoi(line[0])
	m, _ = strconv.Atoi(line[1])
	// add[], query[] 存储插入和询问操作的数据
	add := make([]PII, 0)
	query := make([]PII, 0)
	// alls[] 存储（所有与插入和查询有关的）坐标
	alls := make([]int, 0)
	for i := 0; i < n; i++ {
		var x, c int
		scanner.Scan()
		line = strings.Split(scanner.Text(), " ")
		x, _ = strconv.Atoi(line[0])
		c, _ = strconv.Atoi(line[1])
		add = append(add, PII{x, c})

		alls = append(alls, x)
	}
	for i := 0; i < m; i++ {
		var l, r int
		scanner.Scan()
		line = strings.Split(scanner.Text(), " ")
		l, _ = strconv.Atoi(line[0])
		r, _ = strconv.Atoi(line[1])
		query = append(query, PII{l, r})

		alls = append(alls, l)
		alls = append(alls, r)
	}

	// alls[]数组排序、去重
	sortInts(alls)
	alls = unique(alls)
	// 执行插入数据操作
	for _, item := range add {
		x := find(item.first, alls)
		a[x] += item.second
	}

	// 预处理前缀和
	for i := 1; i <= len(alls); i++ {
		s[i] = s[i-1] + a[i]
	}

	// 处理询问
	for _, item := range query {
		l := find(item.first, alls)
		r := find(item.second, alls)
		res := s[r] - s[l-1]
		fmt.Printf("%d\n", res)
	}
}

func sortInts(a []int) {
	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})
}

// 返回的是输入的坐标的离散化下标
func find(x int, alls []int) int {
	l, r := 0, len(alls)-1
	for l < r {
		mid := (l + r) >> 1
		if alls[mid] >= x {
			r = mid
		} else {
			l = mid + 1
		}
	}
	// 因为后续要使用前缀和，所以返回的坐标要加上1, 映射到从1开始的下标
	return l + 1
}

func unique(a []int) []int {
	// 双指针
	res := make([]int, len(a))
	// j 已经存了多少个数
	j := 0
	for i := 0; i < len(a); i++ {
		// i==0表示是首个元素
		if i == 0 || a[i] != a[i-1] {
			res[j] = a[i]
			j++
		}
	}
	res = res[:j]
	return res
}
