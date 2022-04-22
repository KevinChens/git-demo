package lc701

/*
分析：模拟
时间复杂度：O(n), 最多访问到叶子节点，深度为n
空间复杂度：O(1), 没有额外空间
 */

func insertIntoBST2(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	node := root

	for node != nil {
		if val < node.Val {
			if node.Left == nil {
				node.Left = &TreeNode{Val: val}
				break
			}
			node = node.Left
		} else {
			if node.Right == nil {
				node.Right = &TreeNode{Val: val}
				break
			}
			node = node.Right
		}
	}
	return root
}

