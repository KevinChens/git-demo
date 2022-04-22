package lc102

/*
分析：用queue记录一层的元素，然后扫描这一层元素添加下一层元素到queue（一个数进去出来一次）
时间复杂度：O(n), 每个节点访问一次
空间复杂度：O(n), queue最多为n个节点
 */

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	result := make([][]int ,0)
	queue := make([]*TreeNode, 0)

	// push
	queue = append(queue, root)

	for len(queue) > 0{
		// 保存当前层元素
		list := make([]int, 0)
		// 记录当前层的元素个数，遍历当前层，再添加下一层
		l := len(queue)
		for i := 0; i < l; i++ {
			// pop
			node := queue[0]
			queue = queue[1:]

			list = append(list, node.Val)
			// push
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, list)
	}
	return result
}
