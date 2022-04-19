package lc98

import "math"

// 非递归的中序遍历

 func isValidBST1_1(root *TreeNode) bool {
 	stack := []*TreeNode{}
 	inorder := math.MinInt64

 	for root != nil || len(stack) > 0 {
 		for root != nil {
 			// push
 			stack = append(stack, root)

 			root = root.Left
		}
		// pop
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if node.Val <= inorder {
			return false
		}
		inorder = node.Val

		root = node.Right
	}
	return true
 }
