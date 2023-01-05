# 逆序对个数

首先我们给出逆序对的定义：对于数列的第 i 个和第 j 个元素，如果满足 i < j 且 a[i] > a[j]，则其为一个逆序对。

那么第二步是分析问题，这里我们可以使用分治法解决问题。

我们将序列从中间分开，将逆序对分成三类：

- 两个元素都在左边
- 两个元素都在右边
- 两个元素一个在左一个在右
- 因此这就是我们算法的大致框架

计算逆序对的数量（序列）：
1. 递归算左边的；
2. 递归算右边的；
3. 算一个左一个右的；
4. 把他们加到到一起。

这个时候我们注意到一个很重要的性质，左右半边的元素在各自内部任意调换顺序，是不影响第三步计数的，因此我们可以数完就给它排序。这么做的好处在于，如果序列是有序的，会让第三步计数很容易。如果无序暴力数的话这一步是O(n^2)的。

比如序列是这样的：

```
2 3 4 5 6 | 1 2 3
```

当你发现左边的4 比右边的 3 大的时候，那么左边剩下的5和6都必然比右边的所有元素大，因此就可以不用数5和6的情形了，直接分别加上右半边的元素个数就可以了，这一步就降低到了O(n), 我们知道排序是O(nlogn)，所以成本是可以接受的，并且这一问题下，可以很自然地使用归并排序。