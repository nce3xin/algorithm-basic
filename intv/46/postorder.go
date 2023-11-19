package main

func verifySequenceOfBST(sequence []int) bool {
	seq := sequence
	return dfs(seq, 0, len(seq)-1)
}

func dfs(seq []int, l, r int) bool {
	if l >= r {
		return true
	}
	root := seq[r]
	k := l
	// 从数组头开始遍历
	for k < r && seq[k] < root {
		k++
	}
	for ; k < r; k++ {
		if seq[k] < root {
			return false
		}
	}
	return dfs(seq, l, k-1) && dfs(seq, k, r-1)
}
