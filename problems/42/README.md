# 栈的压入、弹出序列

输入两个整数序列，第一个序列表示栈的压入顺序，请判断第二个序列是否可能为该栈的弹出顺序。

用一个新栈s来模拟实时进出栈操作：

- 在for loop里依次喂数，每push一个数字就检查有没有能pop出来的。
- 如果最后s为空，说明一进一出刚刚好。

时间复杂度分析：O(N)
