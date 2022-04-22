package lc98

/*
分析：中序遍历，检查结果列表是否已经有序
时间复杂度：O(n), 每个节点仅访问一次
空间复杂度：O(n), 递归不超过n层
 */

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
func isValidBST(root *TreeNode) bool {
	result := make([]int, 0)
	inOrder(root, &result)
	// check
	for i := 0; i < len(result)-1; i++ {
		if result[i] >= result[i+1] {
			return false
		}
	}
	return true
}

func inOrder(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}

	inOrder(root.Left, result)
	*result = append(*result, root.Val)
	inOrder(root.Right, result)
}
