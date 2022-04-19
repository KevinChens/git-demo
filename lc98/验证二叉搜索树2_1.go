package lc98

import "math"

// 递归

func isValidBST2_1(root *TreeNode) bool {
	return helper2_1(root, math.MinInt64, math.MaxInt64)
}

func helper2_1(root *TreeNode, lower, upper int) bool {
	if root == nil {
		return true
	}

	if root.Val <= lower || root.Val >= upper {
		return false
	}
	return helper2_1(root.Left, lower, root.Val) && helper2_1(root.Right, root.Val, upper)
}
