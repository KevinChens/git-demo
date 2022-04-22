package lc701

/*
分析：找到最后一个叶子节点满足插入条件即可
时间复杂度：O(n), 最多访问到叶子节点，深度为n
空间复杂度：O(n), 最多n层
 */

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		root = &TreeNode{Val: val}
		return root
	}
	if root.Val > val {
		root.Left = insertIntoBST(root.Left, val)
	} else {
		root.Right = insertIntoBST(root.Right, val)
	}
	return root
}

