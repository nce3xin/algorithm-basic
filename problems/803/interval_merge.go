package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const N int = 1e5 + 10

type PII pair

type pair struct {
	first  int
	second int
}

func main() {
	var n int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := strings.Split(scanner.Text(), " ")
	n, _ = strconv.Atoi(line[0])
	intervals := make([]PII, 0)
	for i := 0; i < n; i++ {
		var l, r int
		scanner.Scan()
		line = strings.Split(scanner.Text(), " ")
		l, _ = strconv.Atoi(line[0])
		r, _ = strconv.Atoi(line[1])
		intervals = append(intervals, PII{l, r})
	}
	res := mergeIntervals(intervals)
	fmt.Printf("%d", len(res))
}

func mergeIntervals(intervals []PII) []PII {
	// 排序后的区间之间关系，只有3种
	// 	按左端点排序
	sortPairs(intervals)
	// start和end维护一个合并区间，初始只要维护一个空区间
	// 由于本题数据范围是[-1e9,+1e9], 所以只要搞一个比-1e9小的数作为初始值即可
	var start, end int = -2e9, -2e9
	res := make([]PII, 0)
	for _, interval := range intervals {
		// 当前维护的区间的右端点小于下一个区间的左端点，没有交集
		if end < interval.first {
			// 判断，不要把初始化的最左边的空区间加进结果里了
			if start != -2e9 {
				res = append(res, PII{start, end})
			}
			start = interval.first
			end = interval.second
		} else { // 有交集
			end = max(end, interval.second)
		}
	}
	// 这里之所以要判断start != -2e9，是因为当输入为空时，{-2e9,+2e9} 会被加入结果，实际不应该加入
	// 所以要排除这种情况
	if start != -2e9 {
		res = append(res, PII{start, end})
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func sortPairs(a []PII) {
	sort.Slice(a, func(i, j int) bool {
		return a[i].first < a[j].first
	})
}
