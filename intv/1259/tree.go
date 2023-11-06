package main

import "fmt"

type TreeNode struct {
	Val   byte // byte, not int
	Left  *TreeNode
	Right *TreeNode
}

var res string // 最终结果，前序遍历

func main() {
	// 中序遍历和层序遍历
	var inorder, lorder string
	fmt.Scanf("%s", &inorder)
	fmt.Scanf("%s", &lorder)
	// 哈希表存中序遍历里每个元素的下标
	pos := make(map[byte]int)
	for i := 0; i < len(inorder); i++ {
		pos[inorder[i]] = i
	}
	// 记录中序遍历里每个点有没有被用过，最大长度是26（不相同字母最多有26个），开30够用了
	var st [30]bool
	// 把宽搜的过程还原出来
	// 需要定义一个队列，维护一下当前这层处理的节点是什么
	var q [30]*TreeNode
	// 层序遍历恢复出来
	q[0] = &TreeNode{lorder[0], nil, nil}
	// 按照宽搜的顺序，把层序遍历搜出来，一层一层来做
	// 按层遍历，i是当前这层的起点，j是下一层的起点
	for i, j := 0, 1; j < len(lorder); {
		// 因为j在不断增加，所以这里需要先把下一层的起点j先存起来
		for end := j; i < end; i++ { // 遍历当前这层
			// 判断左儿子是否存在
			p := pos[lorder[i]] // 当前节点在层序遍历中的下标
			st[p] = true
			if p > 0 && !st[p-1] { // 当前节点的左边有值且这个值没有被访问过，说明左儿子存在
				q[i].Left = &TreeNode{
					Val:   lorder[j],
					Left:  nil,
					Right: nil,
				}
				// 左儿子放到队列里面
				q[j] = q[i].Left
				j++
			}
			// 判断右儿子是否存在
			if p+1 < len(lorder) && !st[p+1] {
				q[i].Right = &TreeNode{lorder[j], nil, nil}
				// 右儿子放到队列里面
				q[j] = q[i].Right
				j++
			}
		}
	}
	dfs(q[0]) // 先序遍历根节点
	fmt.Printf("%s\n", res)
}

func dfs(root *TreeNode) {
	if root == nil {
		return
	}
	res += string(root.Val)
	dfs(root.Left)
	dfs(root.Right)
}
