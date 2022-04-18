package lc124

import "math"

/*
分析：递归法
时间复杂度：O(n),
空间复杂度：O(height), 
 */

func maxPathSum2(root *TreeNode) int {
	maxSum := math.MinInt32
	var maxGain func(node *TreeNode) int
	maxGain = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		// 递归计算左右子节点的最大贡献值
		// 只有最大贡献值大于0时，才会选取对应子节点
		leftGain := max(maxGain(node.Left), 0)
		rightGain := max(maxGain(node.Right), 0)

		// 节点的最大路径和取决于该节点的值与该节点的左右子节点的最大贡献者
		priceNewPath := node.Val + leftGain + rightGain
		// 更新答案
		maxSum = max(maxSum, priceNewPath)
		// 返回节点的最大贡献值
		return node.Val+max(leftGain, rightGain)
	}
	maxGain(root)
	return maxSum
}