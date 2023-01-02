package mergesort

var tmp []int = make([]int, 1e5+10)

func mergeSort(q []int, l, r int) {
	// 如果当前数组中只有一个数或者没有数，就不用排序了
	if l >= r {
		return
	}
	// mid取当前区间中点
	mid := (l + r) >> 1
	mergeSort(q, l, mid)
	mergeSort(q, mid+1, r)
	// merge
	k := 0           // k是已排好序的元素个数
	i, j := l, mid+1 // i,j分别指向左、右数组的起点
	// 当左右两边都没有循环空的时候
	for i <= mid && j <= r {
		// 每次把小的放在当前位置上
		if q[i] < q[j] {
			tmp[k] = q[i]
			k++
			i++
		} else {
			tmp[k] = q[j]
			k++
			j++
		}
	}
	// 如果左右两边有一边没有循环完的话，剩下的数直接接到答案里面去
	for i <= mid {
		tmp[k] = q[i]
		k++
		i++
	}
	for j <= r {
		tmp[k] = q[j]
		k++
		j++
	}
	// 把临时数组里的数，存回到q数组中去
	for i, j := l, 0; i <= r; i, j = i+1, j+1 {
		q[i] = tmp[j]
	}
}
