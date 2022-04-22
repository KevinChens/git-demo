package lc110

/*
分析：分治法，左边平衡 && 右边平衡 && 左右两边高度 <= 1;
因为需要返回是否平衡及高度，要么返回两个数据，要么合并两个数据，
所以用-1表示不平衡，>0表示树高度（二义性：一个变量有两种含义）
时间复杂度：O(n), 其中n为二叉树节点的个数，每个节点在递归中只被遍历一次。递归访问到空节点时退出。
空间复杂度：O(height), 需要递归计算左右子树的最大深度, 其中height表示二叉树的高度；
递归函数需要栈空间，而栈空间取决于递归的深度，因此空间复杂度等价于二叉树的高度。
 */

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if maxDepth(root) == -1 {
		return false
	}
	return true
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	// 为什么返回-1？变量具有二义性
	if left == -1 || right == -1 || left-right > 1 || right-left > 1 {
		return -1
	}
	// >0表示树高度
	if left > right {
		return left+1
	}
	return right+1
}