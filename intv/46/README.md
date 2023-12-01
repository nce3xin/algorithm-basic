# 二叉搜索树的后序遍历序列

剑指offer。

输入一个整数数组，判断该数组是不是某二叉搜索树的后序遍历的结果。

如果是则返回true，否则返回false。

假设输入的数组的任意两个数字都互不相同。

## 解法

后序遍历，数组最后一个数是根节点的值root。

对于二叉搜索树，左子树的值都比根小，右子树的值都比根大。所以看是否能把数组分成2半，左边都比root小，右边都比root大。