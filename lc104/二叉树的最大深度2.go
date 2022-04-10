package lc104

/*
分析：分解问题计算答案的思路，动态规划思想, DFS；
⼀棵⼆叉树的最⼤深度可以通过⼦树的最⼤⾼度推导计算。
时间复杂度：O(n), 其中n为二叉树节点的个数，每个节点在递归中只被遍历一次。递归访问到空节点时退出。
空间复杂度：O(height), 需要递归计算左右子树的最大深度, 其中height表示二叉树的高度；
递归函数需要栈空间，而栈空间取决于递归的深度，因此空间复杂度等价于二叉树的高度。
*/

// 通过计算子树的最大高度，返回二叉树的最大深度
func maxDepth2(root *TreeNode) int {
	if root == nil {
		//遇到空节点返回
		return 0
	}
	//递归计算左右子树的最大深度
	leftMaxDepth := maxDepth2(root.Left)
	rightMaxDepth := maxDepth2(root.Right)
	//后序位置，在离开一个二叉树节点后执行
	//整颗二叉树的最大深度=max(leftMaxDepth, rightMaxDepth)+1
	maxDepth := func(x, y int) int {
		if x < y {
			return y
		}
		return x
	}(leftMaxDepth, rightMaxDepth)
	return maxDepth + 1
}

