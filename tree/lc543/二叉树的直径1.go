package lc543

/*
分析：每⼀条⼆叉树的直径⻓度，就是⼀个节点的左右⼦树的最⼤深度之和；
遍历整棵树中的每个节点，然后通过每个节点的左右⼦树的最⼤深度，
计算出每个节点的直径，最后把所有直径求个最⼤值即可。
时间复杂度：O(n^2), traverse遍历每个节点的时候还会调⽤递归函数maxDepth，⽽maxDepth是要遍历⼦树的所有节点的;
空间复杂度：O(height), 递归栈空间为二叉树的高度
 */

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

//记录最大直径的长度
var maxDiameter int

func diameterOfBinaryTree1(root *TreeNode) int {
	//全局变量初始化
	maxDiameter = 0
	//对每个节点计算直径，计算出最大直径
	traverse(root)
	return maxDiameter
}

// 遍历二叉树
func traverse(root *TreeNode) {
	if root == nil {
		return
	}
	//前序位置，在进入二叉树节点前执行
	//对每一个节点计算直径
	leftMax := maxDepth(root.Left)
	rightMax := maxDepth(root.Right)
	diameter := leftMax + rightMax
	//更新全局最大直径
	maxDiameter = max(diameter, maxDiameter)
	traverse(root.Left)
	traverse(root.Right)
}

// 计算二叉树的最大深度
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftMax := maxDepth(root.Left)
	rightMax := maxDepth(root.Right)
	return max(leftMax, rightMax) + 1
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
