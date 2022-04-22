package lc94

/*
分析：非递归，借助stack, 通过stack保存已经访问的元素
时间复杂度：O(n)
空间复杂度：O(n)
 */

func inorderTraversal2(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	result := make([]int, 0)
	stack := make([]*TreeNode, 0)

	for root != nil || len(stack) > 0 {
		for root != nil {
			// push
			stack = append(stack, root)
			root = root.Left
		}
		// pop
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		result = append(result, node.Val)

		root = node.Right
	}
	return result
}
