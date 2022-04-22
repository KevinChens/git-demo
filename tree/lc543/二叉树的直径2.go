package lc543

/*
分析：解法1使用前序位置⽆法获取⼦树信息，所以只能让每个节点调⽤maxDepth函数去算⼦树的深度；
那如何优化？
把计算直径的逻辑放在后序位置，放在maxDepth的后序位置，因为maxDepth的后序位置是知道左右⼦树的最⼤深度的。
时间复杂度：O(n), 只遍历一次二叉树
空间复杂度：O(height), 递归栈空间为二叉树的高度
 */

func diameterOfBinaryTree2(root *TreeNode) int {
	maxDiameter = 0
	//在计算左右子树最大深度的同时，计算直径
	maxDepth2(root)
	return maxDiameter
}

func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftMax := maxDepth2(root.Left)
	rightMax := maxDepth2(root.Right)
	//后序位置，计算最大直径
	diameter := leftMax + rightMax
	maxDiameter = max(diameter, maxDiameter)
	return max(leftMax, rightMax) + 1
}