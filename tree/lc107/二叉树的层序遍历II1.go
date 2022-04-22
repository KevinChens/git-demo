package lc107

/*
分析：层序遍历+翻转
时间复杂度：O(n), 每个节点仅访问一次
空间复杂度：O(n)，queue最大为n
 */

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
func levelOrderBottom(root *TreeNode) [][]int {
	result := levelOrder(root)
	// 翻转
	reverse(result)
	return result
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	result := make([][]int, 0)
	queue := make([]*TreeNode, 0)

	// push
	queue = append(queue, root)

	for len(queue) > 0 {
		list := make([]int, 0)
		l := len(queue)

		for i := 0; i < l; i++ {
			// pop
			node := queue[0]
			queue = queue[1:]

			list = append(list, node.Val)
			// push
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, list)
	}
	return result
}

func reverse(nums [][]int) {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}
