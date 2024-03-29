# 转一维

## 回顾二维写法

```
for i := 1; i <= n; i++ {
	for j := 0; j <= m; j++ {
		f[i][j] = f[i-1][j]
		if j >= v[i] {
			f[i][j] = max(f[i][j], f[i-1][j-v[i]]+w[i])
		}
	}
}
fmt.Printf("%d\n", f[n][m])
```

## 为什么能转一维

- 注意到，f[i]只用到了f[i-1]这一层，f[i-2]到f[0]是没有用到的。所以可以用滚动数组来做。把第一维直接删掉。
- 又注意到，f[j]只用到了j和j-v[i], 它们都是小于等于j的（非常重要，说明只用到了上一层，没有用到下一层），没有说在j的两侧。所以就可以改成一维数组来算。

## 回到代码

```
f[i][j] = f[i-1][j]
// 删掉第一维，变成如下恒等式，这行可以直接删掉
f[j] = f[j]
```

```
for i := 1; i <= n; i++ {
	for j := 0; j <= m; j++ {
		f[i][j] = f[i-1][j]
		// 当 j 小于 v[i]， if判断为假，可以直接划过去。所以 j 可以直接从v[i]开始循环。判断也可以删掉。
		if j >= v[i] {
			f[i][j] = max(f[i][j], f[i-1][j-v[i]]+w[i])
		}
	}
}
fmt.Printf("%d\n", f[n][m])
```

但这句：

```
f[i][j] = max(f[i][j], f[i-1][j-v[i]]+w[i])
```

如果直接删掉第一维，和之前就不等价了:

```
f[j] = max(f[j], f[j-v[i]]+w[i])
```

因为 j-v[i] < j, j 又是从小到大枚举的，所以当我们更新f[j]时我们已经把f[j - v[i]]更新了。也就是说我们使用的是第[i]个循环里面的f[j - v[i]]，而我们实际要用的是第[i-1]个循环里面的f[j - v[i]]。怎么解决这个问题呢？把内层 j 的循环顺序变成从大到小就可以了。这样在算f[j-v[i]]的时候，由于是从大到小枚举的所有体积，在算f[j]的时候，f[j-v[i]]还没有被更新过。那么它存的就是第i-1层的j-v[i]，也就是f[i-1][j-v[i]]。所以和原状态转移方程等价。

如果不改j的判断条件为 j>=v[i]，下面这样写也能AC:

```
package main

import "fmt"

const N int = 1010

// 在二维版本中，f[i]只用到了f[i-1]这一层，所以可以用滚动数组来做, 直接把第一维删掉
var f [N]int    // 状态，一维优化版
var n, m int    // n表示物品个数，m表示背包容量
var v, w [N]int // v表示物品的体积，w表示物品的价值

// 一维优化版

func main() {
	fmt.Scanf("%d%d", &n, &m)
	for i := 1; i <= n; i++ {
		fmt.Scanf("%d%d", &v[i], &w[i])
	}
	for i := 1; i <= n; i++ {
		for j := m; j >=0; j-- {
		    if j >= v[i] {
		        f[j] = max(f[j], f[j-v[i]]+w[i])
		    }
		}
	}
	fmt.Printf("%d\n", f[m])
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

```
