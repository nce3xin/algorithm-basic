# 递归实现组合型枚举 II

和1572类似，按照元素的种类枚举每种该选几个。

![](imgs/1.png)

参数：

- 当前枚举到了第几段。其实，u表示的不是第几段，而是每一段的第一个数在整个a数组里的下标。
- 哪些数被选了：bool数组
- 当前一共选了几个数。