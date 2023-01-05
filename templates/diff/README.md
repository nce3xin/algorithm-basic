# 差分数组

## 背景

如果给你一个包含5000万个元素的数组，然后会有频繁区间修改操作，那什么是频繁的区间修改操作呢？比如让第1个数到第1000万个数每个数都加上1，而且这种操作时频繁的。

此时你应该怎么做？很容易想到的是，从第1个数开始遍历，一直遍历到第1000万个数，然后每个数都加上1，如果这种操作很频繁的话，那这种暴力的方法在一些实时的系统中可能就拉跨了。

## 介绍

差分和前缀和为逆运算。b数组称为a数组的差分。a就称为b的前缀和。

差分用来解决数组区间 [l,r] 都 +C(常数) 的情况。

差分数组定义：
$$
b_1 = a_1\\
b_2 = a_2-a_1\\
b_3 = a_3-a_2\\
...\\
b_n = a_n-a_{n-1}
$$
对数组 [l,r] 区间的每个数都加上一个常数c，不需要傻傻地遍历数组的[l，r]范围，然后再分别给每个值加上c，我们此时更改差分数组 b 即可：

```go
func insert(l, r, c int) {
	b[l] += c
	b[r+1] -= c
}
```

如何得到差分数组b：

假定a数组初始的时候全部是0，那么差分数组b也全部都是0。但是数组a不是0啊，怎么办？我们可以看作是a数组进行了n次插入操作，第一次是a数组 [1,1]+a1, 第二次是 [2,2]+a2，一直到最后 [n,n] + an。所以差分数组就不用去想如何构造了，直接复用上方的 insert() 函数即可。

## 题目

- 797

## 总结

差分只有1个操作，如果想给数组 [l,r] 区间整段都加一个数的话，那就让 b[l] += c, b[r+1] -= c 即可。

## 模板

![](imgs/diff.png)

# 二维差分

## 介绍

![](imgs/2d-diff.png)

## 二维差分模板

```
给以(x1, y1)为左上角，(x2, y2)为右下角的子矩阵中的所有元素加上c：
b[x1, y1] += c, b[x2 + 1, y1] -= c, b[x1, y2 + 1] -= c, b[x2 + 1, y2 + 1] += c
```


