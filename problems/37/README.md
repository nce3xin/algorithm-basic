# 树的子结构

(二叉树，递归) O(nm)

代码分为两个部分：

1. 遍历树A中的所有非空节点R；
2. 判断树A中以R为根节点的子树是不是包含和树B一样的结构，且我们从根节点开始匹配；

对于第一部分，我们直接递归遍历树A即可，遇到非空节点后，就进行第二部分的判断。

对于第二部分，我们同时从根节点开始遍历两棵子树：

- 如果树B中的节点为空，则表示当前分支是匹配的，返回true；
- 如果树A中的节点为空，但树B中的节点不为空，则说明不匹配，返回false；
- 如果两个节点都不为空，但数值不同，则说明不匹配，返回false；
- 否则说明当前这个点是匹配的，然后递归判断左子树和右子树是否分别匹配即可；

## 时间复杂度

最坏情况下，我们对于树A中的每个节点都要递归判断一遍，每次判断在最坏情况下需要遍历完树B中的所有节点。

所以时间复杂度是O(nm), 其中n是树A中的节点数，m是树B中的节点数。