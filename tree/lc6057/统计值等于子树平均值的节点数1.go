package lc6057

/*
分析：DFS, 自底向上，递归返回，保存总和
需要计算子树节点数量以及子树节点值之和，进行深度搜索递归调用，
返回值为子树节点值之和sum以及子树节点数量cnt，
递归出口为空节点，若节点不为空，先获得左右子树节点和及数量，
并计算当前节点是否满足，更新全局res值，向上层返回该子树节点和以及节点数量
时间复杂度：O(n)
空间复杂度：O(n)
 */

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

var ans int

func AverageOfSubtree(root *TreeNode) int {
	ans = 0
	dfs(root)
	return ans
}

func dfs(root *TreeNode) (int, int){
	sum, cnt := root.Val, 1
	if root == nil {
		return 0, 0
	}
	if root.Left != nil {
		s, c := dfs(root.Left)
		sum += s
		cnt += c
	}
	if root.Right != nil {
		s, c := dfs(root.Right)
		sum += s
		cnt += c
	}
	if root.Val == sum/cnt {
		ans++
	}
	return sum, cnt
}
