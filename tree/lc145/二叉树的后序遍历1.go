package lc145

/*
分析：递归
时间复杂度：O(n)
空间复杂度：O(n)
 */

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	postOrder(root, &result)
	return result
}

func postOrder(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}
	postOrder(root.Left, result)
	postOrder(root.Right, result)
	*result = append(*result, root.Val)
}
