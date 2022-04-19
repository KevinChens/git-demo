package lc103

/*
分析：层序遍历+toggle切换，是否reverse入队，保证z字形
时间复杂度：O(n), 每个节点仅访问一次
空间复杂度：O(n), queue最大为n
 */

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	result := make([][]int, 0)
	queue := make([]*TreeNode, 0)
	// 切换，保证z字形入队
	toggle := false

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
		if toggle {
			reverse(list)
		}
		result = append(result, list)
		toggle = !toggle
	}
	return result
}

func reverse(nums []int) {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}