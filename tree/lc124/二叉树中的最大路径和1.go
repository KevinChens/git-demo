package lc124

/*
分析：思路：分治法，分为三种情况：左子树最大路径和最大，右子树最大路径和最大，左右子树最大加根节点最大;
需要保存两个变量：一个保存子树最大路径和，一个保存左右加根节点和，然后比较这个两个变量选择最大值即可
时间复杂度：O(n),
空间复杂度：O(height),
 */

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

type ResultType struct {
	SinglePath int  // 保存单边最大值
	MaxPath int  // 保存最大值（单边或者两个单边+根的值）
}

func maxPathSum(root *TreeNode) int {
	result := helper(root)
	return result.MaxPath
}

func helper(root *TreeNode) ResultType {
	// check
	if root == nil {
		return ResultType{
			SinglePath: 0,
			MaxPath: -(1<<31),
		}
	}
	// divide
	left := helper(root.Left)
	right := helper(root.Right)
	// conquer
	result := ResultType{}
	// 求单边最大值
	if left.SinglePath > right.SinglePath {
		result.SinglePath = max(left.SinglePath+root.Val, 0)
	} else {
		result.SinglePath = max(right.SinglePath+root.Val, 0)
	}
	// 求两边加根最大值
	maxPath := max(right.MaxPath, left.MaxPath)
	result.MaxPath = max(maxPath, left.SinglePath+right.SinglePath+root.Val)
	return result
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
