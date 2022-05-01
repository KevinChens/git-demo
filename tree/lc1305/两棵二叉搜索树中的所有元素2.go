package lc1305


/*
分析：中序遍历+双指针归并
时间复杂度：O(m+n)
空间复杂度：O(m+n)
 */

func getAllElements2(root1 *TreeNode, root2 *TreeNode) []int {
	nums1 := inOrder2(root1)
	nums2 := inOrder2(root2)
	res := merge(nums1, nums2)
	return res
}

func inOrder2(root *TreeNode) []int {
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

func merge(nums1, nums2 []int) []int {
	// 双指针合并
	p1, l1 := 0, len(nums1)
	p2, l2 := 0, len(nums2)
	merged := make([]int, 0)

	for {
		if p1 == l1 {
			return append(merged, nums2[p2:]...)
		}
		if p2 == l2 {
			return append(merged, nums1[p1:]...)
		}
		if nums1[p1] < nums2[p2] {
			merged = append(merged, nums1[p1])
			p1++
		} else {
			merged = append(merged, nums2[p2])
			p2++
		}
	}
	return merged
}