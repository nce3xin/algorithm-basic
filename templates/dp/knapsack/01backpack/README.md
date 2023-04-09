# 01背包问题

给N个物品和一个容量是V的背包，每个物品两个属性：

1. 体积：vi （i是下标）
2. 价值：wi (为什么用w来表示呢，因为习惯用w来表示权重，这里的价值其实是权重的意思)

每件物品最多用一次。用0次或者用1次。这也是01背包这个问题的由来。

问题是：我们要从这些物品里挑一些物品，使得总体积小于等于V。也就是背包能装的下的情况下，能挑出来的最大价值是多少。

## 特点

每件物品最多只能用一次。

## 题目

- 2

## 思路

- 状态表示。集合：从前i个物品中选，总体积<=j的所有选法。集合的属性：所有选法中，价值的最大值。
- 状态计算。看下面这张图的最底下，集合划分成2个子集：含i和不含i。注意左边这种情况是一定存在的，但是右边这种情况不一定存在。当$j<v_i$时（当j这个体积装不下第i个物品时），右边就是空集。
- 初始化。初始化要枚举所有的状态，也就是f[0-n][0-m]。f[0][0-m]表示的是考虑0件物品，总体积不超过0，不超过1…一直到不超过m，这样的情况下，最大价值是多少。由于1件物品都没选，所以f[0][0-m]都是0. 由于f数组定义成了全局变量，它默认就是0。所以f[0][0-m]=0这句初始化就可以不写。所以i从1开始枚举（因为f[0][0-m]都已经全部是0了），然后枚举所有的体积。

![](imgs/1.png)