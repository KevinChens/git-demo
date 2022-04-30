package lc145

/*
分析：非递归，利用stack，lastVisit记录是否访问，保证root在Right弹出之后弹出
时间复杂度：O(n)
空间复杂度：O(n)
 */

func postorderTraversal2(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	result := make([]int, 0)
	stack := make([]*TreeNode, 0)
	var lastVisit *TreeNode

	for root != nil || len(stack) > 0 {
		for root != nil {
			// push
			stack = append(stack, root)

			root = root.Left
		}
		// read
		node := stack[len(stack)-1]
		// 判断是否pop
		if node.Right == nil || node.Right == lastVisit {
			// pop
			stack = stack[:len(stack)-1]
			
			result = append(result, node.Val)
			// mark
			lastVisit = node
		} else {
			root = node.Right
		}
	}
	return result
}
