package lc110

/*
分析：自顶向下，递归
时间复杂度：O(n^2)，最坏情况下，二叉树是满二叉树，需要遍历二叉树中的所有节点，时间复杂度是O(n)。
对于节点p，如果它的高度是d，则height(p)最多会被调用d次（即遍历到它的每一个祖先节点时）。
对于平均的情况，一棵树的高度h满足O(h)=O(logn)，因为d≤h，所以总时间复杂度为O(nlogn)。
对于最坏的情况，二叉树形成链式结构，高度为O(n)，此时总时间复杂度为 O(n^2)
空间复杂度：O(height), 递归层数为树的高度
 */

func isBalanced2(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return abs(height(root.Left)-height(root.Right)) <= 1 && isBalanced2(root.Left) && isBalanced2(root.Right)
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(height(root.Left), height(root.Right))+1
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}