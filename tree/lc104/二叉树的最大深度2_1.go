package lc104

// 分治法

func maxDepth2_1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// divide
	left := maxDepth2_1(root.Left)
	right := maxDepth2_1(root.Right)

	// conquer
	if left > right {
		return left+1
	}
	return right+1
}
