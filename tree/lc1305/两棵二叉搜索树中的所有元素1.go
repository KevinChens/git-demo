package lc1305

import "sort"

/*
分析：中序遍历，返回两个有序数组，再将两个有序数组升序合并
时间复杂度：O(n+m), n, m为两棵树节点
空间复杂度：O(n+m)
 */

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	nums1 := inOrder(root1)
	nums2 := inOrder(root2)
	// 拼接
	nums1 = append(nums1, nums2...)
	sort.Ints(nums1)
	return nums1
}

func inOrder(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := make([]int, 0)
	stack := make([]*TreeNode, 0)

	for root != nil || len(stack) > 0 {
		for root != nil {
			// push
			stack = append(stack, root)
			root = root.Left
		}
		// pop
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		res = append(res, node.Val)
		root = node.Right
	}
	return res
}
